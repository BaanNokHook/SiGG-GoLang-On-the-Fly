// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var spiGetOps = &ffapi.Route{
	Name:            "spiGetOps",
	Path:            "namespaces/{ns}/operations",
	Method:          http.MethodGet,
	QueryParams:     nil,
	FilterFactory:   database.OperationQueryFactory,
	Description:     coremsgs.APIEndpointsAdminGetOps,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Operation{} },
	JSONOutputCodes: []int{http.StatusOK},
	Tag:             routeTagNonDefaultNamespace,
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetOperations(cr.ctx, r.Filter))
		},
	},
}
