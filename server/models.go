package server

import (
	"net/http"
	"time"
)

type user struct {
	Admin        bool   `json:"admin"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
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
	keyGenerator KeyGenerator
	api          *http.Server
	SigningKey   string `json:"signingkey"`
	WordsSHA256  string `json:"wordsSHA256"`
	TinyAddress  string `json:"tinyaddress"`
	BindAddress  string `json:"bindaddress"`
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
	URL string `json:"url"`
}
