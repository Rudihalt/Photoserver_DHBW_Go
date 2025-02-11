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

func DiashowHandler(w http.ResponseWriter, r *http.Request) {
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

		// Idee, um die letzten hinzugefügten Fotos in Diashow anzuzuzeigen: Jan Dietzel
		allPhotos := packageObjects.GetAllPhotosByUser(user.Username)

		amountShowPhotos := 5

		var lastPhotos []packageObjects.Photo

		photoLength := len(*allPhotos)

		// get last 5 photos
		if photoLength < amountShowPhotos {
			lastPhotos = *allPhotos
		} else {
			lastPhotos = (*allPhotos)[photoLength-amountShowPhotos : photoLength]
		}

		// hand over the last 5 photos
		err = DiaShowTemplate.Execute(w, lastPhotos)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
