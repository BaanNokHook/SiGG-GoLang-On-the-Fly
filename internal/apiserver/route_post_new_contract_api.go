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

var postNewContractAPI = &ffapi.Route{
	Name:       "postNewContractAPI",
	Path:       "apis",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true, Example: "true"},
	},
	Description:     coremsgs.APIEndpointsPostNewContractAPI,
	JSONInputValue:  func() interface{} { return &core.ContractAPI{} },
	JSONOutputValue: func() interface{} { return &core.ContractAPI{} },
	JSONOutputCodes: []int{http.StatusOK, http.StatusAccepted},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			api := r.Input.(*core.ContractAPI)
			err = cr.or.DefinitionSender().DefineContractAPI(cr.ctx, cr.apiBaseURL, api, waitConfirm)
			return api, err
		},
	},
}
