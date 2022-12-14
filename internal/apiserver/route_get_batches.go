package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getBatches = &ffapi.Route{
	Name:            "getBatches",
	Path:            "batches",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.BatchQueryFactory,
	Description:     coremsgs.APIEndpointsGetBatches,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.BatchPersisted{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetBatches(cr.ctx, r.Filter))
		},
	},
}
