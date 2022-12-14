// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postVerifiersResolve = &ffapi.Route{
	Name:            "postVerifiersResolve",
	Path:            "verifiers/resolve",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostVerifiersResolve,
	JSONInputValue:  func() interface{} { return &core.VerifierRef{} },
	JSONOutputValue: func() interface{} { return &core.VerifierRef{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Identity().ResolveInputSigningKey(cr.ctx, r.Input.(*core.VerifierRef))
		},
	},
}
