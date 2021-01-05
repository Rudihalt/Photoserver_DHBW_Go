package packageHandler

import "net/http"

func RESTHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><body>Test</body></html>"
	w.Write([]byte(responseString))
}
