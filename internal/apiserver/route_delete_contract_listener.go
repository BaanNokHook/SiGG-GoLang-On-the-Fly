
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
)

var deleteContractListener = &ffapi.Route{
	Name:   "deleteContractListener",
	Path:   "contracts/listeners/{nameOrId}",
	Method: http.MethodDelete,
	PathParams: []*ffapi.PathParam{
		{Name: "nameOrId", Description: coremsgs.APIParamsContractListenerNameOrID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsDeleteContractListener,
	JSONInputValue:  nil,
	JSONOutputValue: nil,
	JSONOutputCodes: []int{http.StatusNoContent}, // Sync operation, no output
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			err = cr.or.Contracts().DeleteContractListenerByNameOrID(cr.ctx, r.PP["nameOrId"])
			return nil, err
		},
	},
}
