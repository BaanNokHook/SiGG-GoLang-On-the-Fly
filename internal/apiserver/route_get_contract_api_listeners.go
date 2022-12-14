package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getContractAPIListeners = &ffapi.Route{
	Name:   "getContractAPIListeners",
	Path:   "apis/{apiName}/listeners/{eventPath}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "apiName", Description: coremsgs.APIParamsContractAPIName},
		{Name: "eventPath", Description: coremsgs.APIParamsEventPath},
	},
	QueryParams:     []*ffapi.QueryParam{},
	FilterFactory:   database.ContractListenerQueryFactory,
	Description:     coremsgs.APIEndpointsGetContractListeners,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.ContractListener{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Contracts().GetContractAPIListeners(cr.ctx, r.PP["apiName"], r.PP["eventPath"], r.Filter))
		},
	},
}
