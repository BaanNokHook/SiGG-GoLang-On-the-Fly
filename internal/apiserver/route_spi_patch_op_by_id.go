// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var spiPatchOpByID = &ffapi.Route{
	Name:   "spiPatchOpByID",
	Path:   "operations/{nsopid}",
	Method: http.MethodPatch,
	PathParams: []*ffapi.PathParam{
		{Name: "nsopid", Description: coremsgs.APIParamsOperationNamespacedID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsAdminPatchOpByID,
	JSONInputValue:  func() interface{} { return &core.OperationUpdateDTO{} },
	JSONOutputValue: func() interface{} { return &core.EmptyInput{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			err = cr.mgr.ResolveOperationByNamespacedID(cr.ctx, r.PP["nsopid"], r.Input.(*core.OperationUpdateDTO))
			return &core.EmptyInput{}, err
		},
	},
}
