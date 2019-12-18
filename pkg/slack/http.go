package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"github.com/rs/zerolog"
)

func NewHTTPHandler(svc Service, logger zerolog.Logger) http.Handler {
	eventLogger := logger.With().Str("path", "/slack/events").Logger()
	eventHandler := newEventHandler(svc, eventLogger)

	router := mux.NewRouter()
	router.HandleFunc("/slack/events", eventHandler)
	return router
}

func newEventHandler(svc Service, logger zerolog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		rawMessage := json.RawMessage(body)
		event, err := slackevents.ParseEvent(rawMessage, slackevents.OptionNoVerifyToken())

		if err != nil {
			logger.Warn().Err(err).Msg("failed parsing event")
			w.WriteHeader(http.StatusBadRequest)
		}

		switch event.Type {
		case slackevents.URLVerification:
			var challengeReq *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &challengeReq)
			if err != nil {
				logger.Warn().Err(err).Msg("failed unmarshaling event")
				w.WriteHeader(http.StatusBadRequest)
			}

			w.Header().Set("Content-Type", "text")
			w.Write([]byte(challengeReq.Challenge))

		case slackevents.CallbackEvent:
			innerEvent := event.InnerEvent
			switch eventData := innerEvent.Data.(type) {
			case *slackevents.MessageEvent:
				go svc.HandleMessage(ctx, handleMessageInput{
					Message:     eventData.Text,
					ReferenceID: eventData.TimeStamp,
					SenderID:    eventData.User,
				})

			case *slack.ReactionAddedEvent:
				go svc.HandleReaction(ctx, handleReactionInput{
					RecipientID: eventData.ItemUser,
					ReferenceID: eventData.Item.Timestamp,
					SenderID:    eventData.User,
					Type:        eventData.Reaction,
				})
			}
		}
	}
}
