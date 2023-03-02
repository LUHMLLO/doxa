package lib

import (
	"fmt"
)

func (s *Database) Query_beforeInsertUsers(u *User) (*User, error) {
	query := `select * from users where (username = $1 OR email=$2 OR phone=$3)`

	rows, err := s.db.Query(query, u.Username, u.Email, u.Phone)
	if err != nil {
		return nil, err
	}

	user := &User{}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
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

func (s *Database) Query_beforeSigninUsers(u *SigninUser) (*User, error) {
	query := `select * from users where username = $1`

	rows, err := s.db.Query(query, u.Username)
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

	if user.Username == "" {
		return nil, fmt.Errorf("invalid username")
	}

	check := CheckPasswordHash(u.Password, user.Password)
	if !check {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}

// func (s *Database) Query_readUsers_Username(username string) (*User, error) {
// 	query := `select * from users where username = $1`

// 	rows, err := s.db.Query(query, username)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := &User{}
// 	for rows.Next() {
// 		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return user, err
// }

// func (s *Database) Query_readUsers_Email(email string) (*User, error) {
// 	query := `select * from users where email = $1`

// 	rows, err := s.db.Query(query, email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user := &User{}
// 	for rows.Next() {
// 		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Role, &user.Created, &user.Modified)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return user, err
// }
