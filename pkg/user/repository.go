package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByExternalID(ctx context.Context, externalID string) (*User, error)
}

type repository struct {
	psql *sqlx.DB
}

func (rp *repository) Insert(ctx context.Context, user *User) (*User, error) {
	stmt, err := rp.psql.PreparexContext(ctx, `
		insert into users (
			id,
			external_id,
			created_at,
			updated_at
		) values ($1, $2, $3, $4)
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	now := time.Now().UTC()
	_, err = stmt.ExecContext(
		ctx,
		user.ID,
		user.ExternalID,
		now,
		now,
	)

	return user, errors.WithStack(err)
}

func (rp *repository) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	stmt, err := rp.psql.PreparexContext(ctx, `
		select
			id,
			external_id
		from users
		where id = $1
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	var user User
	err = stmt.GetContext(ctx, &user, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, errors.WithStack(err)
}

func (rp *repository) GetByExternalID(ctx context.Context, externalID string) (*User, error) {
	stmt, err := rp.psql.PreparexContext(ctx, `
		select
			id,
			external_id
		from users
		where external_id = $1
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	var user User
	err = stmt.GetContext(ctx, &user, externalID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, errors.WithStack(err)
}

func NewRepository(psql *sqlx.DB) Repository {
	return &repository{psql: psql}
}
