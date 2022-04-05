package storage

import (
	"bloc/utils/config"
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"time"
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

func (fs FilseSystem) Get(id string) (io.Reader, error) {
	f, err := os.Open(fs.Path + "/" + id)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)
	buf := make([]byte, 16)

	go func() {
		for {
			_, err := reader.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Fatal(err)
				}
				break
			}
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		f.Close()
	}()

	return reader, err
}
