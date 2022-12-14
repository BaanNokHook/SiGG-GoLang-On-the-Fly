package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getDatatypes = &ffapi.Route{
	Name:            "getDatatypes",
	Path:            "datatypes",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   database.DatatypeQueryFactory,
	Description:     coremsgs.APIEndpointsGetDatatypes,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.Datatype{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return r.FilterResult(cr.or.GetDatatypes(cr.ctx, r.Filter))
		},
	},
}
