// SiGG-GoLang-On-the-Fly //
package orchestrator

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

func (or *orchestrator) getHistogramIntervals(startTime int64, endTime int64, numBuckets int64) (intervals []core.ChartHistogramInterval) {
	timeIntervalLength := (endTime - startTime) / numBuckets

	for i := startTime; i < endTime; i += timeIntervalLength {
		intervals = append(intervals, core.ChartHistogramInterval{
			StartTime: fftypes.UnixTime(i),
			EndTime:   fftypes.UnixTime(i + timeIntervalLength),
		})
	}

	return intervals
}

func (or *orchestrator) GetChartHistogram(ctx context.Context, startTime int64, endTime int64, buckets int64, collection database.CollectionName) ([]*core.ChartHistogram, error) {
	if buckets > core.ChartHistogramMaxBuckets || buckets < core.ChartHistogramMinBuckets {
		return nil, i18n.NewError(ctx, coremsgs.MsgInvalidNumberOfIntervals, core.ChartHistogramMinBuckets, core.ChartHistogramMaxBuckets)
	}
	if startTime > endTime {
		return nil, i18n.NewError(ctx, coremsgs.MsgHistogramInvalidTimes)
	}

	intervals := or.getHistogramIntervals(startTime, endTime, buckets)

	histogram, err := or.database().GetChartHistogram(ctx, or.namespace.Name, intervals, collection)
	if err != nil {
		return nil, err
	}

	return histogram, nil
}
