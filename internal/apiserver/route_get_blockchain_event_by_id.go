package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getBlockchainEventByID = &ffapi.Route{
	Name:   "getBlockchainEventByID",
	Path:   "blockchainevents/{id}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "id", Description: coremsgs.APIParamsBlockchainEventID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetBlockchainEventByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.BlockchainEvent{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.GetBlockchainEventByID(cr.ctx, r.PP["id"])
		},
	},
}
