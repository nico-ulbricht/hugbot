package event

import (
	"github.com/google/uuid"
)

var ReactionCreatedEventType Type = "reaction:created"

type ReactionCreatedPayload struct {
	RecipientID uuid.UUID `json:"recipientID"`
	SenderID    uuid.UUID `json:"senderID"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
}

type ReactionCreated struct {
	Meta    Meta                   `json:"meta"`
	Payload ReactionCreatedPayload `json:"payload"`
}

func (event *ReactionCreated) GetMeta() Meta {
	return event.Meta
}
