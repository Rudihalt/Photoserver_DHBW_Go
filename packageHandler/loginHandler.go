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
		var cookie, _ = r.Cookie("csrftoken")
		if cookie != nil {
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		} else {
			err := NavTemplate.Execute(w, nil)
			err = LoginTemplate.Execute(w, nil)
			if err != nil {
				// already logged in
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		log.Println("username:", username, "password:", password)

		ok, token := packageObjects.CheckPassword(username, password)
		if ok {
			expiration := time.Now().Add(24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: token, Expires: expiration}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		} else {
			Login = LoginData{PasswordNotCorrect: true}
			err := NavTemplate.Execute(w, nil)
			err = LoginTemplate.Execute(w, Login)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
