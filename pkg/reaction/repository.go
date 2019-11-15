package reaction

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Insert(ctx context.Context, reaction *Reaction) (*Reaction, error)
	GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error)
	GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error)
}
