package reaction

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/nico-ulbricht/hugbot/pkg/db"
)

type Repository interface {
	Insert(ctx context.Context, reaction *Reaction) (*Reaction, error)
	GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error)
	GetByReferenceID(ctx context.Context, referenceID string) (*Reaction, error)
	GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error)
}

type repository struct{}

func (rp *repository) Insert(ctx context.Context, reaction *Reaction) (*Reaction, error) {
	tx, err := db.TxFromContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stmt, err := tx.PreparexContext(ctx, `
		insert into reactions (
			id,
			recipient_id,
			sender_id,
			reference_id,
			amount,
			type,
			created_at,
			updated_at
		) values ($1, $2, $3, $4, $5, $6, $7, $8)
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	now := time.Now().UTC()
	_, err = stmt.ExecContext(
		ctx,
		reaction.ID,
		reaction.RecipientID,
		reaction.SenderID,
		reaction.ReferenceID,
		reaction.Amount,
		reaction.Type,
		now,
		now,
	)

	return reaction, errors.WithStack(err)
}

func (rp *repository) GetByRecipientID(ctx context.Context, recipientID uuid.UUID) ([]*Reaction, error) {
	tx, err := db.TxFromContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stmt, err := tx.PreparexContext(ctx, `
		select
			id,
			recipient_id,
			sender_id,
			reference_id,
			amount,
			type,
		from reactions
		where recipient_id = $1
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	var reactions []*Reaction
	err = stmt.SelectContext(ctx, &reactions, recipientID)
	return reactions, errors.WithStack(err)
}

func (rp *repository) GetByReferenceID(ctx context.Context, referenceID string) (*Reaction, error) {
	tx, err := db.TxFromContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stmt, err := tx.PreparexContext(ctx, `
		select
			id,
			recipient_id,
			sender_id,
			reference_id,
			amount,
			type,
		from reactions
		where reference_id = $1
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	var reaction Reaction
	err = stmt.GetContext(ctx, &reaction, referenceID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &reaction, errors.WithStack(err)
}

func (rp *repository) GetBySenderID(ctx context.Context, senderID uuid.UUID) ([]*Reaction, error) {
	tx, err := db.TxFromContext(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	stmt, err := tx.PreparexContext(ctx, `
		select
			id,
			recipient_id,
			sender_id,
			reference_id,
			amount,
			type,
		from reactions
		where sender_id = $1
	`)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { _ = stmt.Close() }()
	var reactions []*Reaction
	err = stmt.SelectContext(ctx, &reactions, senderID)
	return reactions, errors.WithStack(err)
}

func NewRepository() Repository {
	return &repository{}
}
