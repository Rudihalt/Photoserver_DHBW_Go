package packageHandler

import "net/http"

type testdreck struct {
	Path string
}

func DiashowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

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

	err := NavTemplate.Execute(w, nil)
	err = DiashowTemplate.Execute(w, anus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
