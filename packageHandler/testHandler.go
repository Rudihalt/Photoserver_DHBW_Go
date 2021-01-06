package packageHandler

import (
	"html/template"
	"net/http"
)


var testTemplate *template.Template

type ViewData struct {
	Name string
}

func InitTemplate(){
	var err error
	testTemplate, err = template.ParseFiles("static/test.html")

	if err != nil{
		panic(err)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{"John Smith"}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
