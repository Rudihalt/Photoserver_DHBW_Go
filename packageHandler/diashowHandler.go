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

type testdreck struct {
	Path string
}

func DiashowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		user := packageObjects.GetUserByToken(cookie.Value)
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)


		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	var anus = []testdreck{
		{
			Path: "/images/p1.jpg",
		},
		{
			Path: "/images/p2.jpg",
		},
		{
			Path: "/images/p3.jpg",
		},
	}


	err := DiashowTemplate.Execute(w, anus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
