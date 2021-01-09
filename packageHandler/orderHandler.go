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

type OrderViewData struct {
	OrderElements []packageObjects.OrderElement
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		user := packageObjects.GetUserByToken(cookie.Value)
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)

		orderElements := packageObjects.GetAllOrderElementsByUser(user.Username)

		orderViewData := OrderViewData{
			OrderElements: *orderElements,
		}

		err = OrderTemplate.Execute(w, orderViewData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
