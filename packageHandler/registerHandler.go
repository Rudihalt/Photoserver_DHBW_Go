package packageHandler

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	NavBarData.CurrentPage = "register"

	err := NavTemplate.Execute(w, NavBarData)
	err = RegisterTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
