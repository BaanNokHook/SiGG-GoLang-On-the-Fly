// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getStatus = &ffapi.Route{
	Name:            "getStatus",
	Path:            "status",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetStatus,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.NamespaceStatus{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetStatus(cr.ctx)
			return output, err
		},
	},
}
