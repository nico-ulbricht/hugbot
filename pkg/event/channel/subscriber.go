package channel

import (
	"context"

	"github.com/nico-ulbricht/hugbot/pkg/event"

	"github.com/pkg/errors"
)

type subscriber struct {
	eventChannel    chan event.Event
	handlerRegistry map[event.Type]event.HandleFunc
}

func (sub *subscriber) Consume(ctx context.Context, errChan chan error) {
	for {
		event := <-sub.eventChannel
		handleFunc := sub.handlerRegistry[event.GetMeta().Type]
		if handleFunc == nil {
			continue
		}

		err := handleFunc(ctx, event)
		if err != nil {
			errChan <- errors.WithStack(err)
		}
	}
}

func (sub *subscriber) Subscribe(eventType event.Type, handleFunc event.HandleFunc) {
	sub.handlerRegistry[eventType] = handleFunc
}

func NewSubscriber(eventChannel chan event.Event) event.Subscriber {
	return &subscriber{
		eventChannel: eventChannel,
	}
}
