package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Insert(ctx context.Context, input CreateInput) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByExternalID(ctx context.Context, externalID string) (*User, error)
}
