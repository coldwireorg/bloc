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

func (u Share) Add() error
func (u Share) Revoke() error
func (u Share) Find() error
