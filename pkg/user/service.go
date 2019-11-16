package user

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, input CreateInput) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByExternalID(ctx context.Context, externalID string) (*User, error)
}

type CreateInput struct {
	ExternalID string
	Name       string
}

type service struct {
	userRepository Repository
}

func (svc *service)	Create(ctx context.Context, input CreateInput) (*User, error) {
	panic("TODO")
}

func (svc *service)	GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	panic("TODO")
}

func (svc *service)	GetByExternalID(ctx context.Context, externalID string) (*User, error) {
	panic("TODO")
}

func NewService(userRepository Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}
