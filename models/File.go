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

func (f File) Create() error
func (f File) Delete() error
func (f File) Update() error
func (f File) Find() error
func (f File) Move() error
