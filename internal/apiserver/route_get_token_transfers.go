// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTokenTransfers = &ffapi.Route{
	Name:       "getTokenTransfers",
	Path:       "tokens/transfers",
	Method:     http.MethodGet,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "fromOrTo", Description: coremsgs.APIParamsTokenTransferFromOrTo},
	},
	FilterFactory:   database.TokenTransferQueryFactory,
	Description:     coremsgs.APIEndpointsGetTokenTransfers,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenTransfer{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			filter := r.Filter
			if fromOrTo, ok := r.QP["fromOrTo"]; ok {
				fb := database.TokenTransferQueryFactory.NewFilter(cr.ctx)
				filter = filter.Condition(
					fb.Or().
						Condition(fb.Eq("from", fromOrTo)).
						Condition(fb.Eq("to", fromOrTo)))
			}
			return r.FilterResult(cr.or.Assets().GetTokenTransfers(cr.ctx, filter))
		},
	},
}
