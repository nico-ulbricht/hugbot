package reaction

import (
	"context"
	"time"

	"github.com/nico-ulbricht/hugbot/pkg/event"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Reaction, error)
	GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error)
	GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error)
}

type CreateInput struct {
	RecipientID uuid.UUID
	SenderID    uuid.UUID
	ReferenceID string
	Amount      int
	Type        string
}

type service struct {
	config             configuration
	reactionPublisher  event.Publisher
	reactionRepository Repository
}

func (svc service) Create(ctx context.Context, input CreateInput) (*Reaction, error) {
	existingReaction, err := svc.reactionRepository.GetByReferenceIDAndType(ctx, input.ReferenceID, input.Type)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if existingReaction != nil {
		return existingReaction, nil
	}

	var isSupported bool
	for _, aSupportedType := range svc.config.SupportedTypes {
		if aSupportedType == input.Type {
			isSupported = true
			break
		}
	}

	if isSupported == false {
		return nil, nil
	}

	reaction := newReaction(
		input.RecipientID,
		input.SenderID,
		input.ReferenceID,
		input.Amount,
		input.Type,
	)

	reaction, err = svc.reactionRepository.Insert(ctx, reaction)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	reactionCreatedEvent := event.ReactionCreated{
		Meta: event.Meta{
			Created: time.Now().UTC(),
			Type:    event.ReactionCreatedType,
		},
		Payload: event.ReactionCreatedPayload{
			Amount:      reaction.Amount,
			RecipientID: reaction.RecipientID,
			SenderID:    reaction.SenderID,
			Type:        reaction.Type,
		},
	}

	err = svc.reactionPublisher.Publish(ctx, reactionCreatedEvent)
	return reaction, errors.WithStack(err)
}

func (svc service) GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error) {
	return svc.reactionRepository.GetByRecipientID(ctx, recipientID)
}

func (svc service) GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error) {
	return svc.reactionRepository.GetBySenderID(ctx, senderID)
}

type configuration struct {
	SupportedTypes []string `envconfig:"REACTION_SUPPORTED_TYPES"`
}

func NewService(
	reactionPublisher event.Publisher,
	reactionRepository Repository,
) Service {
	var config configuration
	envconfig.MustProcess("", &config)

	return &service{
		config:             config,
		reactionPublisher:  reactionPublisher,
		reactionRepository: reactionRepository,
	}
}
