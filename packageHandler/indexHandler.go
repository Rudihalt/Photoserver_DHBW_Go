package packageHandler

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	if len(name) == 0 {
		name = "World"
	}
	responseString := "<html><body>Index Page<br>Hello " + name + "</body></html>"
	w.Write([]byte(responseString)) // unbedingt Templates verwenden!
}
