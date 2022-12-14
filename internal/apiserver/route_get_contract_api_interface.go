package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
)

var getContractAPIInterface = &ffapi.Route{
	Name:   "getContractAPIInterface",
	Path:   "apis/{apiName}/interface",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "apiName", Description: coremsgs.APIParamsContractAPIName},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetContractAPIInterface,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &fftypes.FFI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Contracts().GetContractAPIInterface(cr.ctx, r.PP["apiName"])
		},
	},
}
