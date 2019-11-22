package event

import "time"

type Type string

type Event interface {
	GetMeta() Meta
}

type Meta struct {
	Created time.Time `json:"createdAt"`
	Type    Type      `json:"type"`
}
