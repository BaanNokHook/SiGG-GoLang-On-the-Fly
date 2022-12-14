// SiGG-GoLang-On-the-Fly //
package multiparty

import (
	"github.com/hyperledger/firefly/test/e2e"
	"github.com/hyperledger/firefly/test/e2e/client"
)

type ContractMigrationV1TestSuite struct {
	ContractMigrationTestSuite
}

func (suite *ContractMigrationV1TestSuite) SetupSuite() {
	suite.testState = beforeE2ETest(suite.T())
}

func (suite *ContractMigrationV1TestSuite) BeforeTest(suiteName, testName string) {
	suite.testState = beforeE2ETest(suite.T())
}

func (suite *ContractMigrationV1TestSuite) AfterTest(suiteName, testName string) {
	e2e.VerifyAllOperationsSucceeded(suite.T(), []*client.FireFlyClient{suite.testState.client1, suite.testState.client2}, suite.testState.startTime)
}

func (suite *ContractMigrationV1TestSuite) TestContractMigration() {
	defer suite.testState.Done()

	address1 := deployContract(suite.T(), suite.testState.stackName, "firefly/FireflyV1.json")
	address2 := deployContract(suite.T(), suite.testState.stackName, "firefly/Firefly.json")
	runMigrationTest(&suite.ContractMigrationTestSuite, address1, address2, true)
}
