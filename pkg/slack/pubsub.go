package slack

import (
	"context"

	"github.com/nico-ulbricht/hugbot/pkg/event"
)

func SubscribeReactionEventHandlers(
	svc Service,
	subscriber event.Subscriber,
) {
	reactionCreatedHandleFunc := newReactionCreatedHandleFunc(svc)
	subscriber.Subscribe(event.ReactionCreatedType, reactionCreatedHandleFunc)
}

func newReactionCreatedHandleFunc(svc Service) event.HandleFunc {
	return func(ctx context.Context, event event.Event) error {
		return nil
	}
}
