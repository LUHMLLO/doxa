package api

import (
	"fmt"
	"log"
)

// the idea here is to have general use functions

func CreateTable(table string, cols_vals []string) {
	query := fmt.Sprintf("create table if not exists %s (\n%s)", table, utils.StringToQuery(cols_vals))

	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertToTable(table string, cols_vals []interface{}) {
	query := fmt.Sprintf("insert into %s values (\n%s)", table, utils.InterfaceToQuery(cols_vals))

	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFromTable(table string, dataType []interface{}) ([]interface{}, error) {
	query := fmt.Sprintf("select * from %s", table)

	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	users := []*User{}
	for rows.Next() {
		user := &User{}

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Avatar, &user.Name, &user.Email, &user.Phone, &user.Created, &user.Modified)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}
