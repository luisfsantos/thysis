package model

type Model struct {
	DB
}

type DB interface {
	UserDB
}