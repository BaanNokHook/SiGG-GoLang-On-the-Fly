// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postPinsRewind = &ffapi.Route{
	Name:            "postPinsRewind",
	Path:            "pins/rewind",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostPinsRewind,
	JSONInputValue:  func() interface{} { return &core.PinRewind{} },
	JSONOutputValue: func() interface{} { return &core.PinRewind{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.RewindPins(cr.ctx, r.Input.(*core.PinRewind))
		},
	},
}
