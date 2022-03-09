package storage

import (
	"bloc/config"
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

/*

S3 backend
  Write file to the polar decentralized network

*/

type S3 struct {
	*minio.Client
}

func NewS3() S3 {
	client, err := minio.New(config.Conf.Storage.S3.Address, &config.Conf.Storage.S3.Options)
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
