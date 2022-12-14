package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getBatchByID = &ffapi.Route{
	Name:   "getBatchByID",
	Path:   "batches/{batchid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "batchid", Description: coremsgs.APIParamsBatchID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetBatchBbyID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.BatchPersisted{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetBatchByID(cr.ctx, r.PP["batchid"])
			return output, err
		},
	},
}
