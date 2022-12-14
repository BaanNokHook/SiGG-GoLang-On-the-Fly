// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getNetworkNode = &ffapi.Route{
	Name:   "getNetworkNode",
	Path:   "network/nodes/{nameOrId}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "nameOrId", Description: coremsgs.APIParamsNodeNameOrID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetNetworkNode,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Identity{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.NetworkMap().GetNodeByNameOrID(cr.ctx, r.PP["nameOrId"])
			return output, err
		},
	},
}
