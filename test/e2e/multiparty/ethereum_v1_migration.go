// SiGG-GoLang-On-the-Fly //
package multiparty

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestEthereumMultipartyV1MigrationE2ESuite(t *testing.T) {
	suite.Run(t, new(ContractMigrationV1TestSuite))
}
