// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var putContractAPI = &ffapi.Route{
	Name:   "putContractAPI",
	Path:   "apis/{id}",
	Method: http.MethodPut,
	PathParams: []*ffapi.PathParam{
		{Name: "id", Example: "id", Description: coremsgs.APIParamsContractAPIName},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true, Example: "true"},
	},
	Description:     coremsgs.APIParamsContractAPIID,
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
			api.ID, err = fftypes.ParseUUID(cr.ctx, r.PP["id"])
			if err == nil {
				err = cr.or.DefinitionSender().DefineContractAPI(cr.ctx, cr.apiBaseURL, api, waitConfirm)
			}
			return api, err
		},
	},
}
