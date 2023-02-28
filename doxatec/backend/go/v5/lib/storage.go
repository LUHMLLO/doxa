package lib

type Storage interface {
	CreateTable(table string, params []string) error
	InsertToTable(table string, params []interface{}) error
	ReadFromTable(table string) error
}
