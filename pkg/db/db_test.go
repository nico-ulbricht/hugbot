package db_test

import (
	"testing"

	. "github.com/nico-ulbricht/hugbot/pkg/db"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
}

func (suite *DatabaseSuite) TestNewWithConfig() {
	suite.Run("should return valid database instance", func() {
		instance := NewWithConfig(Config{
			Database: "hugbot",
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",

			MaxConnectionRetries: 0,
		})

		require.NotNil(suite.T(), instance)
	})
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, &DatabaseSuite{})
}
