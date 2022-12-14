package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
)

var deleteSubscription = &ffapi.Route{
	Name:   "deleteSubscription",
	Path:   "subscriptions/{subid}",
	Method: http.MethodDelete,
	PathParams: []*ffapi.PathParam{
		{Name: "subid", Description: coremsgs.APIParamsSubscriptionID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsDeleteSubscription,
	JSONInputValue:  nil,
	JSONOutputValue: nil,
	JSONOutputCodes: []int{http.StatusNoContent}, // Sync operation, no output
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			err = cr.or.DeleteSubscription(cr.ctx, r.PP["subid"])
			return nil, err
		},
	},
}
