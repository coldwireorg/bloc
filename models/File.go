package models

/* File model */

type File struct {
	Id         string
	Name       string
	Size       int
	IsFavorite bool
	Key        string
	Parent     int
	Owner      string
}

func (f File) Create() error {
	return nil
}

func (f File) Delete() error {
	return nil
}

func (f File) Update() error {
	return nil
}

func (f File) Find() error {
	return nil
}

func (f File) Move() error {
	return nil
}
