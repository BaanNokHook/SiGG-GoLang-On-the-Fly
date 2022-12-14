// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postContractAPIInvoke = &ffapi.Route{
	Name:   "postContractAPIInvoke",
	Path:   "apis/{apiName}/invoke/{methodPath}",
	Method: http.MethodPost,
	PathParams: []*ffapi.PathParam{
		{Name: "apiName", Description: coremsgs.APIParamsContractAPIName},
		{Name: "methodPath", Description: coremsgs.APIParamsMethodPath},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true, Example: "true"},
	},
	Description:     coremsgs.APIEndpointsPostContractAPIInvoke,
	JSONInputValue:  func() interface{} { return &core.ContractCallRequest{} },
	JSONOutputValue: func() interface{} { return &core.Operation{} },
	JSONOutputCodes: []int{http.StatusOK, http.StatusAccepted},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			req := r.Input.(*core.ContractCallRequest)
			req.Type = core.CallTypeInvoke
			return cr.or.Contracts().InvokeContractAPI(cr.ctx, r.PP["apiName"], r.PP["methodPath"], req, waitConfirm)
		},
	},
}
