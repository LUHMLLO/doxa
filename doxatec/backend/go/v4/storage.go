package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	ReadUsers() ([]*User, error)
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
		id varchar(250) primary key,
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

func (s *PostgresStore) CreateUser(u *User) error {
	query := (`
		insert into users 
		(id, avatar, username, password, customer, created, modified, accessed)
		values 
		($1,$2,$3,$4,$5,$6,$7,$8)
	`)

	res, err := s.db.Query(query, u.ID, u.Avatar, u.Username, u.Password, u.Customer, u.Created, u.Modified, u.Accessed)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}

func (s *PostgresStore) ReadUsers() ([]*User, error) {
	rows, err := s.db.Query(`select * from users`)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user := new(User)
		if err := rows.Scan(&user.ID, &user.Avatar, &user.Username, &user.Password, &user.Customer, &user.Created, &user.Modified, &user.Accessed); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
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
