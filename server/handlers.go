package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//GetRedirect gets the stranslation from the keystore, then issues a redirect
var GetRedirect = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	j := vars["target"]

	log.Println("REDIRECT: key to use", j)
	y, err := Config.GET(db, j)
	if err != nil {
		writeStatus(w, err.Error(), false, 404)
		// http.Redirect(w, r, "/ui", http.StatusMovedPermanently)
		return
	}

	http.Redirect(w, r, y.OldURL, http.StatusMovedPermanently)
	go func() {
		y.Clicks++
		Config.PUT(db, y.Wordkey, y)
	}()
})

//GetJSON returns the URLTranslation JSON object
var GetJSON = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["target"]

	ut, err := Config.GET(db, t)
	if err != nil {
		writeStatus(w, "key does not exist", false, 404)
		return
	} else if err != nil {
		writeStatus(w, err.Error(), false, 500)
		return
	}
	writeJSON(w, ut, 200)
})

var UserLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u user
	err := decoder.Decode(&u)
	log.Println("POST: User login: ", u)
	if err != nil {
		log.Println(err)
	}

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims*/
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "mmerrick"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signingkey := []byte(Config.SigningKey)
	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(signingkey)

	fmt.Println(tokenString)
	/* Finally, write the token to the browser window */

	type login struct {
		ID     string `json:"id_token"`
		Access string `json:"access_token"`
	}

	session := login{"not_implemented", tokenString}
	jdata, _ := json.Marshal(session)

	w.Write(jdata)
})

//PostTranslation converts adds a translation to the store
var PostTranslation = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u URLbase
	err := decoder.Decode(&u)
	log.Println("POST: long url:", u.URL)
	json.NewDecoder(r.Body).Decode(&u)

	if strings.HasPrefix(u.URL, "http") == false {
		u.URL = "http://" + u.URL
	}

	uTranslation := MakeURLTranslation(u.URL)

	// if we have a custom URL specified, we use that instead
	if len(u.Custom) > 0 {
		uTranslation.Wordkey = u.Custom
	}

	log.Println("POST: struct generated", uTranslation)

	_, err = Config.GET(db, uTranslation.Wordkey)
	// err == redis.Nil, then there is no key and we can add
	// run this check to preserve click count and creation time
	if err != nil {
		writeStatus(w, uTranslation.NewURL, true, 200)
		Config.PUT(db, uTranslation.Wordkey, uTranslation)

	} else {
		// the url mapping already exists, but we'll retun the new url anyway
		writeStatus(w, uTranslation.NewURL, true, 200)
	}
})

//GetStatus handler writes the current server status indicated in the global server status struct
var GetStatus = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(ServerStatus)
	fmt.Fprintf(w, "%s", j)
})

// GetSettings will expose all public settings
var GetSettings = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// Normally we could just serialize the Config, but it contains net/http which has a mutex lock
	type settings struct {
		WordsSHA256     string `json:"wordsSHA256"`
		TinyAddress     string `json:"tinyaddress"`
		DatabaseType    string `json:"db_type"`
		DatabaseAddress string `json:"db_address"`
		DatabasePort    int    `json:"db_port"`
	}

	var settmp settings
	settmp.WordsSHA256 = Config.WordsSHA256
	settmp.TinyAddress = Config.TinyAddress
	settmp.DatabaseType = Config.DatabaseType
	settmp.DatabaseAddress = Config.DatabaseAddress
	settmp.DatabasePort = Config.DatabasePort
	fmt.Println(settmp)

	j, _ := json.Marshal(settmp)
	fmt.Fprintf(w, "%s", j)
})

var PostSettings = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// Normally we could just serialize the Config, but it contains net/http which has a mutex lock
	type settings struct {
		WordsSHA256     string `json:"wordsSHA256"`
		TinyAddress     string `json:"tinyaddress"`
		DatabaseType    string `json:"db_type"`
		DatabaseAddress string `json:"db_address"`
		DatabasePort    int    `json:"db_port"`
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var settmp settings
	err := decoder.Decode(&settmp)
	if err != nil {
		writeStatus(w, error.Error(err), false, 500)
	} else {
		fmt.Println(settmp)
		Config.TinyAddress = settmp.TinyAddress
		writeStatus(w, "successfully updated config", true, 200)
	}
})

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.SigningKey), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
