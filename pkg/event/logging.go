package event

import (
	"github.com/rs/zerolog"
)

type loggedSubscriber struct {
	next   Subscriber
	logger zerolog.Logger
}

func (subscriber *loggedSubscriber) Consume(errChan chan error) {
	subscriber.logger.Info().
		Str("method", "consume").
		Send()

	subscriber.next.Consume(errChan)
}

func (subscriber *loggedSubscriber) Subscribe(eventType Type, handleFunc HandleFunc) {
	subscriber.logger.Info().
		Str("method", "subscribe").
		Str("event_type", string(eventType)).
		Send()

	subscriber.next.Subscribe(eventType, handleFunc)
}

func NewLoggedSubscriber(next Subscriber, logger zerolog.Logger) Subscriber {
	return &loggedSubscriber{
		next:   next,
		logger: logger,
	}
}
