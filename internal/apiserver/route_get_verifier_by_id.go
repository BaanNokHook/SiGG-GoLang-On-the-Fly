// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getVerifierByID = &ffapi.Route{
	Name:   "getVerifierByID",
	Path:   "verifiers/{hash}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "hash", Example: "hash", Description: coremsgs.APIParamsVerifierHash},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetVerifierByHash,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Verifier{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.NetworkMap().GetVerifierByHash(cr.ctx, r.PP["hash"])
		},
	},
}
