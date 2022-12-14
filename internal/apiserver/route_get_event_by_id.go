package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getEventByID = &ffapi.Route{
	Name:   "getEventByID",
	Path:   "events/{eid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "eid", Description: coremsgs.APIParamsEventID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchreference", Example: "true", Description: coremsgs.APIParamsFetchReference, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsGetEventByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Event{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchreference"], "true") {
				return cr.or.GetEventByIDWithReference(cr.ctx, r.PP["eid"])
			}
			return cr.or.GetEventByID(cr.ctx, r.PP["eid"])
		},
	},
}
