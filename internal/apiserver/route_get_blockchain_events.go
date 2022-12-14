
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getBlockchainEvents = &ffapi.Route{
	Name:            "getBlockchainEvents",
	Path:            "blockchainevents",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.BlockchainEventQueryFactory,
	Description:     coremsgs.APIEndpointsListBlockchainEvents,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.BlockchainEvent{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetBlockchainEvents(cr.ctx, r.Filter))
		},
	},
}
