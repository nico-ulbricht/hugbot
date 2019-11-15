package user

import "context"

type Service interface {
	Create(ctx context.Context, input CreateInput) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByExternalID(ctx context.Context, externalID string) (*User, error)
}

type CreateInput struct {
	ExternalID string
	Name       string
}
