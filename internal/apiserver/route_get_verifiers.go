// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getVerifiers = &ffapi.Route{
	Name:            "getVerifiers",
	Path:            "verifiers",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.VerifierQueryFactory,
	Description:     coremsgs.APIEndpointsGetVerifiers,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &[]*core.Verifier{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.NetworkMap().GetVerifiers(cr.ctx, r.Filter))
		},
	},
}
