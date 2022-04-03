package models

type Share struct {
	Id         string
	Name       string
	Size       int
	IsFavorite bool
	Key        string
	Parent     string
	Owner      string
}

func (u Share) Add() error {
	return nil
}

func (u Share) Revoke() error {
	return nil
}

func (u Share) Find() error {
	return nil
}
