package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getGroups = &ffapi.Route{
	Name:            "getGroups",
	Path:            "groups",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.GroupQueryFactory,
	Description:     coremsgs.APIEndpointsGetGroups,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Group{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.PrivateMessaging() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.PrivateMessaging().GetGroups(cr.ctx, r.Filter))
		},
	},
}
