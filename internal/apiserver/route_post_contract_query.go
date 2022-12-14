// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postContractQuery = &ffapi.Route{
	Name:            "postContractQuery",
	Path:            "contracts/query",
	Method:          http.MethodPost,
	PathParams:      nil,
	Description:     coremsgs.APIEndpointsPostContractQuery,
	QueryParams:     []*ffapi.QueryParam{},
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
			return cr.or.Contracts().InvokeContract(cr.ctx, req, true)
		},
	},
}
