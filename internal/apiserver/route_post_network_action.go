// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNetworkAction = &ffapi.Route{
	Name:            "postNetworkAction",
	Path:            "network/action",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostNetworkAction,
	JSONInputValue:  func() interface{} { return &core.NetworkAction{} },
	JSONOutputValue: func() interface{} { return &core.NetworkAction{} },
	JSONOutputCodes: []int{http.StatusAccepted},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.MultiParty() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			err = cr.or.SubmitNetworkAction(cr.ctx, r.Input.(*core.NetworkAction))
			return r.Input, err
		},
	},
}
