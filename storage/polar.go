package storage

import (
	"bloc/utils/config"
	"mime/multipart"
)

/*

Polar backend
  Write file to the polar decentralized network

*/

// TODO: implement it when Polar is OK
// https://codeberg.org/coldwire/polar
type Polar struct {
	Url    string
	Secret string
}

func NewPolar() Polar {
	return Polar{
		Url:    config.Conf.Storage.Polar.Url,
		Secret: config.Conf.Storage.Polar.Secret,
	}
}

func (p Polar) Create(id string, file *multipart.FileHeader) error {
	return nil
}

func (p Polar) Delete(id string) error {
	return nil
}

func (p Polar) Get(id string) ([]byte, error) {
	return nil, nil
}
