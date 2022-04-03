package models

type Folder struct {
	Id     string
	Name   string
	Owner  string
	Parent string
}

func (f Folder) Create() error {
	return nil
}

func (f Folder) Delete() error {
	return nil
}

func (f Folder) Find() error {
	return nil
}

func (f Folder) Move() error {
	return nil
}
