package server

// Database is the overarching interface that all db's must match to
type Database interface {
	Get(string) (URLTranslation, error)
	Put(string, URLTranslation) error
	NewUser(string, string) error
	IsUser(user) (bool, error)
}

// CreateDB creates a connection to a database, currently either Redis or Mongodb
func CreateDB(host string, port int, username, password string) Database {

	var db Database = NewRedisdb(host, port)
	return db
}
