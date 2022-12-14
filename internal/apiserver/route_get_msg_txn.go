package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getMsgTxn = &ffapi.Route{
	Name:   "getMsgTxn",
	Path:   "messages/{msgid}/transaction",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "msgid", Description: coremsgs.APIParamsMessageID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetMsgTxn,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Transaction{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetMessageTransaction(cr.ctx, r.PP["msgid"])
			return output, err
		},
	},
}
