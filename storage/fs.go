package storage

import (
	"bloc/utils/config"
	"io/ioutil"
	"mime/multipart"
	"os"
)

/*

File system backend
  simply write to the local file system

*/

type FilseSystem struct {
	Path string
}

func NewFileSystem() FilseSystem {
	return FilseSystem{
		Path: config.Conf.Storage.FileSystem.Path,
	}
}

func (fs FilseSystem) Create(id string, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return os.WriteFile(fs.Path+"/"+id, b, 0777)
}

func (fs FilseSystem) Delete(id string) error {
	return os.Remove(fs.Path + "/" + id)
}
