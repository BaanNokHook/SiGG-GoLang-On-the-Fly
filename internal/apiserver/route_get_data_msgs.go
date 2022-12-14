package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getDataMsgs = &ffapi.Route{
	Name:   "getDataMsgs",
	Path:   "data/{dataid}/messages",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "dataid", Description: coremsgs.APIParamsDataID},
	},
	QueryParams:     nil,
	FilterFactory:   database.MessageQueryFactory,
	Description:     coremsgs.APIEndpointsGetDataMsgs,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Message{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.MultiParty() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetMessagesForData(cr.ctx, r.PP["dataid"], r.Filter))
		},
	},
}
