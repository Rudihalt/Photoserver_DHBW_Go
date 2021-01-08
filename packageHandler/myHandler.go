package packageHandler

import (
	"html/template"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	//var img = packageObjects.GetImageByName("static/images/p1.jpg")
	//packageObjects.WriteImageWithTemplate(w, img)


	// responseString := "<html><body>My</body></html>"
	// w.Write([]byte(responseString))

	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{"John Smith"}
	err := testTemplate.Execute(w, vd)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


var myTemplate *template.Template

func InitTemplateMy(){
	var err error
	testTemplate, err = template.ParseFiles("static/my.html")

	if err != nil{
		panic(err)
	}
}
