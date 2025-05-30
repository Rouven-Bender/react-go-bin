package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	db *sql.DB
}

func NewSqliteStore() (*sqliteStore, error) {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &sqliteStore{
		db: db,
	}, nil
}

func (s *sqliteStore) LookupContent(id string) (*content, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	query := `select * from content where id = ?`
	rows, err := tx.Query(query, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		c, err := scanIntoContent(rows)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, fmt.Errorf("content %s not found", id)
}

func scanIntoContent(rows *sql.Rows) (*content, error) {
	c := new(content)
	err := rows.Scan(
		&c.Id,
		&c.Ctype,
		&c.Data,
	)
	if err != nil {
		return nil, err
	}
	return c, err
}

type content struct {
	Id    string      `json:"id"`
	Ctype contenttype `json:"type"`
	Data  string      `json:"data"`
}
type contenttype uint8

const (
	Link contenttype = iota
	Plaintext
)
