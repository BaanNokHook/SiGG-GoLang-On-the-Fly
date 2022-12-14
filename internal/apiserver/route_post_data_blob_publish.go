// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
)

var postDataBlobPublish = &ffapi.Route{
	Name:   "postDataBlobPublish",
	Path:   "data/{dataid}/blob/publish",
	Method: http.MethodPost,
	PathParams: []*ffapi.PathParam{
		{Name: "dataid", Description: coremsgs.APIParamsBlobID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsPostDataBlobPublish,
	JSONInputValue:  func() interface{} { return &core.PublishInput{} },
	JSONOutputValue: func() interface{} { return &core.Data{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Broadcast() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.Broadcast().PublishDataBlob(cr.ctx, r.PP["dataid"], r.Input.(*core.PublishInput).IdempotencyKey)
		},
	},
}
