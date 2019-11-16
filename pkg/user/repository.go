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

type repository struct{}

func (rp *repository)	Insert(ctx context.Context, input CreateInput) (*User, error) {
	panic("TODO")
}

func (rp *repository)	GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	panic("TODO")
}

func (rp *repository)	GetByExternalID(ctx context.Context, externalID string) (*User, error) {
	panic("TODO")
}

func NewRepository() Repository {
	return &repository{}
}
