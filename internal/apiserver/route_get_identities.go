package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getIdentities = &ffapi.Route{
	Name:       "getIdentities",
	Path:       "identities",
	Method:     http.MethodGet,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchverifiers", Example: "true", Description: coremsgs.APIParamsFetchVerifiers, IsBool: true},
	},
	FilterFactory:   database.IdentityQueryFactory,
	Description:     coremsgs.APIEndpointsGetIdentities,
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
