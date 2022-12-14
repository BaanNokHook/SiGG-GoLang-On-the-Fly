package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var getContractListenerByNameOrID = &ffapi.Route{
	Name:   "getContractListenerByNameOrID",
	Path:   "contracts/listeners/{nameOrId}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "nameOrId", Description: coremsgs.APIParamsContractListenerNameOrID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchstatus", Example: "true", Description: coremsgs.APIParamsFetchStatus, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsGetContractListenerByNameOrID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.ContractListener{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Contracts() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchstatus"], "true") {
				return cr.or.Contracts().GetContractListenerByNameOrIDWithStatus(cr.ctx, r.PP["nameOrId"])
			}
			return cr.or.Contracts().GetContractListenerByNameOrID(cr.ctx, r.PP["nameOrId"])
		},
	},
}
