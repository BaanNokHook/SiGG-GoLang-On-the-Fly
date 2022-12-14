// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postTokenApproval = &ffapi.Route{
	Name:       "postTokenApproval",
	Path:       "tokens/approvals",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	Description: coremsgs.APIEndpointsPostTokenApproval,
	JSONInputValue: func() interface{} {
		return &core.TokenApprovalInput{
			TokenApproval: core.TokenApproval{
				Approved: true,
			},
		}
	},
	JSONOutputValue: func() interface{} { return &core.TokenApproval{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			return cr.or.Assets().TokenApproval(cr.ctx, r.Input.(*core.TokenApprovalInput), waitConfirm)
		},
	},
}
