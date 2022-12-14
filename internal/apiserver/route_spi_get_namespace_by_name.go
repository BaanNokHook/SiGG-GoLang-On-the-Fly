// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var spiGetNamespaceByName = &ffapi.Route{
	Name:   "spiGetNamespaceByName",
	Path:   "namespaces/{ns}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "ns", Description: coremsgs.APIParamsNamespace},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostNewNamespace,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Namespace{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			or, err := getOrchestrator(cr.ctx, cr.mgr, routeTagNonDefaultNamespace, r)
			if err == nil {
				output = or.GetNamespace(cr.ctx)
			}
			return output, err
		},
	},
}
