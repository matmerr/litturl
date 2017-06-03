package server

import (
	"context"
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

	"github.com/julienschmidt/httprouter"
)

func loadConfig() {
	bytes, err := ioutil.ReadFile("conf/server-config.json")

	// error reading config, open api for generation
	if err != nil {
		createConfig()
	} else {
		// file loaded, error marshalling config, open api for generation
		err = json.Unmarshal(bytes, &Config)
		if err != nil {
			createConfig()
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

var stopserver chan bool

// sample curl command
// curl -H "Content-Type: application/json" -X POST -d '{"signingkey":"19","tinyaddress":"http://127.0.0.1/", "bindaddress":"0.0.0.0" }' http://192.168.91.1:8000/config
func createConfig() {
	log.Println("no config was found")

	// signal to this channel to stop config server
	stopserver = make(chan bool)

	api := httprouter.New()
	api.POST("/config", PostConfig)

	srv := http.Server{Addr: ":8000", Handler: api}
	fmt.Println("opening config api on", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// wait until stop server signal is received
	<-stopserver

	log.Println("config accepted")

	// wait 10 seconds before shutting config server down
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// ensure all paths cancel, avoid context leak
	defer cancel()

	srv.Shutdown(ctx)

	// generate a hash for the words list, and save to config
	Config.WordsSHA256 = HashWords("conf/nounlist.txt")

	// flush to disk
	saveConfig()

}

//PostConfig converts adds a translation to the store
func PostConfig(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err)
		writeStatus(w, "invalid config", false, 200)
	} else {
		// if the tiny address supplied doesn't have a trailing /, add one on
		// required for returning the url and keys
		if (Config.TinyAddress[len(Config.TinyAddress)-1]) != '/' {
			Config.TinyAddress = Config.TinyAddress + "/"
		}
		writeStatus(w, "successfully loaded config", true, 200)
		stopserver <- true
	}
}

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
