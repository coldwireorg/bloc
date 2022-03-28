package models

type Folder struct {
	Id     string
	Name   string
	Owner  string
	Parent string
}

func (f Folder) Create() error
func (f Folder) Delete() error
func (f Folder) Find() error
func (f Folder) Move() error
