// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postContractAPIListeners = &ffapi.Route{
	Name:   "postContractAPIListeners",
	Path:   "apis/{apiName}/listeners/{eventPath}",
	Method: http.MethodPost,
	PathParams: []*ffapi.PathParam{
		{Name: "apiName", Description: coremsgs.APIParamsContractAPIName},
		{Name: "eventPath", Description: coremsgs.APIParamsEventPath},
	},
	QueryParams:     []*ffapi.QueryParam{},
	Description:     coremsgs.APIEndpointsPostNewContractListener,
	JSONInputValue:  func() interface{} { return &core.ContractListener{} },
	JSONOutputValue: func() interface{} { return &core.ContractListener{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Contracts().AddContractAPIListener(cr.ctx, r.PP["apiName"], r.PP["eventPath"], r.Input.(*core.ContractListener))
		},
	},
}
