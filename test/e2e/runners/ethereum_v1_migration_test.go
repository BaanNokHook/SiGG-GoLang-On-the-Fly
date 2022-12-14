// SiGG-GoLang-On-the-Fly //
package runners

import (
	"testing"

	"github.com/hyperledger/firefly/test/e2e/multiparty"
	"github.com/stretchr/testify/suite"
)

// This suite can only be run once per stack.
// Specifically, it must be run on a stack that has never used a V1 FireFly multiparty contract.
// The test deploys a V1 contract, then migrates to V2. This is a one-time operation that cannot
// be performed again on the same stack.
func TestEthereumV1MigrationE2ESuite(t *testing.T) {
	suite.Run(t, new(multiparty.ContractMigrationV1TestSuite))
}
