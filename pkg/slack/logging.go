package slack

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

type loggingService struct {
	logger zerolog.Logger
	next   Service
}

func (svc *loggingService) SendMessage(ctx context.Context, msg *Message) (err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "send_message").
			Str("recipient_id", msg.RecipientID.String()).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Send()
		}
	}(time.Now())
	return svc.next.SendMessage(ctx, msg)
}

func (svc *loggingService) HandleMessage(ctx context.Context, input handleMessageInput) (err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "handle_message").
			Str("reference_id", input.ReferenceID).
			Str("sender_id", input.SenderID).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Send()
		}
	}(time.Now())
	return svc.next.HandleMessage(ctx, input)
}

func (svc *loggingService) HandleReaction(ctx context.Context, input handleReactionInput) (err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "handle_reaction").
			Str("reference_id", input.ReferenceID).
			Str("sender_id", input.SenderID).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Send()
		}
	}(time.Now())
	return svc.next.HandleReaction(ctx, input)
}

func NewLoggingService(next Service, logger zerolog.Logger) Service {
	return &loggingService{
		next:   next,
		logger: logger,
	}
}
