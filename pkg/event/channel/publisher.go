package channel

import (
	"context"

	"github.com/nico-ulbricht/hugbot/pkg/event"
)

type publisher struct {
	eventChannel chan event.Event
}

func (pub *publisher) Publish(ctx context.Context, event event.Event) error {
	pub.eventChannel <- event
	return nil
}

func NewPublisher(eventChannel chan event.Event) event.Publisher {
	return &publisher{
		eventChannel: eventChannel,
	}
}
