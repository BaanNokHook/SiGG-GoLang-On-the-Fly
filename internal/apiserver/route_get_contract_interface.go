package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
)

var getContractInterface = &ffapi.Route{
	Name:   "getContractInterface",
	Path:   "contracts/interfaces/{interfaceId}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "interfaceId", Description: coremsgs.APIParamsContractInterfaceID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchchildren", Example: "true", Description: coremsgs.APIParamsContractInterfaceFetchChildren, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsGetContractInterface,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &fftypes.FFI{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			interfaceID, err := fftypes.ParseUUID(cr.ctx, r.PP["interfaceId"])
			if err != nil {
				return nil, err
			}
			if strings.EqualFold(r.QP["fetchchildren"], "true") {
				return cr.or.Contracts().GetFFIByIDWithChildren(cr.ctx, interfaceID)
			}
			return cr.or.Contracts().GetFFIByID(cr.ctx, interfaceID)
		},
	},
}
