package reaction

import (
	"context"

	"github.com/nico-ulbricht/hugbot/pkg/event"

	"github.com/google/uuid"
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

	// TODO: validate type
	reaction := newReaction(
		input.RecipientID,
		input.SenderID,
		input.ReferenceID,
		input.Amount,
		input.Type,
	)

	return svc.reactionRepository.Insert(ctx, reaction)
}

func (svc service) GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error) {
	return svc.reactionRepository.GetByRecipientID(ctx, recipientID)
}

func (svc service) GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error) {
	return svc.reactionRepository.GetBySenderID(ctx, senderID)
}

func NewService(
	reactionPublisher event.Publisher,
	reactionRepository Repository,
) Service {
	return &service{
		reactionPublisher:  reactionPublisher,
		reactionRepository: reactionRepository,
	}
}
