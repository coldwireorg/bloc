package config

import (
	"bloc/utils/env"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Export config to the whole project
var Conf *Config

// Init config to ensure that env variable are loaded
func Init(file string) {
	// if config file is specified we use it
	if file != "" {
		log.Info().Msg("Loading config file: " + file)
		f, err := os.ReadFile(file)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		var c Config
		_, err = toml.Decode(string(f), &c)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		Conf = &c

		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	Conf = &Config{
		Server: ServerConfig{
			Address: env.Get("BLOC_SERVER_ADDRESS", "0.0.0.0"),
			Port:    env.Get("BLOC_SERVER_PORT", "3000"),
		},

		Database: DatabaseConfig{
			Driver: env.Get("BLOC_DATABASE_DRIVER", "sqlite"),
			Postgres: DatabasePostgresConfig{
				Address:  env.Get("BLOC_DATABASE_ADDRESS", "127.0.0.1"),
				Port:     env.Get("BLOC_DATABASE_PORT", "5432"),
				User:     env.Get("BLOC_DATABASE_USER", "postgres"),
				Password: env.Get("BLOC_DATABASE_PASSWORD", "123456789"),
				Name:     env.Get("BLOC_DATABASE_NAME", "bloc"),
			},
			Sqlite: DatabaseSqliteConfig{
				Path: env.Get("BLOC_DATABASE_PATH", "/tmp/bloc.sqlite"),
			},
		},

		Storage: StorageConfig{
			Driver: env.Get("BLOC_STORAGE_DRIVER", "fs"),
			Quota:  env.Get("BLOC_STORAGE_QUOTA", "4G"),
			FileSystem: StorageFileSystemConfig{
				Path: env.Get("BLOC_STORAGE_FS_PATH", "/tmp/bloc/"),
			},
			S3: StorageS3Config{
				Address: env.Get("BLOC_STORAGE_S3_ADDRESS", ""),
				Bucket:  env.Get("BLOC_STORAGE_S3_BUCKET", "bloc"),
				Id:      env.Get("BLOC_STORAGE_S3_ID", ""),
				Secret:  env.Get("BLOC_STORAGE_S3_SECRET", ""),
				Token:   env.Get("BLOC_STORAGE_S3_TOKEN", ""),
				Region:  env.Get("BLOC_STORAGE_S3_REGION", ""),
			},
			Polar: StoragePolarConfig{
				Url:    env.Get("BLOC_STORAGE_POLAR_URI", "unix:///var/run/polar.sock"),
				Secret: env.Get("BLOC_STORAGE_POLAR_SECRET", "acab"),
			},
		},

		Oauth: OauthConfig{
			Server:   env.Get("BLOC_OAUTH_SERVER", ""),
			Id:       env.Get("BLOC_OAUTH_CLIENT", "bloc"),
			Secret:   env.Get("BLOC_OAUTH_SECRET", ""),
			Callback: env.Get("BLOC_OAUTH_CALLBACK", "https://bloc.coldwire.org/api/user/auth/oauth2/callback"),
		},
	}
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Storage  StorageConfig
	Oauth    OauthConfig
}

type ServerConfig struct {
	Address string // Address for webserver to listen on
	Port    string // Port for webserver to listen on
}

type DatabaseConfig struct {
	Driver   string                 // postgres or sqlite
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

/* STORAGE DRIVERS */
type StorageConfig struct {
	Driver     string
	Quota      string // Total size of files that users cans upload
	FileSystem StorageFileSystemConfig
	Polar      StoragePolarConfig
	S3         StorageS3Config
}

type StoragePolarConfig struct {
	Url    string // unix socket/address of the polar node
	Secret string // Opetional secret to access a private polar node
}

type StorageFileSystemConfig struct {
	Path string // Path where to store encrypted files
}

type StorageS3Config struct {
	Bucket  string
	Address string
	Id      string
	Secret  string
	Token   string
	Region  string
}

type OauthConfig struct {
	Server   string
	Id       string
	Secret   string
	Callback string
}
