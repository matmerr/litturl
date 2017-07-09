package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"log"

	"github.com/go-redis/redis"
)

//redisdb is the database structure type for Redis, which
//complies with the database{} interface
type redisdb struct {
	Address     string `json:"address"`
	url_client  *redis.Client
	user_client *redis.Client
}

//NewRedisdb creates a new redis client
func NewRedisdb(host string, port int) (*redisdb, error) {
	var red redisdb
	red.url_client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "", // no password
		DB:       0,  // use url DB
	})
	t, err := red.url_client.Ping().Result()

	red.user_client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "", // no password
		DB:       0,  // use user DB
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(t)
	return &red, nil
}

//Put adds the URLdata json to the key string in redis
func (r redisdb) Put(key string, urldata URLTranslation) error {
	fmt.Println(urldata)
	bs, _ := json.Marshal(urldata)
	err := r.url_client.Set(key, bs, 0).Err()
	return err
}

//Get uses the key to return the URL translation
func (r redisdb) Get(key string) (URLTranslation, error) {
	fmt.Println("Key to use: ", key)
	jsonresult, err := r.url_client.Get(key).Result()
	var u URLTranslation
	fmt.Println(u)
	if err != nil {
		return u, err
	}
	resultReader := strings.NewReader(jsonresult)
	json.NewDecoder(resultReader).Decode(&u)
	return u, nil
}

func (r redisdb) NewUser(username, password, group string) error {
	u := user{username, password, group}
	bs, _ := json.Marshal(u)
	if len(u.Username) == 0 {
		return errors.New("invalid username")
	}
	if len(u.PasswordHash) == 0 {
		return errors.New("invalid password")
	}
	err := r.user_client.Set(u.Username, bs, 0).Err()
	fmt.Println(u)
	return err
}

func (r redisdb) IsUser(testuser user) (bool, error) {
	jsonresult, err := r.user_client.Get(testuser.Username).Result()
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
