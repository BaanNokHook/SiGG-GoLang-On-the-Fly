// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var spiGetNamespaces = &ffapi.Route{
	Name:            "spiGetNamespaces",
	Path:            "namespaces",
	Method:          http.MethodGet,
	QueryParams:     nil,
	FilterFactory:   nil,
	Description:     coremsgs.APIEndpointsAdminGetNamespaces,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Namespace{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.mgr.GetNamespaces(cr.ctx)
		},
	},
}
