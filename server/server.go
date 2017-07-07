package server

//Start keygen translation

// Sample test commands with curl:
// GET:
// curl 192.168.91.1:8000/BogWaspMark/json
// POST:
// curl -H "Content-Type: application/json" -X POST -d '{"apikey":"lolkey","url":"github.com/matmerr"}' http://192.168.91.1:8000/add

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

//ServerStatus is the type which is sent when /status is requested on the api port
var ServerStatus serverStatus

//Config for the primary shortener server
var Config serverConfig

//JSONObject is the base interface for serializing structs
type JSONObject interface{}

func (sc serverConfig) PUT(db Database, key string, ut URLTranslation) error {
	return db.Put(key, ut)
}
func (sc serverConfig) GET(db Database, key string) (URLTranslation, error) {
	return db.Get(key)
}

// this channel is reserved for stopping webservers
var stopserver chan bool

var db Database

// Start the server
func Start() {
	loadConfig()

	// if there is no existing connection to the database
	// reloading config from cold start, set it up
	if db == nil {
		err := ConnectDB()
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// create the keygenerator used for
	Config.keyGenerator, _ = MakeKeyGenerator(3, 4, Config.TinyAddress)

	// setup URL shortener redirect point
	urlrtr := mux.NewRouter()
	urlrtr.Handle("/{target}", GetRedirect).Methods("GET")

	Config.urlServer = http.Server{
		Addr:    Config.bindAddress + ":8000",
		Handler: urlrtr,
	}

	// start the URL server in a seperate goroutine
	go func() {
		log.Println("URL Redirect listening on", Config.urlServer.Addr)
		if err := Config.urlServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// setup API shortner redirect
	SetServerStatus("server ready", true)
	apirtr := mux.NewRouter()
	apirtr.Handle("/settings", GetSettings).Methods("GET")
	dir, _ := os.Executable()
	fmt.Println(dir)

	// serve up the static
	apirtr.PathPrefix("/static").Handler(http.FileServer(http.Dir("client/dist")))
	apirtr.Handle("/ui", http.HandlerFunc(IndexHandler("client/dist/index.html")))
	apirtr.Handle("/ui/{page}", http.HandlerFunc(IndexHandler("client/dist/index.html")))

	apirtr.Handle("/api/settings", PostSettings).Methods("POST")
	apirtr.Handle("/api/status", GetStatus).Methods("GET")
	apirtr.Handle("/api/user/login", UserLogin).Methods("POST")
	apirtr.Handle("/api/url/add", jwtMiddleware.Handler(PostTranslation)).Methods("POST")
	apirtr.Handle("/{target}", GetRedirect).Methods("GET")
	log.Println("API listening on", "0.0.0.0:8001")

	Config.apiServer = http.Server{
		Addr:    "0.0.0.0:8001",
		Handler: apirtr,
	}

	go func() {
		if err := Config.apiServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// wait until stop server signal is received
	<-stopserver

	// wait 10 seconds before shutting config server down
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// ensure all paths cancel, avoid context leak
	defer cancel()

	Config.apiServer.Shutdown(ctx)
	Config.urlServer.Shutdown(ctx)
	return
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

//SetServerStatus accepts a string describing a state, and if that state is ready
func SetServerStatus(newcomment string, ready bool) {
	ServerStatus.Comment = newcomment
	ServerStatus.Ready = ready
}

func writeJSON(w http.ResponseWriter, jo JSONObject, httpstatus int) {
	jb, _ := json.Marshal(jo)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(httpstatus)
	fmt.Fprintf(w, "%s", jb)
}
