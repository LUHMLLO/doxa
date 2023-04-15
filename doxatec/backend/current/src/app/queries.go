package app

import (
	"doxapi/utils"
	"fmt"
	"reflect"
)

func (s *Postgres) QueryList(entity string, t reflect.Type) (interface{}, error) {
	query := fmt.Sprintf("sqls/%s/table/read.sql", entity)

	rows := utils.RowsQL(s.db, query)

	slice := reflect.MakeSlice(reflect.SliceOf(t), 0, 0)

	for rows.Next() {
		item := reflect.New(t).Elem()

		fields := make([]interface{}, item.NumField())

		for i := 0; i < item.NumField(); i++ {
			fields[i] = item.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fields...); err != nil {
			return nil, err
		}

		slice = reflect.Append(slice, item)
	}

	return slice.Interface(), nil
}

func (s *Postgres) QueryCreate(entity string, values ...interface{}) error {
	query := fmt.Sprintf("sqls/%s/crud/create.sql", entity)

	result := utils.ExecQL(s.db, query, values...)

	_, err := result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
