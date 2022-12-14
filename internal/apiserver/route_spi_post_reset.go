// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
)

var spiPostReset = &ffapi.Route{
	Name:            "spiPostReset",
	Path:            "reset",
	Method:          http.MethodPost,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   nil,
	Description:     coremsgs.APIEndpointsAdminPostReset,
	JSONInputValue:  nil,
	JSONOutputValue: nil,
	JSONOutputCodes: []int{http.StatusNoContent},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			cr.mgr.Reset(cr.ctx)
			return nil, nil
		},
	},
}
