package event

import "context"

type Subscriber interface {
	Consume(errChan chan error)
	Subscribe(eventType Type, handleFunc HandleFunc)
}

type HandleFunc func(ctx context.Context, event Event) error
