package database

import (
	"bloc/utils/config"
	"context"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

//go:embed schema.sql
var tables string

// Expose database to the whole project
var DB *pgxpool.Pool
var Ready bool = false

// Connect to the database
func Connect() error {
	var err error

	log.Info().Msg("Using driver: " + config.Conf.Database.Driver)

	// Run until the database is connected
	for {
		switch config.Conf.Database.Driver {
		case "postgres":
			DSN := []string{
				"postgresql://",
				config.Conf.Database.Postgres.Address,
				":" + config.Conf.Database.Postgres.Port,
				"/" + config.Conf.Database.Postgres.Name,
				"?user=" + config.Conf.Database.Postgres.User,
				"&password=" + config.Conf.Database.Postgres.Password,
			}

			DB, err = pgxpool.Connect(context.Background(), strings.Join(DSN[:], ""))
		default:
			log.Fatal().Msg("Database driver not found!")
		}

		if err != nil {

			fmt.Print(err)

			Ready = false
			log.Err(err).Msg(err.Error())
			time.Sleep(5 * time.Second)
		} else {
			log.Info().Msg("Successfully connected to the database!")
			Ready = true
			break
		}
	}

	// Create tables
	_, err = DB.Exec(context.Background(), tables)
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}
