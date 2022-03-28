package models

type User struct {
	Username   string
	Password   string
	AuthMode   string
	PrivateKey string
	PublicKey  string
	Root       string
}

func (u User) Create() error
func (u User) Delete() error
func (u User) Find() error
func (u User) Exist() bool
