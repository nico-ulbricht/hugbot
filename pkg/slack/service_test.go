package slack_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	. "github.com/nico-ulbricht/hugbot/pkg/slack"
)

type ServiceSuite struct {
	suite.Suite

	slackService Service
}

func (suite *ServiceSuite) SetupSuite() {
	suite.slackService = NewService(nil, nil)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, &ServiceSuite{})
}
