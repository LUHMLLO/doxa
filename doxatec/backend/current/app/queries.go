package app

import (
	"doxapi/utils"
	"fmt"
	"log"
	"reflect"
)

func (s *Database) QueryList(entity string, t reflect.Type) (interface{}, error) {
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

func (s *Database) QueryCreate(entity string, params ...interface{}) {
	//query := fmt.Sprintf("sqls/%s/crud/create.sql", entity)
	log.Println(params...)
	//utils.ExecQL(s.db, query, params...)
}

func (s *Database) QueryRead(entity string, t reflect.Type, id string) (interface{}, error) {
	query := fmt.Sprintf("sqls/%s/crud/read.sql", entity)

	rows := utils.RowsQL(s.db, query, id)

	item := reflect.New(t).Elem()

	for rows.Next() {
		fields := make([]interface{}, item.NumField())

		for i := 0; i < item.NumField(); i++ {
			fields[i] = item.Field(i).Addr().Interface()
		}

		if err := rows.Scan(fields...); err != nil {
			return nil, err
		}
	}

	return item.Interface(), nil
}

func (s *Database) QueryUpdate(entity string, id string, params interface{}) (interface{}, error) {
	for key, value := range params.(map[string]interface{}) {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}

	return nil, nil
}

func (s *Database) QueryDelete(entity string, id string) {
	query := fmt.Sprintf("sqls/%s/crud/delete.sql", entity)

	utils.ExecQL(s.db, query, id)
}
