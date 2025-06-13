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

func (s *sqliteStore) CreateNewContent(c *content) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	err = func(c *content) error {
		query := `insert into content values (?, ?, ?)`
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(c.Id, c.Ctype, c.Data)
		if err != nil {
			return err
		}
		return nil
	}(c)
	if err == nil {
		return tx.Commit()
	} else {
		_ = tx.Rollback()
		return fmt.Errorf("insert logic failed: %w", err)
	}
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

func (s *sqliteStore) CheckForSecretKey(key string) bool {
	tx, err := s.db.Begin()
	if err != nil {
		return false
	}
	defer tx.Rollback()

	query := `select * from creds where key_hash = ?`
	row := tx.QueryRow(query, key)
	dbkey := ""
	if err := row.Scan(&dbkey); err != nil {
		return false
	}
	return dbkey == key
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
	Image
)
