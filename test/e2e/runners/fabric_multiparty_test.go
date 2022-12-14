// SiGG-GoLang-On-the-Fly //
package runners

import (
	"testing"

	"github.com/hyperledger/firefly/test/e2e/multiparty"
	"github.com/stretchr/testify/suite"
)

func TestFabricMultipartyE2ESuite(t *testing.T) {
	suite.Run(t, new(multiparty.OnChainOffChainTestSuite))
	suite.Run(t, new(multiparty.IdentityTestSuite))
	suite.Run(t, new(multiparty.FabricContractTestSuite))
}
