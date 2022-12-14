// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
)

var postNewDatatype = &ffapi.Route{
	Name:       "postNewDatatype",
	Path:       "datatypes",
	Method:     http.MethodPost,
	PathParams: nil,
	QueryParams: []*ffapi.QueryParam{
		{Name: "confirm", Description: coremsgs.APIConfirmQueryParam, IsBool: true, Example: "true"},
	},
	Description:     coremsgs.APIEndpointsPostNewDatatype,
	JSONInputValue:  func() interface{} { return &core.Datatype{} },
	JSONOutputValue: func() interface{} { return &core.Datatype{} },
	JSONOutputCodes: []int{http.StatusAccepted, http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			waitConfirm := strings.EqualFold(r.QP["confirm"], "true")
			r.SuccessStatus = syncRetcode(waitConfirm)
			err = cr.or.DefinitionSender().DefineDatatype(cr.ctx, r.Input.(*core.Datatype), waitConfirm)
			return r.Input, err
		},
	},
}
