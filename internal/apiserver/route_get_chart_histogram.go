package apiserver

import (
	"net/http"
	"strconv"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getChartHistogram = &ffapi.Route{
	Name:   "getChartHistogram",
	Path:   "charts/histogram/{collection}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "collection", Description: coremsgs.APIParamsCollectionID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "startTime", Description: coremsgs.APIHistogramStartTimeParam, IsBool: false},
		{Name: "endTime", Description: coremsgs.APIHistogramEndTimeParam, IsBool: false},
		{Name: "buckets", Description: coremsgs.APIHistogramBucketsParam, IsBool: false},
	},
	Description:     coremsgs.APIEndpointsGetChartHistogram,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.ChartHistogram{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			startTime, err := fftypes.ParseTimeString(r.QP["startTime"])
			if err != nil {
				return nil, i18n.NewError(cr.ctx, coremsgs.MsgInvalidChartNumberParam, "startTime")
			}
			endTime, err := fftypes.ParseTimeString(r.QP["endTime"])
			if err != nil {
				return nil, i18n.NewError(cr.ctx, coremsgs.MsgInvalidChartNumberParam, "endTime")
			}
			buckets, err := strconv.ParseInt(r.QP["buckets"], 10, 64)
			if err != nil {
				return nil, i18n.NewError(cr.ctx, coremsgs.MsgInvalidChartNumberParam, "buckets")
			}
			return cr.or.GetChartHistogram(cr.ctx, startTime.UnixNano(), endTime.UnixNano(), buckets, database.CollectionName(r.PP["collection"]))
		},
	},
}
