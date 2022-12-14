// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTokenAccounts = &ffapi.Route{
	Name:            "getTokenAccounts",
	Path:            "tokens/accounts",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.TokenAccountQueryFactory,
	Description:     coremsgs.APIEndpointsGetTokenAccounts,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenAccount{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Assets().GetTokenAccounts(cr.ctx, r.Filter))
		},
	},
}
