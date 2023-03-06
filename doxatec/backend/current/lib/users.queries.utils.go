package lib

import "fmt"

func (storer *Database) Query_before_insert_user(u *User) (*User, error) {
	query := `select * from users where (jwt = $1 OR username = $2 OR email=$3 OR phone=$4)`

	rows, err := storer.db.Query(query, &u.JWT, u.Username, u.Email, u.Phone)
	if err != nil {
		return nil, err
	}

	user := &User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.JWT, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
		if err != nil {
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

func (storer *Database) Query_before_signin_user(username, password string) (*User, error) {
	query := `select * from users where username = $1`

	rows, err := storer.db.Query(query, username)
	if err != nil {
		return nil, err
	}

	user := &User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.JWT, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
		if err != nil {
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
