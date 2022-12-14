package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var getDatatypeByName = &ffapi.Route{
	Name:   "getDatatypeByName",
	Path:   "datatypes/{name}/{version}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "name", Description: coremsgs.APIParamsDatatypeName},
		{Name: "version", Description: coremsgs.APIParamsDatatypeVersion},
	},
	QueryParams:     nil,
	Description:     coremsgs.APIEndpointsGetDatatypeByName,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &core.Datatype{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			output, err = cr.or.GetDatatypeByName(cr.ctx, r.PP["name"], r.PP["version"])
			return output, err
		},
	},
}
