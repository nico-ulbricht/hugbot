package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/nico-ulbricht/hugbot/pkg/db"
	"github.com/nico-ulbricht/hugbot/pkg/event"
	"github.com/nico-ulbricht/hugbot/pkg/event/channel"
	"github.com/nico-ulbricht/hugbot/pkg/reaction"
	"github.com/nico-ulbricht/hugbot/pkg/slack"
	"github.com/nico-ulbricht/hugbot/pkg/user"
)

func init() {
	godotenv.Load()
}

type configuration struct {
	Port int `envconfig:"PORT" default:"8080"`
}

func main() {
	var config configuration
	envconfig.MustProcess("", &config)
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	psql := db.New()
	db.MustMigrate(psql, "file://migrations")
	subscriberLogger := logger.With().Str("component", "subscriber").Logger()

	reactionChannel := make(chan event.Event)
	reactionPublisher := channel.NewPublisher(reactionChannel)
	reactionLogger := logger.With().Str("component", "reaction").Logger()
	reactionRepository := reaction.NewRepository(psql)
	reactionService := reaction.NewService(reactionPublisher, reactionRepository)
	reactionService = reaction.NewLoggingService(reactionService, reactionLogger)

	userLogger := logger.With().Str("component", "user").Logger()
	userRepository := user.NewRepository(psql)
	userService := user.NewService(userRepository)
	userService = user.NewLoggingService(userService, userLogger)

	slackSubscriber := channel.NewSubscriber(reactionChannel)
	slackSubscriber = event.NewLoggedSubscriber(slackSubscriber, subscriberLogger)
	slackLogger := logger.With().Str("component", "slack").Logger()
	slackService := slack.NewService(reactionService, userService)
	slackService = slack.NewLoggingService(slackService, slackLogger)
	slackHandler := slack.NewHTTPHandler(slackService, slackLogger)
	slack.SubscribeReactionEventHandlers(slackService, slackSubscriber)

	port := fmt.Sprintf(":%d", config.Port)
	server := mux.NewRouter()
	server.PathPrefix("/slack").Handler(slackHandler)

	errChan := make(chan error)
	go slackSubscriber.Consume(errChan)
	go func(chan error) {
		logger.Info().
			Str("component", "server").
			Str("method", "listen_and_server").
			Str("port", port).
			Send()

		errChan <- http.ListenAndServe(port, server)
	}(errChan)

	select {
	case err := <-errChan:
		logger.Error().Err(err).Send()
	}
}
