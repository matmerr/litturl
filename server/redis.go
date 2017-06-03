package server

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

//RedisClient is the database structure type for Redis, which
//complies with the database{} interface
type RedisClient struct {
	Address string `json:"address"`
	client  *redis.Client
}

//NewRedisClient creates a new redis client
func NewRedisClient(host string, port int) (*RedisClient, error) {
	var red RedisClient
	red.client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: "", // no password
		DB:       0,  // use default DB
	})
	t, err := red.client.Ping().Result()
	fmt.Println(t)
	return &red, err
}

//Put adds the URLdata json to the key string in redis
func (r RedisClient) Put(key string, urldata URLTranslation) error {
	fmt.Println(urldata)
	bs, _ := json.Marshal(urldata)
	err := r.client.Set(key, bs, 0).Err()
	return err
}

//Get uses the key to return the URL translation
func (r RedisClient) Get(key string) (URLTranslation, error) {
	fmt.Println("Key to use: ", key)
	jsonresult, err := r.client.Get(key).Result()
	var u URLTranslation
	fmt.Println(u)
	if err != nil {
		return u, err
	}
	result := strings.NewReader(jsonresult)
	json.NewDecoder(result).Decode(&u)
	return u, nil
}

func (r RedisClient) NewUser(username, password string) error {
	u := user{false, username, password}
	bs, _ := json.Marshal(u)
	err := r.client.Set(u.Username, bs, 0).Err()
	return err
}

func (r RedisClient) IsUser(testuser user) (bool, error) {
	_, err := r.client.Get(testuser.Username).Result()
	if err != nil {
		return false, err
	}
	fmt.Println("if you like it then you should put a little crypto on it")
	return true, nil
}
