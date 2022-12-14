package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getDataByID = &ffapi.Route{
	Name:   "getDataByID",
	Path:   "data/{dataid}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "dataid", Description: coremsgs.APIParamsDataID},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetDataByID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Data{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetDataByID(cr.ctx, r.PP["dataid"])
			return output, err
		},
	},
}
