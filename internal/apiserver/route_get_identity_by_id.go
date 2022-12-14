package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getIdentityByID = &ffapi.Route{
	Name:   "getIdentityByID",
	Path:   "identities/{iid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "iid", Example: "id", Description: coremsgs.APIParamsIdentityID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchverifiers", Example: "true", Description: coremsgs.APIParamsFetchVerifiers, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsGetIdentityByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Identity{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchverifiers"], "true") {
				return cr.or.NetworkMap().GetIdentityByIDWithVerifiers(cr.ctx, r.PP["iid"])
			}
			return cr.or.NetworkMap().GetIdentityByID(cr.ctx, r.PP["iid"])
		},
	},
}
