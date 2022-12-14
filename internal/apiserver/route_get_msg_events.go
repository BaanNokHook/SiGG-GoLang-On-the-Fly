// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getMsgEvents = &ffapi.Route{
	Name:   "getMsgEvents",
	Path:   "messages/{msgid}/events",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "msgid", Description: coremsgs.APIParamsMessageID},
	},
	QueryParams:     nil,
	FilterFactory:   database.EventQueryFactory,
	Description:     coremsgs.APIEndpointsGetMsgEvents,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Event{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetMessageEvents(cr.ctx, r.PP["msgid"], r.Filter))
		},
	},
}
