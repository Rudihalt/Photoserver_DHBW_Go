package packageHandler

import (
	"net/http"
)

func GalleryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := NavTemplate.Execute(w, nil)
	err = GalleryTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}