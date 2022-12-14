// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/batch"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
)

var getStatusBatchManager = &ffapi.Route{
	Name:            "getStatusBatchManager",
	Path:            "status/batchmanager",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetStatusBatchManager,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &batch.ManagerStatus{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.BatchManager() != nil
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.BatchManager().Status(), nil
		},
	},
}
