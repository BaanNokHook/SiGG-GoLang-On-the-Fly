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

var postContractDeploy = &ffapi.Route{
	Name:       "postContractDeploy",
	Path:       "contracts/deploy",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true, Example: "true"},
	},
	Description:     coremsgs.APIEndpointsPostContractDeploy,
	JSONInputValue:  func() interface{} { return &core.ContractDeployRequest{} },
	JSONOutputValue: func() interface{} { return &core.Operation{} },
	JSONOutputCodes: []int{http.StatusOK, http.StatusAccepted},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			req := r.Input.(*core.ContractDeployRequest)
			return cr.or.Contracts().DeployContract(cr.ctx, req, waitConfirm)
		},
	},
}
