package reaction

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*Reaction, error)
	GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error)
	GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error)
}

type CreateInput struct {
	RecipientID uuid.UUID
	ReferenceID string
	SenderID    uuid.UUID
	Amount      int
	Type        string
}

type service struct {
	reactionRepository Repository
}

func (svc service) Create(ctx context.Context, input CreateInput) (*Reaction, error) {
	panic("TODO")
}

func (svc service) GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error) {
	panic("TODO")
}

func (svc service) GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error) {
	panic("TODO")
}

func NewService(reactionRepository Repository) Service {
	return &service{
		reactionRepository: reactionRepository,
	}
}
