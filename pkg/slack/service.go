package slack

import (
	"context"
	"regexp"

	"github.com/google/uuid"
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
	Message     string
	ReferenceID string
	SenderID    string
}

type handleReactionInput struct {
	RecipientID string
	ReferenceID string
	SenderID    string
	Type        string
}

type service struct {
	reactionService reaction.Service
	slackClient     *slack.Client
	userService     user.Service
}

var reactionRegexp = regexp.MustCompile(":(\\w+):")
var userRegexp = regexp.MustCompile("\\<@(.*?)\\>")

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
	sender, err := svc.userService.Upsert(ctx, user.UpsertInput{ExternalID: input.SenderID})
	if err != nil {
		return errors.WithStack(err)
	}

	amountsByReaction := map[string]int{}
	reactionMatches := reactionRegexp.FindAllStringSubmatch(input.Message, -1)
	for _, aReactionMatch := range reactionMatches {
		amountsByReaction[aReactionMatch[1]]++
	}

	recipientMatches := userRegexp.FindAllStringSubmatch(input.Message, -1)
	recipientIDs := make([]uuid.UUID, len(recipientMatches))
	for idx, aMatch := range recipientMatches {
		recipientID := aMatch[1]
		recipient, err := svc.userService.Upsert(ctx, user.UpsertInput{ExternalID: recipientID})
		if err != nil {
			return errors.WithStack(err)
		}

		recipientIDs[idx] = recipient.ID
	}

	for aReaction, reactionAmount := range amountsByReaction {
		for _, aRecipientID := range recipientIDs {
			_, err = svc.reactionService.Create(ctx, reaction.CreateInput{
				RecipientID: aRecipientID,
				ReferenceID: input.ReferenceID,
				SenderID:    sender.ID,
				Amount:      reactionAmount,
				Type:        aReaction,
			})
		}
	}

	return nil
}

func (svc *service) HandleReaction(ctx context.Context, input handleReactionInput) error {
	sender, err := svc.userService.Upsert(ctx, user.UpsertInput{ExternalID: input.SenderID})
	if err != nil {
		return errors.WithStack(err)
	}

	recipient, err := svc.userService.Upsert(ctx, user.UpsertInput{ExternalID: input.RecipientID})
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = svc.reactionService.Create(ctx, reaction.CreateInput{
		RecipientID: recipient.ID,
		ReferenceID: input.ReferenceID,
		SenderID:    sender.ID,
		Amount:      1,
		Type:        input.Type,
	})

	return errors.WithStack(err)
}

type config struct {
	Token string `envconfig:"SLACK_TOKEN" required:"true"`
}

func NewService(
	reactionService reaction.Service,
	userService user.Service,
) Service {
	var c config
	envconfig.MustProcess("", &c)

	slackClient := slack.New(c.Token)
	return &service{
		reactionService: reactionService,
		slackClient:     slackClient,
		userService:     userService,
	}
}
