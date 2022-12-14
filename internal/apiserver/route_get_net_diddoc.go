package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/networkmap"
)

var getNetworkDIDDocByDID = &ffapi.Route{
	Name:   "getNetworkDIDDocByDID",
	Path:   "network/diddocs/{did:.+}",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "did", Description: coremsgs.APIParamsDID},
	},
	Description:     coremsgs.APIEndpointsGetDIDDocByDID,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &networkmap.DIDDocument{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			return cr.or.NetworkMap().GetDIDDocForIndentityByDID(cr.ctx, r.PP["did"])
		},
	},
}
