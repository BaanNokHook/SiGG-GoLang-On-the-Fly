// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getNetworkIdentities = &ffapi.Route{
	Name:   "getNetworkIdentities",
	Path:   "network/identities",
	Method: http.MethodGet,
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchverifiers", Example: "true", Description: coremsgs.APIParamsFetchVerifiers, IsBool: true},
	},
	FilterFactory:   database.IdentityQueryFactory,
	Description:     coremsgs.APIEndpointsGetNetworkIdentities,
	Deprecated:      true, // use getIdentities instead
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &[]*core.IdentityWithVerifiers{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchverifiers"], "true") {
				return r.FilterResult(cr.or.NetworkMap().GetIdentitiesWithVerifiers(cr.ctx, r.Filter))
			}
			return r.FilterResult(cr.or.NetworkMap().GetIdentities(cr.ctx, r.Filter))
		},
	},
}
