// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var spiGetOpByID = &ffapi.Route{
	Name:   "spiGetOpByID",
	Path:   "operations/{nsopid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "nsopid", Description: coremsgs.APIParamsOperationNamespacedID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsAdminGetOpByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Operation{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.mgr.GetOperationByNamespacedID(cr.ctx, r.PP["nsopid"])
			return output, err
		},
	},
}
