// +Build integration
package integration_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ExampleTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.VariableThatShouldStartAtFive = 5
}

// This will run before before the tests in the suite are run
func (suite *ExampleTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
	fmt.Println("TestExample")
	suite.Equal(suite.VariableThatShouldStartAtFive, 5)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	fmt.Println("TestExampleTestSuite")
	suite.Run(t, new(ExampleTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
// func (suite *ExampleTestSuite) BeforeTest(suiteName, testName string) {
// 	fmt.Println("BeforeTest")
// }

// This will run after test finishes
// and receives the suite and test names as input
// func (suite *ExampleTestSuite) AfterTest(suiteName, testName string) {
// 	fmt.Println("AfterTest")
// }
