package packageHandler

import "net/http"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><body>Login</body></html>"
	w.Write([]byte(responseString))
}
