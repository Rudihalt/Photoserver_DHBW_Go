package packageHandler

import (
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	NavBarData.CurrentPage = "upload"

	err := NavTemplate.Execute(w, NavBarData)
	err = UploadTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
