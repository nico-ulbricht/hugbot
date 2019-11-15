package slack

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"

	"github.com/nico-ulbricht/hugbot/pkg/reaction"
	"github.com/nico-ulbricht/hugbot/pkg/user"
)

type Service interface {
	SendMessage(ctx context.Context, msg *Message) error
	HandleMessage(ctx context.Context, input handleMessageInput) error
	HandleReaction(ctx context.Context, input handleReactionInput) error
}

type handleMessageInput struct {
	Message  string
	SenderID string
}

type handleReactionInput struct {
	SenderID    string
	RecipientID string
	Type        string
}

type service struct {
	slackClient     *slack.Client
	reactionService reaction.Service
	userService     user.Service
}

func (svc *service) SendMessage(ctx context.Context, msg *Message) error {
	usr, err := svc.userService.GetByID(ctx, msg.RecipientID)
	if err != nil {
		return errors.WithStack(err)
	}

	_, _, channelID, err := svc.slackClient.OpenIMChannel(usr.ExternalID)
	if err != nil {
		return errors.WithStack(err)
	}

	_, _, _, err = svc.slackClient.SendMessage(channelID, slack.MsgOptionText(msg.Message, false))
	return errors.WithStack(err)
}

func (svc *service) HandleMessage(ctx context.Context, input handleMessageInput) error {
	panic("TODO")
}

func (svc *service) HandleReaction(ctx context.Context, input handleReactionInput) error {
	panic("TODO")
}

type config struct {
	Token string `envconfig:"SLACK_TOKEN"`
}

func NewService() Service {
	var c config
	envconfig.MustProcess("", &c)

	slackClient := slack.New(c.Token)
	return &service{
		slackClient: slackClient,
	}
}
