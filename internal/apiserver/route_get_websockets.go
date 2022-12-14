// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/events/eifactory"
	"github.com/hyperledger/firefly/internal/events/websockets"
	"github.com/hyperledger/firefly/pkg/core"
)

var getWebSockets = &ffapi.Route{
	Name:            "getWebSockets",
	Path:            "websockets",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetWebSockets,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.WebSocketStatus{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			ws, _ := eifactory.GetPlugin(cr.ctx, "websockets")
			return ws.(*websockets.WebSockets).GetStatus(), nil
		},
	},
}
