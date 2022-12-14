// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getNamespaces = &ffapi.Route{
	Name:            "getNamespaces",
	Path:            "namespaces",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   nil,
	Description:     coremsgs.APIEndpointsGetNamespaces,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Namespace{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.mgr.GetNamespaces(cr.ctx)
		},
	},
}
