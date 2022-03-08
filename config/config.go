package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

// Export config to the whole project
var Conf *Config

// Init config to ensure that env variable are loaded
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	Conf = &Config{
		Server: ServerConfig{
			Address: os.Getenv("BLOC_SERVER_ADDRESS"),
			Port:    os.Getenv("BLOC_SERVER_PORT"),
		},
		Database: DatabaseConfig{
			Type: os.Getenv("BLOC_DATABASE_TYPE"),
			Postgres: DatabasePostgresConfig{
				Address:  os.Getenv("BLOC_DATABASE_ADDRESS"),
				Port:     os.Getenv("BLOC_DATABASE_PORT"),
				User:     os.Getenv("BLOC_DATABASE_USER"),
				Password: os.Getenv("BLOC_DATABASE_PASSWORD"),
				Name:     os.Getenv("BLOC_DATABASE_NAME"),
			},
			Sqlite: DatabaseSqliteConfig{
				Path: os.Getenv("BLOC_DATABASE_PATH"),
			},
		},
		Storage: StorageConfig{
			Path:  os.Getenv("BLOC_STORAGE_PATH"),
			Quota: os.Getenv("BLOC_STORAGE_QUOTA"),
		},
		Polar: PolarConfig{
			Uri:    os.Getenv("BLOC_POLAR_URI"),
			Secret: os.Getenv("BLOC_POLAR_SECRET"),
		},
		Oauth: oauth2.Config{
			ClientID:     os.Getenv("BLOC_OAUTH_CLIENT"),
			ClientSecret: os.Getenv("BLOC_OAUTH_SECRET"),
			RedirectURL:  os.Getenv("BLOC_OAUTH_CALLBACK"),
		},
	}
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Storage  StorageConfig
	Polar    PolarConfig
	Oauth    oauth2.Config
}

type ServerConfig struct {
	Address string // Address for webserver to listen on
	Port    string // Port for webserver to listen on
}

type DatabaseConfig struct {
	Type     string                 // postgres or sqlite
	Postgres DatabasePostgresConfig // Postgres config
	Sqlite   DatabaseSqliteConfig   // Sqlite config
}

type DatabasePostgresConfig struct {
	Address  string // Address of the postgres instance
	Port     string // Port of the postgres instance
	User     string // USer of the database
	Password string // Password of the database
	Name     string // Name of the database
}

type DatabaseSqliteConfig struct {
	Path string // path to the database file
}

type StorageConfig struct {
	Path  string // Path where to store encrypted files
	Quota string // Total size of files that users cans upload
}

type PolarConfig struct {
	Uri    string // unix socket/address of the polar node
	Secret string // Opetional secret to access a private polar node
}
