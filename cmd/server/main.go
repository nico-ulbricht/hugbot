package main

import (
	"github.com/joho/godotenv"
	"github.com/nico-ulbricht/hugbot/pkg/db"
)

func init() {
	godotenv.Load()
}

func main() {
	psql := db.New()
	db.MustMigrate(psql, "file://migrations")
}
