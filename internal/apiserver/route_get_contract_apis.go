package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getContractAPIs = &ffapi.Route{
	Name:            "getContractAPIs",
	Path:            "apis",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.ContractAPIQueryFactory,
	Description:     coremsgs.APIEndpointsGetContractAPIs,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.ContractAPI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Contracts().GetContractAPIs(cr.ctx, cr.apiBaseURL, r.Filter))
		},
	},
}
