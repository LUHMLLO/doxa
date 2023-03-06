package lib

func (storer *Database) Query_create_users_table() error {
	query := `
		create table if not exists users (
			id varchar(250) primary key,
			jwt varchar(250),
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

	_, err := storer.db.Exec(query)
	return err
}

func (storer *Database) Query_read_all_users_from_table() ([]*User, error) {
	query := `select * from users`

	rows, err := storer.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user := &User{}

		err := rows.Scan(&user.ID, &user.JWT, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}
