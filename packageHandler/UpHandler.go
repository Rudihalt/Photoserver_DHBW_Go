package packageHandler

import "net/http"

func UpHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/up.html")
}
