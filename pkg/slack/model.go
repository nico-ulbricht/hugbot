package slack

type Message struct {
	ID          uuid.UUID `db:"id"`
	RecipientID uuid.UUID `db:"recipient_id"`
	Message     string    `db:"message"`
}
