package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/networkmap"
)

var getIdentityDID = &ffapi.Route{
	Name:   "getIdentityDID",
	Path:   "identities/{iid}/did",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "iid", Example: "id", Description: coremsgs.APIParamsIdentityID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetIdentityDID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &networkmap.DIDDocument{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.NetworkMap().GetDIDDocForIndentityByID(cr.ctx, r.PP["iid"])
		},
	},
}
