package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (svc *service) Create(ctx context.Context, input CreateInput) (*User, error) {
	panic("TODO")
}

func (svc *service) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	user, err := svc.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if user == nil {
		return nil, ErrNotFound{UserID: userID}
	}

	return user, nil
}

func (svc *service) GetByExternalID(ctx context.Context, externalID string) (*User, error) {
	user, err := svc.userRepository.GetByExternalID(ctx, externalID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if user == nil {
		return nil, ErrNotFound{ExternalID: externalID}
	}

	return user, nil
}

func NewService(userRepository Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}
