// SiGG-GoLang-On-the-Fly //
package runners

import (
	"testing"

	"github.com/hyperledger/firefly/test/e2e/gateway"
	"github.com/stretchr/testify/suite"
)

func TestFabricGatewayE2ESuite(t *testing.T) {
	suite.Run(t, new(gateway.FabricContractTestSuite))
}
