package server

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"log"

	"github.com/go-redis/redis"
)

//RedisDB is the database structure type for Redis, which
//complies with the database{} interface
type RedisDB struct {
	Address    string `json:"address"`
	urlClient  *redis.Client
	userClient *redis.Client
}

//NewRedisDB creates a new redis client
func NewRedisDB(host string, port int) (*RedisDB, error) {
	var red RedisDB
	red.urlClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "", // no password
		DB:       0,  // use url DB
	})
	t, err := red.urlClient.Ping().Result()

	red.userClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "", // no password
		DB:       0,  // use user DB
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(t)
	return &red, nil
}

//Put adds the URLdata json to the key string in redis
func (r RedisDB) Put(key string, urldata URLTranslation) error {
	bs, _ := json.Marshal(urldata)
	err := r.urlClient.Set(key, bs, 0).Err()
	return err
}

//Get uses the key to return the URL translation
func (r RedisDB) Get(key string) (URLTranslation, error) {

	jsonresult, err := r.urlClient.Get(key).Result()
	var u URLTranslation

	if err != nil {
		return u, err
	}
	resultReader := strings.NewReader(jsonresult)
	json.NewDecoder(resultReader).Decode(&u)
	return u, nil
}

// NewUser creates a new user in the db
func (r RedisDB) NewUser(username, password, group string) error {

	pwhash := hashPassword(password)
	u := user{username, pwhash, group}
	bs, _ := json.Marshal(u)
	if len(u.Username) == 0 {
		return errors.New("invalid username")
	}
	if len(u.PasswordHash) == 0 {
		return errors.New("invalid password")
	}
	err := r.userClient.Set(u.Username, bs, 0).Err()

	return err
}

// IsUser validates that the credentials are an actual use
func (r RedisDB) IsUser(testuser user) (bool, error) {
	jsonresult, err := r.userClient.Get(testuser.Username).Result()
	if err != nil {
		if err == redis.Nil {
			return false, errors.New("username or password incorrect")
		}
		return false, err
	}
	var storeduser user
	resultReader := strings.NewReader(jsonresult)
	json.NewDecoder(resultReader).Decode(&storeduser)

	if userDiff(storeduser, testuser) {
		return true, nil
	}
	return false, errors.New("username or password incorrect")
}
