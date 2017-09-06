package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"

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
		http.Redirect(w, r, "/ui", http.StatusMovedPermanently)
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

// UserLogin calidate user, and returns jwt token if valid
var UserLogin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var u user
	err := decoder.Decode(&u)
	if err != nil {
		log.Println(err)
		return
	}

	// check if the user is legit
	isUser, err := db.IsUser(u)
	if err != nil && !isUser {
		writeStatus(w, error.Error(err), false, 501)
		return
	}

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims*/
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 8760).Unix()

	signingkey := []byte(Config.SigningKey)
	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(signingkey)

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
		uTranslation.NewURL = Config.TinyAddress + u.Custom
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

	j, _ := json.Marshal(settmp)
	fmt.Fprintf(w, "%s", j)
})

// PostSettings event handler
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

		if (settmp.TinyAddress[len(settmp.TinyAddress)-1]) != '/' {
			Config.TinyAddress = settmp.TinyAddress + "/"
		} else {
			Config.TinyAddress = settmp.TinyAddress
		}
		writeStatus(w, "successfully updated config", true, 200)
		saveConfig()
	}
})

// IndexHandler serves files out of the specified client directory
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

func userDiff(user1, user2 user) bool {
	if strings.Compare(user1.Username, user2.Username) == 0 && strings.Compare(user1.Group, user2.Group) == 0 {
		err := bcrypt.CompareHashAndPassword([]byte(user1.PasswordHash), []byte(user2.PasswordHash))
		if err != nil {
			log.Println(err)
		} else{
			return true
		}
	}
	return false
}

func userGroupCheck(user1, user2 user) bool {
	if (strings.Compare(user1.Group, user2.Group) == 0){
		return true
	}
	return false
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
