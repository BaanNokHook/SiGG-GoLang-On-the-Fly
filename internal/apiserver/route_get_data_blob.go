package apiserver

import (
	"net/http"
	"strconv"

	"github.com/hyperledger/firefly-common/pkg/ffapi"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/orchestrator"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

var getDataBlob = &ffapi.Route{
	Name:   "getDataBlob",
	Path:   "data/{dataid}/blob",
	Method: http.MethodGet,
	PathParams: []*ffapi.PathParam{
		{Name: "dataid", Description: coremsgs.APIParamsBlobID},
	},
	QueryParams:     nil,
	FilterFactory:   database.MessageQueryFactory,
	Description:     coremsgs.APIEndpointsGetDataBlob,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return []byte{} },
	JSONOutputCodes: []int{http.StatusOK},
	Extensions: &coreExtensions{
		EnabledIf: func(or orchestrator.Orchestrator) bool {
			return or.Data().BlobsEnabled()
		},
		CoreJSONHandler: func(r *ffapi.APIRequest, cr *coreRequest) (output interface{}, err error) {
			blob, reader, err := cr.or.Data().DownloadBlob(cr.ctx, r.PP["dataid"])
			if err == nil {
				r.ResponseHeaders.Set(core.HTTPHeadersBlobHashSHA256, blob.Hash.String())
				if blob.Size > 0 {
					r.ResponseHeaders.Set(core.HTTPHeadersBlobSize, strconv.FormatInt(blob.Size, 10))
				}
			}
			return reader, nil
		},
	},
}
