// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTokenAccountPools = &ffapi.Route{
	Name:   "getTokenAccountPools",
	Path:   "tokens/accounts/{key}/pools",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "key", Description: coremsgs.APIParamsTokenAccountKey},
	},
	QueryParams:     nil,
	FilterFactory:   database.TokenAccountPoolQueryFactory,
	Description:     coremsgs.APIEndpointsGetTokenAccountPools,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenAccountPool{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Assets().GetTokenAccountPools(cr.ctx, r.PP["key"], r.Filter))
		},
	},
}
