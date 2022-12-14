// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getTokenTransferByID = &ffapi.Route{
	Name:   "getTokenTransferByID",
	Path:   "tokens/transfers/{transferId}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "transferId", Description: coremsgs.APIParamsTokenTransferID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetTokenTransferByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.TokenTransfer{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.Assets().GetTokenTransferByID(cr.ctx, r.PP["transferId"])
			return output, err
		},
	},
}
