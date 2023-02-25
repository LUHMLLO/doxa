package storages

import (
	"doxatec/types"
	"doxatec/utils"
	"fmt"

	"github.com/google/uuid"
)

func (s *PostgresStore) Query_CreateUserTable() error {
	query := `
		create table if not exists users (
			id varchar(250) primary key,
			username varchar(250),
			password varchar(250),
			
			profile varchar(250),

			created timestamp,
			modified timestamp,
			accessed timestamp
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) Query_CreateUser(u *types.User) error {
	query := (`
		insert into users (
			id, 
			username, 
			password, 
			
			profile,

			created, 
			modified, 
			accessed
		)
		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	`)

	_, err := s.db.Query(
		query,
		u.ID,
		u.Username,
		u.Password,

		u.Profile,

		u.Created,
		u.Modified,
		u.Accessed,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) Query_ReadUsers() ([]*types.User, error) {
	rows, err := s.db.Query(`select * from users`)
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
	for rows.Next() {
		user, err := utils.ScanIntoUsers(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *PostgresStore) Query_ReadUserByID(id uuid.UUID) (*types.User, error) {
	rows, err := s.db.Query("select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return utils.ScanIntoUsers(rows)
	}
	return nil, fmt.Errorf("user %s not found", id)
}

func (s *PostgresStore) Query_UpdateUserByID(id uuid.UUID) (*types.User, error) {
	query := (`
		update users set
			id = $1, 
			username = $2, 
			password = $3, 
			
			profile = $4,

			created = $5, 
			modified = $6, 
			accessed = $7

		where id = $1"
	`)

	_, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("user %s not found", id)
}

func (s *PostgresStore) Query_DeleteUser(id uuid.UUID) error {
	_, err := s.db.Query("delete from users where id = $1", id)
	return err
}
