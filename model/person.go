package model

type User struct {
	ID int64 `db:"id"`
	Username string `db:"username"`
	Email string `db:"email"`
	Password string `db:"password"`
}

type UserDB interface {
	CreateUser(username, email, password string) (int64, error)
	SelectUserByID(ID int64) (*User, error)
	SelectAllUsers() ([]*User, error)
}