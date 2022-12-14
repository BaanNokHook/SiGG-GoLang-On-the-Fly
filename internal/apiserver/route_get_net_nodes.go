// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getNetworkNodes = &ffapi.Route{
	Name:            "getNetworkNodes",
	Path:            "network/nodes",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.IdentityQueryFactory,
	Description:     coremsgs.APIEndpointsGetNetworkNodes,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Identity{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.NetworkMap().GetNodes(cr.ctx, r.Filter))
		},
	},
}
