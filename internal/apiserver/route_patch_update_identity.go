// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var patchUpdateIdentity = &ffapi.Route{
	Name:   "patchUpdateIdentity",
	Path:   "identities/{iid}",
	Method: http.MethodPatch,
	PathParams: []*ffapi.PathParam{
		{Name: "iid", Description: coremsgs.APIParamsIdentityID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsPatchUpdateIdentity,
	JSONInputValue:  func() interface{} { return &core.IdentityUpdateDTO{} },
	JSONOutputValue: func() interface{} { return &core.Identity{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			org, err := cr.or.NetworkMap().UpdateIdentity(cr.ctx, r.PP["iid"], r.Input.(*core.IdentityUpdateDTO), waitConfirm)
			return org, err
		},
	},
}
