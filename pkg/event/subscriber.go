package event

import "context"

type Subscriber interface {
	Consume(ctx context.Context) error
}
