// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postContractAPIQuery = &ffapi.Route{
	Name:   "postContractAPIQuery",
	Path:   "apis/{apiName}/query/{methodPath}",
	Method: http.MethodPost,
	PathParams: []*ffapi.PathParam{
		{Name: "apiName", Description: coremsgs.APIParamsContractAPIName},
		{Name: "methodPath", Description: coremsgs.APIParamsMethodPath},
	},
	QueryParams:     []*ffapi.QueryParam{},
	Description:     coremsgs.APIEndpointsPostContractAPIQuery,
	JSONInputValue:  func() interface{} { return &core.ContractCallRequest{} },
	JSONOutputValue: func() interface{} { return make(map[string]interface{}) },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			req := r.Input.(*core.ContractCallRequest)
			req.Type = core.CallTypeQuery
			return cr.or.Contracts().InvokeContractAPI(cr.ctx, r.PP["apiName"], r.PP["methodPath"], req, true)
		},
	},
}
