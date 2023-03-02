package lib

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
