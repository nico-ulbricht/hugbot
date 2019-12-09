package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	Upsert(ctx context.Context, input UpsertInput) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
}

type UpsertInput struct {
	ExternalID string
	Name       string
}

type service struct {
	userRepository Repository
}

func (svc *service) Upsert(ctx context.Context, input UpsertInput) (*User, error) {
	existingUser, err := svc.userRepository.GetByExternalID(ctx, input.ExternalID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if existingUser == nil {
		return existingUser, nil
	}

	user := newUser(input.ExternalID)
	return svc.userRepository.Insert(ctx, user)
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

func NewService(userRepository Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}
