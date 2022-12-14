// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNewIdentity = &ffapi.Route{
	Name:       "postNewIdentity",
	Path:       "identities",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsPostNewIdentity,
	JSONInputValue:  func() interface{} { return &core.IdentityCreateDTO{} },
	JSONOutputValue: func() interface{} { return &core.Identity{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			org, err := cr.or.NetworkMap().RegisterIdentity(cr.ctx, r.Input.(*core.IdentityCreateDTO), waitConfirm)
			return org, err
		},
	},
}
