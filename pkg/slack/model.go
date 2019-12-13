package slack

import (
	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID `db:"id"`
	RecipientID uuid.UUID `db:"recipient_id"`
	Message     string    `db:"message"`
}

type eventType string

var (
	messageEventType       eventType = "message"
	reactionAddedEventType eventType = "reaction_added"
)

type eventItem struct {
	ID string `json:"ts"`
}

type messageEvent struct {
	ID       string `json:"ts"`
	SenderID string `json:"user"`
	Text     string `json:"text"`
}

type reactionAddedEvent struct {
	ID          string    `json:"event_ts"`
	SenderID    string    `json:"user"`
	RecipientID string    `json:"item_user"`
	Reaction    string    `json:"reaction"`
	Item        eventItem `json:"item"`
}
