// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getOpByID = &ffapi.Route{
	Name:   "getOpByID",
	Path:   "operations/{opid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "opid", Description: coremsgs.APIParamsOperationIDGet},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetOpByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Operation{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetOperationByID(cr.ctx, r.PP["opid"])
			return output, err
		},
	},
}
