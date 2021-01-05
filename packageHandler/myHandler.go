package packageHandler

import (
	"net/http"
	"photoserver/packageObjects"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	var img = packageObjects.GetImageByName("static/images/p1.jpg")
	packageObjects.WriteImageWithTemplate(w, img)
	// responseString := "<html><body>My</body></html>"
	// w.Write([]byte(responseString))
}
