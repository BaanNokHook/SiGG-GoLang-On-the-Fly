// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNewSubscription = &ffapi.Route{
	Name:            "postNewSubscription",
	Path:            "subscriptions",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostNewSubscription,
	JSONInputValue:  func() interface{} { return &core.Subscription{} },
	JSONOutputValue: func() interface{} { return &core.Subscription{} },
	JSONOutputCodes: []int{http.StatusCreated}, // Sync operation
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.CreateSubscription(cr.ctx, r.Input.(*core.Subscription))
			return output, err
		},
	},
}
