package event

import "context"

type Subscriber interface {
	Consume(ctx context.Context, errChan chan error)
	Subscribe(eventType Type, handleFunc HandleFunc)
}

type HandleFunc func(ctx context.Context, event Event) error
