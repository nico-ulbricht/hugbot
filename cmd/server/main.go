package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/nico-ulbricht/hugbot/pkg/db"
	"github.com/nico-ulbricht/hugbot/pkg/reaction"
	"github.com/nico-ulbricht/hugbot/pkg/slack"
	"github.com/nico-ulbricht/hugbot/pkg/user"
)

func init() {
	godotenv.Load()
}

func main() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	psql := db.New()
	db.MustMigrate(psql, "file://migrations")

	reactionLogger := logger.With().Str("component", "reaction").Logger()
	reactionRepository := reaction.NewRepository()
	reactionService := reaction.NewService(reactionRepository)
	reactionService = reaction.NewLoggingService(reactionService, reactionLogger)

	userLogger := logger.With().Str("component", "user").Logger()
	userRepository := user.NewRepository()
	userService := user.NewService(userRepository)
	userService = user.NewLoggingService(userService, userLogger)

	slackLogger := logger.With().Str("component", "slack").Logger()
	slackService := slack.NewService(reactionService, userService)
	slackService = slack.NewLoggingService(slackService, slackLogger)
}
