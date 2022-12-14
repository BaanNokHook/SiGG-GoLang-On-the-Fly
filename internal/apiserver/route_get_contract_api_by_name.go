package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var getContractAPIByName = &ffapi.Route{
	Name:   "getContractAPIByName",
	Path:   "apis/{apiName}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "apiName", Description: coremsgs.APIParamsContractAPIName},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetContractAPIByName,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.ContractAPI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Contracts().GetContractAPI(cr.ctx, cr.apiBaseURL, r.PP["apiName"])
		},
	},
}
