package reaction

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type loggingService struct {
	logger zerolog.Logger
	next   Service
}

func (svc *loggingService) Create(ctx context.Context, input CreateInput) (reaction *Reaction, err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "create").
			Str("recipient_id", input.RecipientID.String()).
			Str("sender_id", input.SenderID.String()).
			Str("type", input.Type).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Send()
		}
	}(time.Now())
	return svc.next.Create(ctx, input)
}

func (svc *loggingService) GetByRecipientID(ctx context.Context, recipientID uuid.UUID) (reactions []*Reaction, err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "get_by_recipient_id").
			Str("recipient_id", recipientID.String()).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Int("reaction_count", len(reactions)).Send()
		}
	}(time.Now())
	return svc.next.GetByRecipientID(ctx, recipientID)
}

func (svc *loggingService) GetBySenderID(ctx context.Context, senderID uuid.UUID) (reactions []*Reaction, err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "get_by_sender_id").
			Str("sender_id", senderID.String()).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Int("reaction_count", len(reactions)).Send()
		}
	}(time.Now())
	return svc.next.GetBySenderID(ctx, senderID)
}

func NewLoggingService(next Service, logger zerolog.Logger) Service {
	return &loggingService{
		next:   next,
		logger: logger,
	}
}
