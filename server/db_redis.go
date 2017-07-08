package server

import (
	"encoding/json"
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
		DB:       1,  // use user DB
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	return &red, err
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
	result := strings.NewReader(jsonresult)
	json.NewDecoder(result).Decode(&u)
	return u, nil
}

func (r redisdb) NewUser(username, password string) error {
	u := user{false, username, password}
	bs, _ := json.Marshal(u)
	err := r.user_client.Set(u.Username, bs, 1).Err()
	return err
}

func (r redisdb) IsUser(testuser user) (bool, error) {
	_, err := r.user_client.Get(testuser.Username).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
