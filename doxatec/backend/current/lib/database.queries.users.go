package lib

import (
	"github.com/google/uuid"
)

func (s *Database) Query_tableUsers() error {
	query := `
		create table if not exists users (
			id varchar(250) primary key,
			username varchar(250),
			password varchar(250),
			avatar varchar(250),
			name varchar(250),
			email varchar(250),
			phone varchar(250),
			role varchar(250),
			created timestamp,
			modified timestamp
		)
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *Database) Query_allUsers() ([]*User, error) {
	query := `select * from users`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user := &User{}

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

func (s *Database) Query_insertUsers(u *User) error {
	query := `
		insert into users (
			id,
			username,
			password,
			avatar,
			name,
			email,
			phone,
			role,
			created,
			modified 
		)
		values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10
		)
	`

	_, err := s.db.Query(query, &u.ID, &u.Username, &u.Password, &u.Avatar, &u.Name, &u.Email, &u.Phone, &u.Role, &u.Created, &u.Modified)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) Query_readUsers(id uuid.UUID) (*User, error) {
	query := `select * from users where id = $1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	user := &User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
		if err != nil {
			return nil, err
		}
	}

	return user, err
}

func (s *Database) Query_updateUsers(id uuid.UUID, u *User) error {
	query := `update users set username=$2, password=$3, avatar=$4, name=$5, email=$6, phone=$7, role=$8, modified=$9 where id = $1`

	_, err := s.db.Exec(query, id, &u.Username, &u.Password, &u.Avatar, &u.Name, &u.Email, &u.Phone, &u.Role, &u.Modified)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) Query_deleteUsers(id uuid.UUID) (uuid.UUID, error) {
	query := `delete from users where id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return id, err
	}

	return id, err
}
