// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTokenBalances = &ffapi.Route{
	Name:            "getTokenBalances",
	Path:            "tokens/balances",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.TokenBalanceQueryFactory,
	Description:     coremsgs.APIEndpointsGetTokenBalances,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenBalance{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Assets().GetTokenBalances(cr.ctx, r.Filter))
		},
	},
}
