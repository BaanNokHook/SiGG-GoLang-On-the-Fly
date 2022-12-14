package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getIdentityByDID = &ffapi.Route{
	Name:   "getIdentityByDID",
	Path:   "identities/{did:did:.+}",
	Method: http.MethodGet,
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchverifiers", Example: "true", Description: coremsgs.APIParamsFetchVerifiers, IsBool: true},
	},
	PathParams: []*ffapi.PathParam{
		{Name: "did", Description: coremsgs.APIParamsDID},
	},
	Description:     coremsgs.APIEndpointsGetIdentityByDID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.IdentityWithVerifiers{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchverifiers"], "true") {
				return cr.or.NetworkMap().GetIdentityByDIDWithVerifiers(cr.ctx, r.PP["did"])
			}
			return cr.or.NetworkMap().GetIdentityByDID(cr.ctx, r.PP["did"])
		},
	},
}
