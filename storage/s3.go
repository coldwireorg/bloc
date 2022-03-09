package storage

import (
	"bloc/utils/config"
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

/*

S3 backend
  Write file to any S3 cluster

*/

type S3 struct {
	*minio.Client
}

func NewS3() S3 {
	client, err := minio.New(config.Conf.Storage.S3.Address, &minio.Options{
		Creds: credentials.New(&credentials.Static{
			Value: credentials.Value{
				AccessKeyID:     config.Conf.Storage.S3.Id,
				SecretAccessKey: config.Conf.Storage.S3.Secret,
				SessionToken:    config.Conf.Storage.S3.Token,
				SignerType:      credentials.SignatureV4,
			},
		}),
		Region: config.Conf.Storage.S3.Region,
	})

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return S3{
		client,
	}
}

func (s S3) Create(id string, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = s.PutObject(context.Background(), config.Conf.Storage.S3.Bucket, id, f, file.Size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s S3) Delete(id string) error {
	return s.RemoveObject(context.Background(), config.Conf.Storage.S3.Bucket, id, minio.RemoveObjectOptions{})
}
