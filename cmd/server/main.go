package main

import (
	"github.com/joho/godotenv"

	"github.com/nico-ulbricht/hugbot/pkg/db"
	"github.com/nico-ulbricht/hugbot/pkg/reaction"
	"github.com/nico-ulbricht/hugbot/pkg/slack"
	"github.com/nico-ulbricht/hugbot/pkg/user"
)

func init() {
	godotenv.Load()
}

func main() {
	psql := db.New()
	db.MustMigrate(psql, "file://migrations")

	reactionRepository := reaction.NewRepository()
	reactionService := reaction.NewService(reactionRepository)

	userRepository := user.NewRepository()
	userService := user.NewService(userRepository)

	_ = slack.NewService(reactionService, userService)
}
