package utils

import (
	"database/sql"
	"doxatec/types"
)

func ScanIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.Avatar, &user.Username, &user.Password, &user.Customer, &user.Created, &user.Modified, &user.Accessed)

	return user, err
}
