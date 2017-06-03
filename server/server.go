package server

//Start keygen translation

// Sample test commands with curl:
// GET:
// curl 192.168.91.1:8000/BogWaspMark/json
// POST:
// curl -H "Content-Type: application/json" -X POST -d '{"apikey":"lolkey","url":"github.com/matmerr"}' http://192.168.91.1:8000/add

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/matmerr/gurl/keygen"
)

//Config for the primary shortener server
var Config serverConfig

//JSONObject is the base interface for serializing structs
type JSONObject interface{}

func (sc serverConfig) PUT(db database, key string, ut URLTranslation) error {
	return db.Put(key, ut)
}
func (sc serverConfig) GET(db database, key string) (URLTranslation, error) {
	return db.Get(key)
}

// because there are no generics, the global database is open to change,
// but as long as the db is a type which inherits from database{}, it's fine
// it would be cool if I could just use a base db type in the serverConfig for all
// database types, var db GenericDatabaseType
// #genericsplz
var db *RedisClient

// Start the server
func Start() {
	loadConfig()

	var err error
	db, err = NewRedisClient("192.168.91.137", 6379)

	if err != nil {
		log.Fatal(err)
	}

	Config.keyGenerator, _ = MakeKeyGenerator(3, 4, Config.TinyAddress)
	Config.api = &http.Server{
		Addr: Config.BindAddress + ":8000",
	}

	// setup URL shortener redirect point
	urlrtr := mux.NewRouter()
	urlrtr.Handle("/{target}", GetRedirect).Methods("GET")
	go func(rtr http.Handler) {
		log.Println("URL Redirect listening on", Config.api.Addr)
		log.Fatal(http.ListenAndServe(Config.api.Addr, rtr))
	}(urlrtr)

	// setup API shortner redirect
	apirtr := mux.NewRouter()
	apirtr.Handle("/user/login2", GetToken).Methods("GET")
	apirtr.Handle("/user/login", GetToken2).Methods("POST")
	apirtr.Handle("/token2", jwtMiddleware.Handler(GetToken)).Methods("GET")
	apirtr.Handle("/json", GetJSON).Methods("GET")
	apirtr.Handle("/add", PostTranslation).Methods("POST")
	go func(rtr http.Handler) {
		log.Println("API listening on", "0.0.0.0:8001")
		log.Fatal(http.ListenAndServe("0.0.0.0:8001", rtr))
	}(apirtr)

	webui := httprouter.New()
	webui.ServeFiles("/*filepath", http.Dir("client"))
	log.Fatal(http.ListenAndServe(":8081", webui))
}

//Shutdown here we send the signal to flush redis, and stop the webserver
func (sc *serverConfig) Shutdown() {
	log.Println("shutting down http server...")
	sc.Shutdown()
}

//MakeURLTranslation creates a URL translation type to be used for redis
func MakeURLTranslation(oldURL string) URLTranslation {
	key := Config.keyGenerator.GenerateKey(oldURL)
	size := len(Config.TinyAddress) + len(key)
	newURL := make([]byte, size)
	b := 0
	b += copy(newURL[b:], Config.TinyAddress)
	b += copy(newURL[b:], key)
	t := time.Now()
	return URLTranslation{key, oldURL, string(newURL), t, 0}
}

func writeStatus(w http.ResponseWriter, comment string, success bool, httpstatus int) {
	s := status{comment, success}
	j, _ := json.Marshal(s)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(httpstatus)
	fmt.Fprintf(w, "%s", j)
}

func writeJSON(w http.ResponseWriter, jo JSONObject, httpstatus int) {
	jb, _ := json.Marshal(jo)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(httpstatus)
	fmt.Fprintf(w, "%s", jb)
}
