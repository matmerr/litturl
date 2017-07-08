package server

import (
	"net/http"
	"time"
)

type user struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	Group        string `json:"group"`
}

type status struct {
	Comment string `json:"comment"`
	Success bool   `json:"success"`
}

type serverStatus struct {
	Comment string `json:"comment"`
	Ready   bool   `json:"ready"`
}

type serverConfig struct {
	keyGenerator    KeyGenerator
	apiServer       http.Server
	urlServer       http.Server
	SigningKey      []byte `json:"signingkey"`
	WordsSHA256     string `json:"wordsSHA256"`
	TinyAddress     string `json:"tinyaddress"`
	DatabaseType    string `json:"db_type"`
	DatabaseAddress string `json:"db_address"`
	DatabasePort    int    `json:"db_port"`
	bindAddress     string
}

//URLTranslation type for the value in redis
type URLTranslation struct {
	Wordkey string    `json:"translation"` // the lookup key for redis, "WordWordWord"
	OldURL  string    `json:"long"`        // https://reallylongoldurl.com/route1/subroute
	NewURL  string    `json:"short"`       // https://u.tiny.com/TestUrlNew
	Created time.Time `json:"time"`
	Clicks  int       `json:"clicks"`
}

//URLbase the datastructure to be used when POSTing to the api server
type URLbase struct {
	URL    string `json:"url"`
	Custom string `json:"custom"`
}
