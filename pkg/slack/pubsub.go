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
	return func(ctx context.Context, req interface{}) error {
		reactionCreatedEvent := req.(event.ReactionCreated)
		return svc.HandleReactionCreated(ctx, handleReactionCreatedInput{
			Amount:      reactionCreatedEvent.Payload.Amount,
			RecipientID: reactionCreatedEvent.Payload.RecipientID,
			SenderID:    reactionCreatedEvent.Payload.SenderID,
			Type:        reactionCreatedEvent.Payload.Type,
		})
	}
}
