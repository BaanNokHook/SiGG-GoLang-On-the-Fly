// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getTxnBlockchainEvents = &ffapi.Route{
	Name:   "getTxnBlockchainEvents",
	Path:   "transactions/{txnid}/blockchainevents",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "txnid", Description: coremsgs.APIParamsTransactionID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetTxnBlockchainEvents,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &[]*core.BlockchainEvent{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetTransactionBlockchainEvents(cr.ctx, r.PP["txnid"]))
		},
	},
}
