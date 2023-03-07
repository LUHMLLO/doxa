package lib

// func (s *Database) Query_create_users_table() error {
// 	query := `
// 		create table if not exists users (
// 			id varchar(250) primary key,
// 			jwt varchar(250),
// 			username varchar(250),
// 			password varchar(250),
// 			avatar varchar(250),
// 			name varchar(250),
// 			email varchar(250),
// 			phone varchar(250),
// 			role varchar(250),
// 			created timestamp,
// 			modified timestamp
// 		)
// 	`

// 	_, err := s.db.Exec(query)
// 	return err
// }

// func (s *Database) Query_read_all_users_from_table() ([]*User, error) {
// 	query := `select * from users`

// 	rows, err := s.db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}

// 	users := []*User{}

// 	for rows.Next() {
// 		user := &User{}

// 		if err := rows.Scan(
// 			&user.ID,
// 			&user.JWT,
// 			&user.Username,
// 			&user.Password,
// 			&user.Avatar,
// 			&user.Name,
// 			&user.Email,
// 			&user.Phone,
// 			&user.Role,
// 			&user.Created,
// 			&user.Modified,
// 		); err != nil {
// 			return nil, err
// 		}

// 		users = append(users, user)
// 	}

// 	return users, err
// }

// func (s *Database) Query_before_insert_user(u *User) (*User, error) {
// 	query := `select * from users where (username = $1 OR email=$2 OR phone=$3)`

// 	rows, err := s.db.Query(query, u.Username, u.Email, u.Phone)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := &User{}

// 	for rows.Next() {
// 		if err := rows.Scan(
// 			&user.ID,
// 			&user.JWT,
// 			&user.Username,
// 			&user.Password,
// 			&user.Avatar,
// 			&user.Name,
// 			&user.Email,
// 			&user.Phone,
// 			&user.Role,
// 			&user.Created,
// 			&user.Modified,
// 		); err != nil {
// 			return nil, err
// 		}

// 		if user.Username == u.Username {
// 			return nil, fmt.Errorf("username already in use")
// 		}

// 		if user.Email == u.Email {
// 			return nil, fmt.Errorf("email already in use")
// 		}

// 		if user.Phone == u.Phone {
// 			return nil, fmt.Errorf("phone already in use")
// 		}
// 	}

// 	return user, nil
// }

// func (s *Database) Query_insert_user(u *User) error {
// 	query := `
// 		insert into users (
// 			id,
// 			jwt,
// 			username,
// 			password,
// 			avatar,
// 			name,
// 			email,
// 			phone,
// 			role,
// 			created,
// 			modified
// 		)
// 		values (
// 			$1,
// 			$2,
// 			$3,
// 			$4,
// 			$5,
// 			$6,
// 			$7,
// 			$8,
// 			$9,
// 			$10,
// 			$11
// 		)
// 	`

// 	if _, err := s.db.Query(
// 		query,
// 		&u.ID,
// 		&u.JWT,
// 		&u.Username,
// 		&u.Password,
// 		&u.Avatar,
// 		&u.Name,
// 		&u.Email,
// 		&u.Phone,
// 		&u.Role,
// 		&u.Created,
// 		&u.Modified,
// 	); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Database) Query_read_user_where_ID_and_column(id uuid.UUID, column string, param any) error {
// 	query := fmt.Sprintf(`select * from users where (id=$1 AND %s=$2)`, column)

// 	_, err := s.db.Exec(query, id, param)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Database) Query_read_user_where_column(column, param string) (*User, error) {
// 	query := fmt.Sprintf(`select * from users where %s=$1`, column)

// 	_, err := s.db.Exec(query, param)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, nil

// }

// func (s *Database) Query_update_user_column_where_ID(id uuid.UUID, column string, param any) error {
// 	query := fmt.Sprintf(`update users set %s=$2 where id=$1`, column)

// 	_, err := s.db.Exec(query, id, param)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (s *Database) Query_delete_user_where_ID(id uuid.UUID) (uuid.UUID, error) {
// 	query := `delete from users where id = $1`

// 	_, err := s.db.Exec(query, id)
// 	if err != nil {
// 		return id, err
// 	}

// 	return id, err
// }
