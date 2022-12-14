// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getMsgByID = &ffapi.Route{
	Name:   "getMsgByID",
	Path:   "messages/{msgid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "msgid", Description: coremsgs.APIParamsMessageID},
	},
	QueryParams: []*ffapi.QueryParam{
		{Name: "fetchdata", IsBool: true, Description: coremsgs.APIFetchDataDesc},
	},
	Description:     coremsgs.APIEndpointsGetMsgByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.MessageInOut{} }, // can include full values
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			if strings.EqualFold(r.QP["data"], "true") || strings.EqualFold(r.QP["fetchdata"], "true") {
				return cr.or.GetMessageByIDWithData(cr.ctx, r.PP["msgid"])
			}
			return cr.or.GetMessageByID(cr.ctx, r.PP["msgid"])
		},
	},
}
