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
		var cookie, err = r.Cookie("csrftoken")
		if cookie != nil {
			// already logged in
			http.Redirect(w, r, "/my", http.StatusSeeOther)
		}
		err = NavTemplate.Execute(w, nil)
		err = RegisterTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		confpassword := r.FormValue("confpassword")

		log.Println("username:", username, "password:", password, "confpassword:", confpassword)

		Register = RegisterData{}

		if password != confpassword {
			Register.PasswordNotCorrect = true
		}

		if packageObjects.UserExists(username) {
			Register.UserAlreadyExist = true
		}

		NavData = NavBarData{}

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
