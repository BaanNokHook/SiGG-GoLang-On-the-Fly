// SiGG-GoLang-On-the-Fly //
package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getTokenApprovals = &ffapi.Route{
	Name:            "getTokenApprovals",
	Path:            "tokens/approvals",
	Method:          http.MethodGet,
	PathParams:      nil,
	FilterFactory:   database.TokenApprovalQueryFactory,
	Description:     coremsgs.APIEndpointsGetTokenApprovals,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []*core.TokenApproval{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			filter := r.Filter
			return r.FilterResult(cr.or.Assets().GetTokenApprovals(cr.ctx, filter))
		},
	},
}
