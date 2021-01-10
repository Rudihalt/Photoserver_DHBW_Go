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
	"strconv"
)

type GalleryData struct {
	Username string
	Photos   *[]packageObjects.Photo
	Amount   int
	Pages    []int
}

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
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

		GData := GalleryData{Username: user.Username,
			Amount: 0,
		}

		// page logic:
		// get pages from get parameter and compare it with the max page amount
		pages := packageObjects.GetPhotoPageAmount(user.Username)
		if pages != 0 {
			page := r.URL.Query().Get("p")
			p, _ := strconv.Atoi(page)
			if p == 0 {
				p = 1
			}
			if p > pages {
				p = pages
			}

			// create paging array for pagination
			var paging []int
			for i := 0; i < pages; i++ {
				paging = append(paging, i+1)
			}
			// set the photos max value from GetPhotosForPage is 9
			GData.Photos = packageObjects.GetPhotosForPage(user.Username, p)
			GData.Amount = len(*GData.Photos)
			GData.Pages = paging
		}

		// hand over the GData to the template
		err = GalleryTemplate.Execute(w, GData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
