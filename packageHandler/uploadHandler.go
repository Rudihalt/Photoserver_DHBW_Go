/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"net/http"
	"photoserver/packageObjects"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// get token from cookie
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		// if no cookie is set redirect to login site
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// get user from cookie token and set navigation bar with username
		user := packageObjects.GetUserByToken(cookie.Value)
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)
		// get images for user and send it to the template instead of nil
		err = UploadTemplate.Execute(w, NavData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
