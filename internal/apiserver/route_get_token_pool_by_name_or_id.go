// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getTokenPoolByNameOrID = &ffapi.Route{
	Name:   "getTokenPoolByNameOrID",
	Path:   "tokens/pools/{nameOrId}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "nameOrId", Description: coremsgs.APIParamsTokenPoolNameOrID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetTokenPoolByNameOrID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.TokenPool{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.Assets().GetTokenPoolByNameOrID(cr.ctx, r.PP["nameOrId"])
			return output, err
		},
	},
}
