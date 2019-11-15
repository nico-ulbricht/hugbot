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
	SenderID    uuid.UUID
	Amount      int
	Type        string
}
