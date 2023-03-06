package lib

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage interface {
	Query_create_users_table() error
	Query_read_all_users_from_table() ([]*User, error)

	Query_before_insert_user(u *User) (*User, error)
	Query_insert_user(u *User) error

	Query_read_user_where_column(column, param string) (*User, error)
	Query_update_user_column_where_ID(id uuid.UUID, column string, param any) error
	Query_delete_user_where_ID(id uuid.UUID) (uuid.UUID, error)

	Query_before_signin_user(username, password string) (*User, error)
}

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

func (storer *Database) Init() error {
	storer.Query_create_users_table()

	return nil
}
