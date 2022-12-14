// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getNetworkOrgs = &ffapi.Route{
	Name:            "getNetworkOrgs",
	Path:            "network/organizations",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.IdentityQueryFactory,
	Description:     coremsgs.APIEndpointsGetNetworkOrgs,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Identity{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.NetworkMap().GetOrganizations(cr.ctx, r.Filter))
		},
	},
}
