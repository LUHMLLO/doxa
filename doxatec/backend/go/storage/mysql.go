package storage

type MySQLStorage struct {
	username string
	password string
	hostname string
	dbname   string
}

func NewMySQLStorage(username string, password string, hostname string, dbname string) *MySQLStorage {
	return &MySQLStorage{
		username: username,
		password: password,
		hostname: hostname,
		dbname:   dbname,
	}
}
