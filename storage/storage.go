package storage

import (
	"mime/multipart"
)

var Driver Interface

type Interface interface {
	Create(id string, file *multipart.FileHeader) error
	Delete(id string) error
}

type Drivers struct {
	S3
	FilseSystem
}

func New(driver string) Interface {
	switch driver {
	case "s3":
		return NewS3()
	default:
		return NewFileSystem()
	}
}

func Init(d string) {
	i := New(d)
	Driver = i
}
