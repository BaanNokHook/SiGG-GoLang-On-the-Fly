// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTxnByID = &ffapi.Route{
	Name:   "getTxnByID",
	Path:   "transactions/{txnid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "txnid", Description: coremsgs.APIParamsTransactionID},
	},
	QueryParams:     nil,
	FilterFactory:   database.TransactionQueryFactory,
	Description:     coremsgs.APIEndpointsGetTxnByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Transaction{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetTransactionByID(cr.ctx, r.PP["txnid"])
			return output, err
		},
	},
}
