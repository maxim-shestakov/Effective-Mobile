package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(user, pass, host, port, dataBase, schema string) *Repository {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", user, pass, host, port, dataBase, schema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{db: db}
}

func (r *Repository) GetDB() *sql.DB {
	return r.db
}

func (p *Repository) Close() error {
	return p.db.Close()
}
