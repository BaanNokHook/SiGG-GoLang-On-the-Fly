// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postOpRetry = &ffapi.Route{
	Name:   "postOpRetry",
	Path:   "operations/{opid}/retry",
	Method: http.MethodPost,
	PathParams: []*ffapi.PathParam{
		{Name: "opid", Description: coremsgs.OperationID},
	},
	QueryParams:     []*ffapi.QueryParam{},
	Description:     coremsgs.APIEndpointsPostOpRetry,
	JSONInputValue:  func() interface{} { return &core.EmptyInput{} },
	JSONOutputValue: func() interface{} { return &core.Operation{} },
	JSONOutputCodes: []int{http.StatusAccepted},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			opid, err := fftypes.ParseUUID(cr.ctx, r.PP["opid"])
			if err != nil {
				return nil, err
			}
			return cr.or.Operations().RetryOperation(cr.ctx, opid)
		},
	},
}
