package packageHandler

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}
