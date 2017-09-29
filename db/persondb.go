package db

import (
	"github.com/luisfsantos/thysis/model"
	"log"
)

const (
	createUser = `
	INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id;
	`
	selectUserByID = `
	SELECT id, username, email, password FROM users WHERE id=$1;
	`
	selectAllUsers = `
	SELECT id, username, email, password FROM users
	`
)

func (db *pgDb) CreateUser(username, email, password string) (int64, error) {
	result, err := db.dbConnection.Exec(createUser, username, email, password)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return -1, err
	}
	return result.LastInsertId()
}

func (db *pgDb) SelectUserByID(ID int64) (*model.User, error)  {
	user := new(model.User)
	err := db.dbConnection.Get(&user, selectUserByID, ID)
	if err != nil {
		log.Printf("Error getting user with id: %v\n", err)
		return nil, err
	}
	return user, nil
}

func (db *pgDb) SelectAllUsers() ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := db.dbConnection.Select(&users, selectAllUsers)
	if err != nil {
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}
