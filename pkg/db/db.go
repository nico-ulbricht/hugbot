package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	Database string `envconfig:"PSQL_DB"`
	Host     string `envconfig:"PSQL_HOST"`
	Password string `envconfig:"PSQL_PASS"`
	Port     int    `envconfig:"PSQL_PORT"`
	SSL      bool   `envconfig:"PSQL_SSL"`
	User     string `envconfig:"PSQL_USER"`

	MaxConnectionLifetime time.Duration `envconfig:"PSQL_MAX_LIFETIME" default:"5m"`
	MaxConnectionRetries  int           `envconfig:"PSQL_MAX_RETRIES" default:"30"`
	MaxIdleConnections    int           `envconfig:"PSQL_MAX_IDLE" default:"10"`
	MaxOpenConnections    int           `envconfig:"PSQL_MAX_OPEN" default:"20"`
}

func New() *sqlx.DB {
	var config Config
	envconfig.MustProcess("", &config)
	return NewWithConfig(config)
}

func NewWithConfig(config Config) *sqlx.DB {
	sslMode := "disable"
	if config.SSL == true {
		sslMode = "require"
	}

	connectionProps := []string{
		fmt.Sprintf("sslmode=%s", sslMode),
		fmt.Sprintf("host=%s", config.Host),
		fmt.Sprintf("port=%d", config.Port),
		fmt.Sprintf("dbname=%s", config.Database),
	}

	if config.User != "" {
		connectionProps = append(connectionProps, fmt.Sprintf("user=%s", config.User))
	}

	if config.Password != "" {
		connectionProps = append(connectionProps, fmt.Sprintf("password=%s", config.Password))
	}

	connectionString := strings.Join(connectionProps, " ")
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	retryCount := 0
	tick := time.Tick(1 * time.Second)
	err = db.Ping()
	for err != nil && retryCount < config.MaxConnectionRetries {
		select {
		case <-tick:
			err = db.Ping()
			retryCount++
		}
	}

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(config.MaxConnectionLifetime)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetMaxOpenConns(config.MaxOpenConnections)
	return db
}
