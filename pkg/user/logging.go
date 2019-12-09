package user

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

func (svc *loggingService) Upsert(ctx context.Context, input UpsertInput) (user *User, err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "upsert").
			Str("external_id", input.ExternalID).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Str("user_id", user.ID.String()).Send()
		}
	}(time.Now())
	return svc.next.Upsert(ctx, input)
}

func (svc *loggingService) GetByID(ctx context.Context, userID uuid.UUID) (user *User, err error) {
	defer func(begin time.Time) {
		methodLogger := svc.logger.
			With().
			Str("method", "get_by_id").
			Str("user_id", userID.String()).
			Dur("duration", time.Since(begin)).
			Logger()

		if err != nil {
			methodLogger.Error().Err(err).Send()
		} else {
			methodLogger.Debug().Send()
		}
	}(time.Now())
	return svc.next.GetByID(ctx, userID)
}

func NewLoggingService(next Service, logger zerolog.Logger) Service {
	return &loggingService{
		next:   next,
		logger: logger,
	}
}
