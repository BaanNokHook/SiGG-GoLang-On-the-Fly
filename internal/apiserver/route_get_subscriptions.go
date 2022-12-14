// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getSubscriptions = &ffapi.Route{
	Name:            "getSubscriptions",
	Path:            "subscriptions",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.SubscriptionQueryFactory,
	Description:     coremsgs.APIEndpointsGetSubscriptions,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Subscription{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetSubscriptions(cr.ctx, r.Filter))
		},
	},
}
