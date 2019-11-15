package event

import "time"

type EventType string

type Event struct {
	Meta    EventMeta   `json:"meta"`
	Payload interface{} `json:"payload"`
}

type EventMeta struct {
	Created time.Time `json:"createdAt"`
	Type    EventType `json:"eventType"`
}
