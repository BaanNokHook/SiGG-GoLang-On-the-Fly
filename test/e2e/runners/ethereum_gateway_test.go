// SiGG-GoLang-On-the-Fly //
package runners

import (
	"testing"

	"github.com/hyperledger/firefly/test/e2e/gateway"
	"github.com/stretchr/testify/suite"
)

func TestEthereumGatewayE2ESuite(t *testing.T) {
	suite.Run(t, new(gateway.TokensTestSuite))
	suite.Run(t, new(gateway.EthereumContractTestSuite))
	suite.Run(t, new(gateway.TokensOnlyTestSuite))
}
