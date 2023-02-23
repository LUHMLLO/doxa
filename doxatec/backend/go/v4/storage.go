package main

import (
	"database/sql"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	ReadUser(uuid.UUID) (*User, error)
	UpdateUser(*User) error
	DeleteUser(uuid.UUID) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connectionString := "user=postgres dbname=postgres password=d@x@dm1n sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateUserTable()
}

func (s *PostgresStore) CreateUserTable() error {
	query := `create table if not exists users (
		id serial primary key,
		avatar varchar(250),
		username varchar(250),
		password varchar(250),
		customer varchar(250),
		created timestamp,
		modified timestamp,
		accessed timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateUser(*User) error {
	return nil
}

func (s *PostgresStore) ReadUser(id uuid.UUID) (*User, error) {
	return nil, nil
}

func (s *PostgresStore) UpdateUser(*User) error {
	return nil
}

func (s *PostgresStore) DeleteUser(id uuid.UUID) error {
	return nil
}
