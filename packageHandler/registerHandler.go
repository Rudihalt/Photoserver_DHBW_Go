package packageHandler

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><body>Register</body></html>"
	w.Write([]byte(responseString))
}
