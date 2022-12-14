// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getTxnStatus = &ffapi.Route{
	Name:   "getTxnStatus",
	Path:   "transactions/{txnid}/status",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "txnid", Description: coremsgs.APIParamsTransactionID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetTxnStatus,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.TransactionStatus{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.GetTransactionStatus(cr.ctx, r.PP["txnid"])
		},
	},
}
