package reaction

import (
	"github.com/google/uuid"
)

type Reaction struct {
	ID          uuid.UUID `db:"id"`
	RecipientID uuid.UUID `db:"recipient_id"`
	SenderID    uuid.UUID `db:"sender_id"`
	ReferenceID string    `db:"reference_id"`

	Amount int    `db:"amount"`
	Type   string `db:"type"`
}

func newReaction(
	recipientID uuid.UUID,
	senderID uuid.UUID,
	referenceID string,
	amount int,
	reactionType string,
) *Reaction {
	return &Reaction{
		ID:          uuid.New(),
		RecipientID: recipientID,
		SenderID:    senderID,
		ReferenceID: referenceID,
		Amount:      amount,
		Type:        reactionType,
	}
}
