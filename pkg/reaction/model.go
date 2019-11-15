package reaction

import (
	"github.com/google/uuid"
)

type Reaction struct {
	ID          uuid.UUID `db:"id"`
	SenderID    uuid.UUID `db:"sender_id"`
	RecipientID uuid.UUID `db:"recipient_id"`

	Amount int    `db:"amount"`
	Type   string `db:"type"`
}
