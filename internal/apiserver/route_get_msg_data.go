package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getMsgData = &ffapi.Route{
	Name:   "getMsgData",
	Path:   "messages/{msgid}/data",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "msgid", Description: coremsgs.APIParamsMessageID},
	},
	QueryParams:     nil,
	FilterFactory:   nil, // No filtering on this route - use data
	Description:     coremsgs.APIEndpointsGetMsgData,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return core.DataArray{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetMessageData(cr.ctx, r.PP["msgid"])
			return output, err
		},
	},
}
