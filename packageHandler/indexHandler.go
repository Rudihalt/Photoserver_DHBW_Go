package packageHandler

import (
	"html/template"
	"log"
	"net/http"
)

var NavData NavBarData

var NavTemplate *template.Template
var LoginTemplate *template.Template
var RegisterTemplate *template.Template
var UploadTemplate *template.Template

type NavBarData struct {
	Username string
}

func InitTemplates() {
	var err error
	NavTemplate, err = template.ParseFiles("static/template/nav.html")
	LoginTemplate, err = template.ParseFiles("static/template/login.html")
	RegisterTemplate, err = template.ParseFiles("static/template/register.html")
	UploadTemplate, err = template.ParseFiles("static/template/upload.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	if len(name) == 0 {
		name = "World"
	}
	responseString := "<html><body>Index Page<br>Hello " + name + "</body></html>"
	w.Write([]byte(responseString)) // unbedingt Templates verwenden!
}
