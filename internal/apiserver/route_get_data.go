package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getData = &ffapi.Route{
	Name:            "getData",
	Path:            "data",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.DataQueryFactory,
	Description:     coremsgs.APIEndpointsGetData,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return core.DataArray{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetData(cr.ctx, r.Filter))
		},
	},
}
