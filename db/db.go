package db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type Configuration struct {
	DBConnection string
}

type pgDb struct {
	dbConnection *sqlx.DB
}

func InitDB (configuration Configuration) (*pgDb, error) {
	dbConnection, err := sqlx.Connect("postgres", configuration.DBConnection)
	if err != nil {
		return nil, err
	} else {
		p := &pgDb {dbConnection:dbConnection}
		err := p.dbConnection.Ping()
		if err != nil {
			return nil, err
		}
		p.createTablesIfNotExist();
		return p, nil
	}
}

func (db *pgDb) createTablesIfNotExist() {
	const (
		userTable = `
		CREATE TABLE IF NOT EXISTS users (
		id bigserial PRIMARY KEY,
		username varchar(12) UNIQUE,
		email text NOT NULL,
		password text NOT NULL);
		`
	)
	db.dbConnection.MustExec(userTable)
}
