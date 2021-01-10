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

var Login LoginData

type LoginData struct {
	PasswordNotCorrect bool
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		// get token from cookie
		var cookie, _ = r.Cookie("csrftoken")
		if cookie != nil {
			// if cookie exist redirect user to gallery
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		} else {
			// else user is not logged in
			err := NavTemplate.Execute(w, nil)
			err = LoginTemplate.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else if r.Method == "POST" {
		// parse the form
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
			return
		}
		// initialize username and password
		username := r.FormValue("username")
		password := r.FormValue("password")
		log.Println("username:", username, "password:", password)

		// check if password is correct
		ok, token := packageObjects.CheckPassword(username, password)
		if ok {
			// set cross site request forgery token and expiration time to 24 hours
			expiration := time.Now().Add(24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: token, Expires: expiration}
			// set cookie and redirect to gallery
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		} else {
			// set PasswordNotCorrect bool and hand it over to the login template where the message will be shown
			// that the password is not correct
			Login = LoginData{PasswordNotCorrect: true}
			err := NavTemplate.Execute(w, nil)
			err = LoginTemplate.Execute(w, Login)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
