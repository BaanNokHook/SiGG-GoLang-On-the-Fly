// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getPins = &ffapi.Route{
	Name:            "getPins",
	Path:            "pins",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.PinQueryFactory,
	Description:     coremsgs.APIEndpointsGetPins,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []core.Pin{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetPins(cr.ctx, r.Filter))
		},
	},
}
