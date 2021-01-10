/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"log"
	"net/http"
	"photoserver/packageObjects"
	"time"
)

var Register RegisterData

type RegisterData struct {
	SuccessfullyRegistered bool
	UserAlreadyExist       bool
	PasswordNotCorrect     bool
	Username               string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		// get token from cookie
		var cookie, err = r.Cookie("csrftoken")
		if cookie != nil {
			// already logged in redirect to gallery
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		}
		err = NavTemplate.Execute(w, nil)
		err = RegisterTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		// parse the form
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
			return
		}
		// initialize username and password and confirm password
		username := r.FormValue("username")
		password := r.FormValue("password")
		confpassword := r.FormValue("confpassword")

		log.Println("username:", username, "password:", password, "confpassword:", confpassword)

		Register = RegisterData{}

		// check if password and confpassword are equal
		if password != confpassword {
			// if not set PasswordNotCorrect to true -> this will show a message
			// which said that the passwords are not equal
			Register.PasswordNotCorrect = true
		}

		// if user already exist the UserAlreadyExist value will set to true
		// this will also pop up a message for the user in the frontend
		if packageObjects.UserExists(username) {
			Register.UserAlreadyExist = true
		}

		NavData = NavBarData{}

		// if everything is correct the user will be created and the cookie will be set
		// with an expiration of 24 hours
		if !Register.UserAlreadyExist && !Register.PasswordNotCorrect {
			user := packageObjects.CreateUser(username, password)
			Register.SuccessfullyRegistered = true
			Register.Username = user.Username
			// https://stackoverflow.com/questions/12130582/setting-cookies-with-net-http
			expiration := time.Now().Add(24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: user.Token, Expires: expiration}
			http.SetCookie(w, &cookie)
			NavData.Username = user.Username
		}
		err := NavTemplate.Execute(w, NavData)

		err = RegisterTemplate.Execute(w, Register)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
