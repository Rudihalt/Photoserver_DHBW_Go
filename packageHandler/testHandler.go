package packageHandler

import "net/http"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><body>Test</body></html>"
	w.Write([]byte(responseString))
}
