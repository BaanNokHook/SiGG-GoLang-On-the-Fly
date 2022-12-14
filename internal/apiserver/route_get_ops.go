// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getOps = &ffapi.Route{
	Name:            "getOps",
	Path:            "operations",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.OperationQueryFactory,
	Description:     coremsgs.APIEndpointsGetOps,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Operation{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetOperations(cr.ctx, r.Filter))
		},
	},
}
