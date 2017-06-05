package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"fmt"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var GetToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims*/
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "mmerrick"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sk = []byte(Config.SigningKey)

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(sk)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})

//GetRedirect gets the stranslation from the keystore, then issues a redirect
var GetRedirect = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	j := vars["target"]

	log.Println("REDIRECT: key to use", j)
	y, err := Config.GET(db, j)
	if err != nil {
		writeStatus(w, err.Error(), false, 404)
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

var GetToken2 = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
	claims["exp"] = time.Now().Add(time.Hours * 24).Unix()

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

	// here we would validate apikey, but that's a TODO
	ut := MakeURLTranslation(u.URL)
	log.Println("POST: struct generated", ut)

	_, err = Config.GET(db, ut.Wordkey)
	// err == redis.Nil, then there is no key and we can add
	// run this check to preserve click count and creation time
	if err != nil {
		writeStatus(w, ut.NewURL, true, 200)
		Config.PUT(db, ut.Wordkey, ut)

	} else {
		// the url mapping already exists, but we'll retun the new url anyway
		writeStatus(w, ut.NewURL, true, 200)
	}
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.SigningKey), nil
	},
	EnableAuthOnOptions: true,
	SigningMethod:       jwt.SigningMethodHS256,
})
