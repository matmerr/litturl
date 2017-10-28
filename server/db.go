package server

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/matmerr/dataface"
)

// ConnectDB creates a connection to a database, currently either Redis or Mongodb
func ConnectDB() error {
	var err error
	log.Println("Establishing database connection...")
	if strings.Compare(Config.DatabaseType, "Redis") == 0 {
		//db, err = NewRedisDB(Config.DatabaseAddress, Config.DatabasePort)
		db, err = dataface.InitializeDatabase("redis", Config.DatabaseAddress, Config.DatabasePort, "", "")
		return err
	} else if strings.Compare(Config.DatabaseType, "MongoDB") == 0 {
		db, err = dataface.InitializeDatabase("mongo", Config.DatabaseAddress, Config.DatabasePort, "", "")
	} else {
		return errors.New("database not correct type")
	}
	return err
}

func getURLTranslation(key string) (*URLTranslation, error) {
	var u URLTranslation
	result, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	resultReader := strings.NewReader(string(result))
	json.NewDecoder(resultReader).Decode(&u)
	return &u, nil
}
func putURLTranslation(key string, urldata *URLTranslation) error {
	bs, _ := json.Marshal(urldata)
	err := db.Put(key, bs)
	return err
}

func newUser(username, password, group string) error {

	pwhash := hashPassword(password)
	u := user{username, pwhash, group}
	bs, _ := json.Marshal(u)
	if len(u.Username) == 0 {
		return errors.New("invalid username")
	}
	if len(u.PasswordHash) == 0 {
		return errors.New("invalid password")
	}
	err := db.Put(u.Username, bs)

	return err
}

// IsUser validates that the credentials are an actual use
func isUser(testuser user) (bool, error) {
	jsonresult, err := db.Get(testuser.Username)
	if err != nil {
		// need to fix
		return false, errors.New("username or password incorrect")

	}
	var storeduser user
	resultReader := strings.NewReader(string(jsonresult))
	json.NewDecoder(resultReader).Decode(&storeduser)

	if userDiff(storeduser, testuser) {
		return true, nil
	}
	return false, errors.New("username or password incorrect")
}
