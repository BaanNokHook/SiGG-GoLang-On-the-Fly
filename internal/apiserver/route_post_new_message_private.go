// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNewMessagePrivate = &ffapi.Route{
	Name:       "postNewMessagePrivate",
	Path:       "messages/private",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true},
	},
	Description:     coremsgs.APIEndpointsPostNewMessagePrivate,
	JSONInputValue:  func() interface{} { return &core.MessageInOut{} },
	JSONOutputValue: func() interface{} { return &core.Message{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.MultiParty() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			return cr.or.PrivateMessaging().SendMessage(cr.ctx, r.Input.(*core.MessageInOut), waitConfirm)
		},
	},
}
