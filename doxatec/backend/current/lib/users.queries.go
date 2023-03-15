package lib

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Database) users_init() error {
	schema := []string{
		"id varchar(250) primary key",
		"username varchar(250)",
		"password varchar(250)",
		"avatar varchar(250)",
		"name varchar(250)",
		"email varchar(250)",
		"phone varchar(250)",
		"role varchar(250)",
		"created timestamp",
		"modified timestamp",
	}

	query := fmt.Sprintf(`create table if not exists users (%s)`, StringToQuery(schema))

	_, err := s.db.Exec(query)
	return err
}

func (s *Database) users_beforeInsert(u *User) (*User, error) {
	query := `select * from users where (username = $1 OR email=$2 OR phone=$3)`

	rows, err := s.db.Query(query, u.Username, u.Email, u.Phone)
	if err != nil {
		return nil, err
	}

	user := &User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Avatar,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.Created,
			&user.Modified,
		); err != nil {
			return nil, err
		}

		if user.Username == u.Username {
			return nil, fmt.Errorf("username already in use")
		}

		if user.Email == u.Email {
			return nil, fmt.Errorf("email already in use")
		}

		if user.Phone == u.Phone {
			return nil, fmt.Errorf("phone already in use")
		}
	}

	return user, nil
}

func (s *Database) users_insert(u *User) error {
	schema := []string{
		"id",
		"username",
		"password",
		"avatar",
		"name",
		"email",
		"phone",
		"role",
		"created",
		"modified",
	}

	cols := []string{}
	for i := 0; i < len(schema); i++ {
		cols = append(cols, fmt.Sprintf("$%d", i+1))
	}

	query := fmt.Sprintf(`insert into users (%s) values (%s)`, StringToQuery(schema), StringToQuery(cols))

	if _, err := s.db.Query(
		query,
		&u.ID,
		&u.Username,
		&u.Password,
		&u.Avatar,
		&u.Name,
		&u.Email,
		&u.Phone,
		&u.Role,
		&u.Created,
		&u.Modified,
	); err != nil {
		return err
	}

	return nil
}

func (s *Database) users_readTable() ([]*User, error) {
	query := `select * from users`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*User{}

	for rows.Next() {
		user := &User{}

		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Avatar,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.Created,
			&user.Modified,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

func (s *Database) users_read(id uuid.UUID) (*User, error) {
	query := `select * from users where id=$1`

	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	user := &User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Avatar,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.Created,
			&user.Modified,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *Database) users_readCol(column string, value any) (*User, error) {
	query := fmt.Sprintf(`select * from users where %s=$1`, column)

	rows, err := s.db.Query(query, value)
	if err != nil {
		return nil, err
	}

	user := &User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Avatar,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.Created,
			&user.Modified,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *Database) users_update(id uuid.UUID, column string, value any) error {
	query := fmt.Sprintf(`update users set %s=$2 where id=$1`, column)

	_, err := s.db.Exec(query, id, value)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) users_delete(id uuid.UUID) error {
	query := `delete from users where id=$1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
func (s *Database) Users_beforeSignin(username, password string) (*User, error) {
	query := `select * from users where username = $1`

	rows, err := s.db.Query(query, username)
	if err != nil {
		return nil, err
	}

	user := &User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Avatar,
			&user.Name,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.Created,
			&user.Modified,
		); err != nil {
			return nil, err
		}
	}

	if user.Username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	check := CheckPasswordHash(password, user.Password)
	if !check {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}
