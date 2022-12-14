package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getIdentityVerifiers = &ffapi.Route{
	Name:   "getIdentityVerifiers",
	Path:   "identities/{iid}/verifiers",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "iid", Example: "id", Description: coremsgs.APIParamsIdentityID},
	},
	QueryParams:     nil,
	FilterFactory:   database.VerifierQueryFactory,
	Description:     coremsgs.APIEndpointsGetIdentityVerifiers,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &[]*core.Verifier{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.NetworkMap().GetIdentityVerifiers(cr.ctx, r.PP["iid"], r.Filter))
		},
	},
}
