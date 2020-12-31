package packageHandler

import "net/http"

func DiashowHandler(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><body>Diashow</body></html>"
	w.Write([]byte(responseString))
}
