package server

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"crypto/rsa"

	"crypto/x509"

	"github.com/gorilla/mux"
)

func loadConfig(webFiles string) {
	bytes, err := ioutil.ReadFile("conf/server-config.json")

	// error reading config, open api for generation
	if err != nil {
		createConfig(webFiles)
	} else {
		// file loaded, error marshalling config, open api for generation
		err = json.Unmarshal(bytes, &Config)
		if err != nil {
			createConfig(webFiles)
		}
		if strings.Compare(Config.WordsSHA256, HashWords("conf/nounlist.txt")) != 0 {
			log.Fatal("Word reference file has changed!")
		}
	}
}

func saveConfig() {
	bytes, err := json.MarshalIndent(Config, "", "  ")
	config := "conf/server-config.json"
	//fmt.Println(Config)
	err = ioutil.WriteFile(config, bytes, 0644)
	if err != nil {
		f, _ := os.Create(config)
		defer f.Close()
		err = ioutil.WriteFile(config, bytes, 0644)
	}
}

// sample curl command
// curl -H "Content-Type: application/json" -X POST -d '{"signingkey":"19","tinyaddress":"http://127.0.0.1/", "bindaddress":"0.0.0.0" }' http://192.168.91.1:8000/config
func createConfig(webFiles string) {
	log.Println("no config was found")

	// signal to this channel to stop config server
	stopserver = make(chan bool)

	SetServerStatus("config missing", false)

	apirtr := mux.NewRouter()
	// serve up the static
	apirtr.PathPrefix("/static").Handler(http.FileServer(http.Dir(webFiles)))
	apirtr.Handle("/ui", http.HandlerFunc(IndexHandler(webFiles+"index.html")))
	apirtr.Handle("/ui/{page}", http.HandlerFunc(IndexHandler(webFiles+"index.html")))

	apirtr.Handle("/api/status", GetStatus).Methods("GET")
	apirtr.Handle("/api/config", PostConfig).Methods("POST")
	apirtr.Handle("/", defaultRedirect).Methods("GET")
	apirtr.Handle("{page}", defaultRedirect).Methods("GET")

	srv := http.Server{Addr: ":8001", Handler: apirtr}

	fmt.Println("opening config api on :8001")

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// wait until stop server signal is received
	<-stopserver

	log.Println("config loaded")
	SetServerStatus("config loaded", true)

	// wait 10 seconds before shutting config server down
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// ensure all paths cancel, avoid context leak
	defer cancel()

	srv.Shutdown(ctx)

	// generate a hash for the words list, and save to config
	Config.WordsSHA256 = HashWords("conf/nounlist.txt")

	// generate a signing key for the jwt tokens.
	// NOTE: secrets don't have to be used with asymmetric keys,
	// they can just be long, but this rsa lib is easy to use

	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	Config.SigningKey = x509.MarshalPKCS1PrivateKey(key)
	// flush to disk
	saveConfig()

}

var defaultRedirect = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui", http.StatusMovedPermanently)
})

//PostConfig converts adds a translation to the store
var PostConfig = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type baseConfig struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		Group           string `json:"group"`
		TinyAddress     string `json:"tinyaddress"`
		DatabaseType    string `json:"db_type"`
		DatabaseAddress string `json:"db_address"`
		DatabasePort    int    `json:"db_port"`
	}
	var init baseConfig
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&init)
	if err != nil {
		fmt.Println("THIS ERROR")
		fmt.Println(err)
		writeStatus(w, "invalid config", false, 200)
	} else {
		// if the tiny address supplied doesn't have a trailing /, add one on
		// required for returning the url and keys
		if (init.TinyAddress[len(init.TinyAddress)-1]) != '/' {
			init.TinyAddress = init.TinyAddress + "/"
		}
		fmt.Println(init)
		Config.bindAddress = "0.0.0.0"
		Config.TinyAddress = init.TinyAddress
		Config.DatabaseType = init.DatabaseType
		Config.DatabaseAddress = init.DatabaseAddress
		Config.DatabasePort = init.DatabasePort

		// validate the connection to the database
		err := ConnectDB()
		if err != nil {
			log.Fatal(err)
			writeStatus(w, err.Error(), true, 200)
			return
		}

		// at this point the connection to the db has been established,
		// let's create the supplied user
		db.NewUser(init.Username, init.Password, init.Group)

		writeStatus(w, "successfully loaded config", true, 200)
		stopserver <- true
	}
})

// HashWords ok
func HashWords(fpath string) string {
	hasher := sha256.New()
	fullpath, err := filepath.Abs(fpath)
	f, err := os.Open(fullpath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(hasher, f)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
