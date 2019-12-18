package event

import (
	"github.com/google/uuid"
)

var ReactionCreatedType Type = "reaction:created"

type ReactionCreatedPayload struct {
	Amount      int       `json:"amount"`
	RecipientID uuid.UUID `json:"recipientID"`
	SenderID    uuid.UUID `json:"senderID"`
	Type        string    `json:"type"`
}

type ReactionCreated struct {
	Meta    Meta                   `json:"meta"`
	Payload ReactionCreatedPayload `json:"payload"`
}

func (event ReactionCreated) GetMeta() Meta {
	return event.Meta
}
