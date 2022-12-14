// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getTokenConnectors = &ffapi.Route{
	Name:            "getTokenConnectors",
	Path:            "tokens/connectors",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetTokenConnectors,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenConnector{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Assets().GetTokenConnectors(cr.ctx), nil
		},
	},
}
