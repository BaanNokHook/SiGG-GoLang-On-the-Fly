// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
)

var postContractInterfaceGenerate = &ffapi.Route{
	Name:            "postGenerateContractInterface",
	Path:            "contracts/interfaces/generate",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     []*ffapi.QueryParam{},
	Description:     coremsgs.APIEndpointsPostContractInterfaceGenerate,
	JSONInputValue:  func() interface{} { return &fftypes.FFIGenerationRequest{} },
	JSONOutputValue: func() interface{} { return &fftypes.FFI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			generationRequest := r.Input.(*fftypes.FFIGenerationRequest)
			return cr.or.Contracts().GenerateFFI(cr.ctx, generationRequest)
		},
	},
}
