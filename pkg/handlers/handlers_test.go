package handlers

import (
	"testing"

	"github.com/carsonclarke570/lair-api/pkg/db"
	"github.com/stretchr/testify/suite"
	"gopkg.in/h2non/gock.v1"
	"upper.io/db.v3/lib/sqlbuilder"
)

type HandlerTestSuite struct {
	suite.Suite
	sess sqlbuilder.Database
}

const (
	host = "localhost"
)

func (suite *HandlerTestSuite) SetupTest() {
	require := suite.Require()

	var err error
	suite.sess, err = db.InitDB()
	require.NoError(err)
}

func (suite *HandlerTestSuite) TearDownTest() {
	suite.sess.Close()
}

func (suite *HandlerTestSuite) TearDownSuite() {
	gock.Off()
}

func (suite *HandlerTestSuite) TestCreate() {

}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
