package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/database"
)

var getDataValue = &ffapi.Route{
	Name:   "getDataValue",
	Path:   "data/{dataid}/value",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "dataid", Description: coremsgs.APIParamsBlobID},
	},
	QueryParams:     nil,
	FilterFactory:   database.MessageQueryFactory,
	Description:     coremsgs.APIEndpointsGetDataValue,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []byte{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			d, err := cr.or.GetDataByID(cr.ctx, r.PP["dataid"])
			if err != nil {
				return nil, err
			}
			return d.Value, err
		},
	},
}
