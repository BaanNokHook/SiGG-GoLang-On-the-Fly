package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
)

var getContractInterfaceNameVersion = &ffapi.Route{
	Name:   "getContractInterfaceByNameAndVersion",
	Path:   "contracts/interfaces/{name}/{version}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "name", Description: coremsgs.APIParamsContractInterfaceName},
		{Name: "version", Description: coremsgs.APIParamsContractInterfaceVersion},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchchildren", Example: "true", Description: coremsgs.APIParamsContractInterfaceFetchChildren, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsGetContractInterfaceNameVersion,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &fftypes.FFI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchchildren"], "true") {
				return cr.or.Contracts().GetFFIWithChildren(cr.ctx, r.PP["name"], r.PP["version"])
			}
			return cr.or.Contracts().GetFFI(cr.ctx, r.PP["name"], r.PP["version"])
		},
	},
}
