package db

type ErrNoTx struct{}

func (err ErrNoTx) Error() string {
	return "no psql tx found in context"
}
