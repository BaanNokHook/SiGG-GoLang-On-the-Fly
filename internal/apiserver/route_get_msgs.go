// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getMsgs = &ffapi.Route{
	Name:       "getMsgs",
	Path:       "messages",
	Method:     http.MethodGet,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchdata", IsBool: true, Description: coremsgs.APIFetchDataDesc},
	},
	FilterFactory:   database.MessageQueryFactory,
	Description:     coremsgs.APIEndpointsGetMsgs,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Message{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["fetchdata"], "true") {
				return r.FilterResult(cr.or.GetMessagesWithData(cr.ctx, r.Filter))
			}
			return r.FilterResult(cr.or.GetMessages(cr.ctx, r.Filter))
		},
	},
}
