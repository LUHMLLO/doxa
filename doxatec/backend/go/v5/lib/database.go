package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "dbname=doxatec user=doxadmin password=d@x@dm1n sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (s *Database) Init() error {
	s.CreateTable("users", []string{
		"id varchar(250) primary key",
		"username varchar(250)",
		"password varchar(250)",
		"avatar varchar(250)",
		"name varchar(250)",
		"email varchar(250)",
		"phone varchar(250)",
		"created varchar(250)",
		"modified varchar(250)",
	})

	s.CreateTable("devices", []string{
		"id varchar(250) primary key",
		"name varchar(250)",
		"tempsup varchar(250)",
		"tempmid varchar(250)",
		"tempsub varchar(250)",
		"created varchar(250)",
		"modified varchar(250)",
	})

	return nil
}

func (s *Database) CreateTable(table string, cols_vals []string) error {
	query := fmt.Sprintf("create table if not exists %s (\n%s)", table, StringToQuery(cols_vals))

	// fmt.Println(query)
	// return nil

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *Database) InsertToTable(table string, cols_vals []interface{}) error {
	query := fmt.Sprintf("insert into %s values (\n%s)", table, InterfaceToQuery(cols_vals))

	// fmt.Println(query)
	// return nil

	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func (s *Database) ReadFromTable(table string) error {
	query := fmt.Sprintf(`select * from %s`, table)

	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}

	users := []any{}
	for rows.Next() {
		var id, username, password, avatar, name, email, phone, created, modified string
		if err := rows.Scan(id, username, password, avatar, name, email, phone, created, modified); err != nil {
			log.Fatal(err)
		}
		users = append(users, id)
	}

	fmt.Println(users)

	return nil
}
