// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postTokenBurn = &ffapi.Route{
	Name:       "postTokenBurn",
	Path:       "tokens/burn",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsPostTokenBurn,
	JSONInputValue:  func() interface{} { return &core.TokenTransferInput{} },
	JSONOutputValue: func() interface{} { return &core.TokenTransfer{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			return cr.or.Assets().BurnTokens(cr.ctx, r.Input.(*core.TokenTransferInput), waitConfirm)
		},
	},
}
