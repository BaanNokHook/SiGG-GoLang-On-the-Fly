// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getSubscriptionByID = &ffapi.Route{
	Name:   "getSubscriptionByID",
	Path:   "subscriptions/{subid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "subid", Description: coremsgs.APIParamsSubscriptionID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchstatus", Description: coremsgs.APIParamsFetchStatus, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsGetSubscriptionByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Subscription{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchstatus"], "true") {
				return cr.or.GetSubscriptionByIDWithStatus(cr.ctx, r.PP["subid"])
			}
			return cr.or.GetSubscriptionByID(cr.ctx, r.PP["subid"])
		},
	},
}
