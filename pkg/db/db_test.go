package db

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DBTestSuite struct {
	suite.Suite
}

func (suite *DBTestSuite) SetupTest() {

}

func (suite *DBTestSuite) TestInitDB() {
	require := suite.Require()

	_, err := InitDB()
	require.NoError(err)
}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}
