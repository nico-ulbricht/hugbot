package db

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var txContextKey = "psql-tx"

type TxProvider interface {
	Beginx() (*sqlx.Tx, error)
}

func NewMiddleware(provider TxProvider) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tx, err := provider.Beginx()
			if err != nil {
				return nil, errors.WithStack(err)
			}

			txContext := context.WithValue(ctx, txContextKey, tx)
			resp, err := next(txContext, req)
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					return nil, errors.WithStack(err)
				}

				return nil, errors.WithStack(err)
			}

			err = tx.Commit()
			return resp, errors.WithStack(err)
		}
	}
}

func TxFromContext(ctx context.Context) (*sqlx.Tx, error) {
	tx, ok := ctx.Value(txContextKey).(*sqlx.Tx)
	if ok == false || tx == nil {
		return nil, ErrNoTx{}
	}

	return tx, nil
}
