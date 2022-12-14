package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var getGroupByHash = &ffapi.Route{
	Name:   "getGroupByHash",
	Path:   "groups/{hash}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "hash", Description: coremsgs.APIParamsGroupHash},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetGroupByHash,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Group{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.PrivateMessaging() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.PrivateMessaging().GetGroupByID(cr.ctx, r.PP["hash"])
		},
	},
}
