package event

import (
	"github.com/google/uuid"
)

var ReactionCreatedEventType EventType = "reaction:created"

type ReactionCreatedPayload struct {
	RecipientID uuid.UUID `json:"recipientID"`
	SenderID    uuid.UUID `json:"senderID"`
	Amount      int       `json:"amount"`
	Type        string    `json:"type"`
}

type ReactionCreatedEvent struct {
	Meta    EventMeta              `json:"meta"`
	Payload ReactionCreatedPayload `json:"payload"`
}
