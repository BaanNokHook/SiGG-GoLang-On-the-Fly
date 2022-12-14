package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/database"
)

var getContractInterfaces = &ffapi.Route{
	Name:            "getContractInterfaces",
	Path:            "contracts/interfaces",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.FFIQueryFactory,
	Description:     coremsgs.APIEndpointsGetContractInterfaces,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*fftypes.FFI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.Contracts().GetFFIs(cr.ctx, r.Filter))
		},
	},
}
