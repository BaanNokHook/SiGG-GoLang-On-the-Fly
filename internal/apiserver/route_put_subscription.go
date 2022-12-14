// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var putSubscription = &ffapi.Route{
	Name:            "putSubscription",
	Path:            "subscriptions",
	Method:          http.MethodPut,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPutSubscription,
	JSONInputValue:  func() interface{} { return &core.Subscription{} },
	JSONOutputValue: func() interface{} { return &core.Subscription{} },
	JSONOutputCodes: []int{http.StatusOK}, // Sync operation
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.CreateUpdateSubscription(cr.ctx, r.Input.(*core.Subscription))
			return output, err
		},
	},
}
