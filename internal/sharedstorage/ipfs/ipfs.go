// SiGG-GoLang-On-the-Fly //
package ipfs

import (
	"context"
	"encoding/json"
	"fmt"

	"io"

	"github.com/go-resty/resty/v2"
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/ffresty"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/pkg/sharedstorage"
)

type IPFS struct {
	ctx          context.Context
	capabilities *sharedstorage.Capabilities
	apiClient    *resty.Client
	gwClient     *resty.Client
}

type ipfsUploadResponse struct {
	Name string      `json:"Name"`
	Hash string      `json:"Hash"`
	Size json.Number `json:"Size"`
}

func (i *IPFS) Name() string {
	return "ipfs"
}

func (i *IPFS) Init(ctx context.Context, config config.Section) error {

	i.ctx = log.WithLogField(ctx, "sharedstorage", "ipfs")

	apiConfig := config.SubSection(IPFSConfAPISubconf)
	if apiConfig.GetString(ffresty.HTTPConfigURL) == "" {
		return i18n.NewError(ctx, coremsgs.MsgMissingPluginConfig, apiConfig.Resolve(ffresty.HTTPConfigURL), "ipfs")
	}
	i.apiClient = ffresty.New(i.ctx, apiConfig)
	gwConfig := config.SubSection(IPFSConfGatewaySubconf)
	if gwConfig.GetString(ffresty.HTTPConfigURL) == "" {
		return i18n.NewError(ctx, coremsgs.MsgMissingPluginConfig, gwConfig.Resolve(ffresty.HTTPConfigURL), "ipfs")
	}
	i.gwClient = ffresty.New(i.ctx, gwConfig)
	i.capabilities = &sharedstorage.Capabilities{}
	return nil
}

func (i *IPFS) SetHandler(namespace string, handler sharedstorage.Callbacks) {
}

func (i *IPFS) Capabilities() *sharedstorage.Capabilities {
	return i.capabilities
}

func (i *IPFS) UploadData(ctx context.Context, data io.Reader) (string, error) {
	var ipfsResponse ipfsUploadResponse
	res, err := i.apiClient.R().
		SetContext(ctx).
		SetFileReader("document", "file.bin", data).
		SetResult(&ipfsResponse).
		Post("/api/v0/add")
	if err != nil || !res.IsSuccess() {
		return "", ffresty.WrapRestErr(i.ctx, res, err, coremsgs.MsgIPFSRESTErr)
	}
	log.L(ctx).Infof("IPFS published %s Size=%s", ipfsResponse.Hash, ipfsResponse.Size)
	return ipfsResponse.Hash, err
}

func (i *IPFS) DownloadData(ctx context.Context, payloadRef string) (data io.ReadCloser, err error) {
	res, err := i.gwClient.R().
		SetContext(ctx).
		SetDoNotParseResponse(true).
		Get(fmt.Sprintf("/ipfs/%s", payloadRef))
	ffresty.OnAfterResponse(i.gwClient, res) // required using SetDoNotParseResponse
	if err != nil || !res.IsSuccess() {
		if res != nil && res.RawBody() != nil {
			_ = res.RawBody().Close()
		}
		return nil, ffresty.WrapRestErr(i.ctx, res, err, coremsgs.MsgIPFSRESTErr)
	}
	log.L(ctx).Infof("IPFS retrieved %s", payloadRef)
	return res.RawBody(), nil
}
