package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getEvents = &ffapi.Route{
	Name:       "getEvents",
	Path:       "events",
	Method:     http.MethodGet,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchreferences", Example: "true", Description: coremsgs.APIParamsFetchReferences, IsBool: true},
		{Name: "fetchreference", Example: "true", Description: coremsgs.APIParamsFetchReference, IsBool: true},
	},
	FilterFactory:   database.EventQueryFactory,
	Description:     coremsgs.APIEndpointsGetEvents,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Event{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchreferences"], "true") || strings.EqualFold(r.QP["fetchreference"], "true") {
				return r.FilterResult(cr.or.GetEventsWithReferences(cr.ctx, r.Filter))
			}
			return r.FilterResult(cr.or.GetEvents(cr.ctx, r.Filter))
		},
	},
}
