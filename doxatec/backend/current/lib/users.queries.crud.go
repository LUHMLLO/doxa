package lib

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Database) Query_insert_user(u *User) error {
	query := `
		insert into users (
			id,
			jwt,
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
			$10,
			$11
		)
	`

	if _, err := s.db.Query(
		query,
		&u.ID,
		&u.JWT,
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

func (s *Database) Query_read_user_where_column(column, param string) (*User, error) {
	query := fmt.Sprintf(`select * from users where %s = $1`, column)

	rows, err := s.db.Query(query, param)
	if err != nil {
		return nil, err
	}

	user := &User{}

	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.JWT,
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

	return user, err
}

func (s *Database) Query_update_user_column_where_ID(id uuid.UUID, column string, param any) error {
	query := fmt.Sprintf(`update users set %s=$2 where id=$1`, column)

	_, err := s.db.Exec(query, id, param)
	if err != nil {
		return err
	}

	return nil
}

func (s *Database) Query_delete_user_where_ID(id uuid.UUID) (uuid.UUID, error) {
	query := `delete from users where id = $1`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return id, err
	}

	return id, err
}
