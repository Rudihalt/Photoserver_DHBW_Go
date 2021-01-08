package packageHandler

import (
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		err := NavTemplate.Execute(w, nil)
		err = LoginTemplate.Execute(w, nil)
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
		log.Println("username:", username, "password:", password)

		responseString := "<html><body>Successfully logged in as " + username + "</body></html>"
		_, err := w.Write([]byte(responseString))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
