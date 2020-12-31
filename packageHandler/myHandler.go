package packageHandler

import "net/http"

func MyHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><body>My</body></html>"
	w.Write([]byte(responseString))
}
