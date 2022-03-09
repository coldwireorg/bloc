package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
			Driver: os.Getenv("BLOC_DATABASE_DRIVER"),
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
			Driver: os.Getenv("BLOC_STORAGE_DRIVER"),
			Quota:  os.Getenv("BLOC_STORAGE_QUOTA"),
			Polar: StoragePolarConfig{
				Url:    os.Getenv("BLOC_STORAGE_POLAR_URI"),
				Secret: os.Getenv("BLOC_STORAGE_POLAR_SECRET"),
			},
			FileSystem: StorageFileSystemConfig{
				Path: os.Getenv("BLOC_STORAGE_FS_PATH"),
			},
			S3: StorageS3Config{
				Address: os.Getenv("BLOC_STORAGE_S3_ADDRESS"),
				Bucket:  os.Getenv("BLOC_STORAGE_S3_BUCKET"),
				Options: minio.Options{
					Creds:  credentials.NewStaticV4(os.Getenv("BLOC_STORAGE_S3_ID"), os.Getenv("BLOC_STORAGE_S3_SECRET"), os.Getenv("BLOC_STORAGE_S3_TOKEN")),
					Region: os.Getenv("BLOC_STORAGE_S3_REGION"),
					Secure: true,
				},
			},
		},

		Oauth: OauthConfig{
			Server: os.Getenv("BLOC_OAUTH_SERVER"),
			Config: oauth2.Config{
				ClientID:    os.Getenv("BLOC_OAUTH_CLIENT"),
				RedirectURL: os.Getenv("BLOC_OAUTH_CALLBACK"),
			},
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
	minio.Options
}

type OauthConfig struct {
	Config oauth2.Config
	Server string
}
