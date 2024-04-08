package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgresql struct {
	db *sql.DB
}

func New(user, pass, host, port, dataBase string) *Postgresql {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dataBase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Postgresql{db: db}
}

func (p *Postgresql) Close() error {
	return p.db.Close()
}
