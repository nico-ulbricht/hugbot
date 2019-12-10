package event

import "context"

type Subscriber interface {
	Consume(ctx context.Context, errChan chan error)
}

type HandleFunc func(ctx context.Context, event Event) error
