// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTokenPools = &ffapi.Route{
	Name:            "getTokenPools",
	Path:            "tokens/pools",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.TokenPoolQueryFactory,
	Description:     coremsgs.APIEndpointsGetTokenPools,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenPool{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Assets().GetTokenPools(cr.ctx, r.Filter))
		},
	},
}
