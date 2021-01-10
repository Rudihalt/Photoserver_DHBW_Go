/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == "GET" {
		// set cookie expiration time to 1 nano second
		expiration := time.Now().Add(1)
		cookie := http.Cookie{Name: "csrftoken", Value: "-", Expires: expiration}
		// set cookie and redirect to login site
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
