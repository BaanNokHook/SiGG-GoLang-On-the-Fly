// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNewMessageRequestReply = &ffapi.Route{
	Name:            "postNewMessageRequestReply",
	Path:            "messages/requestreply",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostNewMessageRequestReply,
	JSONInputValue:  func() interface{} { return &core.MessageInOut{} },
	JSONOutputValue: func() interface{} { return &core.MessageInOut{} },
	JSONOutputCodes: []int{http.StatusOK}, // Sync operation
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.MultiParty() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.RequestReply(cr.ctx, r.Input.(*core.MessageInOut))
			return output, err
		},
	},
}
