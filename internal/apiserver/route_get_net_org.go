// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getNetworkOrg = &ffapi.Route{
	Name:   "getNetworkOrg",
	Path:   "network/organizations/{nameOrId}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "nameOrId", Description: coremsgs.APIParamsOrgNameOrID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetNetworkOrg,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Identity{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.NetworkMap().GetOrganizationByNameOrID(cr.ctx, r.PP["nameOrId"])
			return output, err
		},
	},
}
