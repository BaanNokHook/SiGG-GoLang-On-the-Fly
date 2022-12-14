// SiGG-GoLang-On-the-Fly //
package fabric

import (
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/wsclient"
)

const (
	defaultBatchSize    = 50
	defaultBatchTimeout = 500
	defaultPrefixShort  = "fly"
	defaultPrefixLong   = "firefly"
)

const (
	// FabconnectConfigKey is a sub-key in the config to contain all the ethconnect specific config,
	FabconnectConfigKey = "fabconnect"

	// FabconnectConfigDefaultChannel is the default Fabric channel to use if no "ledger" is specified in requests
	FabconnectConfigDefaultChannel = "channel"
	// FabconnectConfigSigner is the signer identity used to subscribe to FireFly chaincode events
	FabconnectConfigSigner = "signer"
	// FabconnectConfigTopic is the websocket listen topic that the node should register on, which is important if there are multiple
	// nodes using a single fabconnect
	FabconnectConfigTopic = "topic"
	// FabconnectConfigBatchSize is the batch size to configure on event streams, when auto-defining them
	FabconnectConfigBatchSize = "batchSize"
	// FabconnectConfigBatchTimeout is the batch timeout to configure on event streams, when auto-defining them
	FabconnectConfigBatchTimeout = "batchTimeout"
	// FabconnectPrefixShort is used in the query string in requests to ethconnect
	FabconnectPrefixShort = "prefixShort"
	// FabconnectPrefixLong is used in HTTP headers in requests to ethconnect
	FabconnectPrefixLong = "prefixLong"
	// FabconnectConfigChaincodeDeprecated is the Fabric Firefly chaincode deployed to the Firefly channels
	FabconnectConfigChaincodeDeprecated = "chaincode"
)

func (f *Fabric) InitConfig(config config.Section) {
	f.fabconnectConf = config.SubSection(FabconnectConfigKey)
	wsclient.InitConfig(f.fabconnectConf)
	f.fabconnectConf.AddKnownKey(FabconnectConfigDefaultChannel)
	f.fabconnectConf.AddKnownKey(FabconnectConfigChaincodeDeprecated)
	f.fabconnectConf.AddKnownKey(FabconnectConfigSigner)
	f.fabconnectConf.AddKnownKey(FabconnectConfigTopic)
	f.fabconnectConf.AddKnownKey(FabconnectConfigBatchSize, defaultBatchSize)
	f.fabconnectConf.AddKnownKey(FabconnectConfigBatchTimeout, defaultBatchTimeout)
	f.fabconnectConf.AddKnownKey(FabconnectPrefixShort, defaultPrefixShort)
	f.fabconnectConf.AddKnownKey(FabconnectPrefixLong, defaultPrefixLong)
}
