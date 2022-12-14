// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTxns = &ffapi.Route{
	Name:            "getTxns",
	Path:            "transactions",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.TransactionQueryFactory,
	Description:     coremsgs.APIEndpointsGetTxns,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Transaction{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetTransactions(cr.ctx, r.Filter))
		},
	},
}
