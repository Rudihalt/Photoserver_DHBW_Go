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
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		user := packageObjects.GetUserByToken(cookie.Value)
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)

		GData := GalleryData{Username: user.Username,
			Amount: 0,
		}

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

			var paging []int
			for i := 0; i < pages; i++ {
				paging = append(paging, i+1)
			}
			GData.Photos = packageObjects.GetPhotosForPage(user.Username, p)
			GData.Amount = len(*GData.Photos)
			GData.Pages = paging
		}

		err = GalleryTemplate.Execute(w, GData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
