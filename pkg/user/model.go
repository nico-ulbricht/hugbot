package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `db:"id"`
	ExternalID string    `db:"external_id"`
	Name       string    `db:"name"`
}
