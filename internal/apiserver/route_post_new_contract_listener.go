// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNewContractListener = &ffapi.Route{
	Name:            "postNewContractListener",
	Path:            "contracts/listeners",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostNewContractListener,
	JSONInputValue:  func() interface{} { return &core.ContractListenerInput{} },
	JSONOutputValue: func() interface{} { return &core.ContractListener{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Contracts().AddContractListener(cr.ctx, r.Input.(*core.ContractListenerInput))
		},
	},
}
