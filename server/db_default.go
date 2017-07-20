package server

import (
	"errors"
	"log"
	"strings"
)

// Database is the overarching interface that all db's must match to
type Database interface {
	Get(string) (URLTranslation, error)
	Put(string, URLTranslation) error
	NewUser(string, string, string) error
	IsUser(user) (bool, error)
}

// ConnectDB creates a connection to a database, currently either Redis or Mongodb
func ConnectDB() error {
	var err error
	log.Println("Establishing database connection...")
	if strings.Compare(Config.DatabaseType, "Redis") == 0 {
		db, err = NewRedisDB(Config.DatabaseAddress, Config.DatabasePort)
		return err
	} else if strings.Compare(Config.DatabaseType, "MongoDB") == 0 {
		return errors.New("MongoDB Not supported yet")
	} else {
		return errors.New("database not correct type")
	}
}
