package app

import (
	"doxapi/utils"
	"fmt"
)

func (s *Postgres) QueryList(entity string) ([]*any, error) {
	query := fmt.Sprintf("sqls/%s/table/read.sql", entity)

	rows := utils.RowsQL(s.db, query)

	switch entity {
	case "client":
		clients := []*Client{}

		for rows.Next() {
			client := &Client{}

			if err := rows.Scan(
				&client.ID,
				&client.Name,
				&client.Email,
				&client.Phone,
				&client.Created,
				&client.Modified,
			); err != nil {
				return nil, err
			}

			clients = append(clients, client)
		}

		return clients, nil

	case "user":
		users := []*User{}

		for rows.Next() {
			user := &User{}

			if err := rows.Scan(
				&user.ID,
				&user.Avatar,
				&user.Username,
				&user.Password,
				&user.Role,
				&user.Created,
				&user.Modified,
				&user.Accessed,
			); err != nil {
				return nil, err
			}

			users = append(users, user)
		}

		return users, nil
	}

	return nil, nil
}
