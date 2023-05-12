package utils

import (
	"errors"
	"log"
	"reflect"
	"sort"
)

func SortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func ReorderMap(m map[string]interface{}, t reflect.Type) (map[string]interface{}, error) {
	if t.Kind() != reflect.Struct {
		return nil, errors.New("provided type is not a struct")
	}

	var newMap map[string]interface{}

	log.Println(m)
	log.Println(t)

	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		if value, ok := m[fieldName]; ok {
			newMap[fieldName] = value
		}
	}

	log.Println("new map: ", newMap)

	return newMap, nil
}
