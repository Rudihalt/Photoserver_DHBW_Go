/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"html/template"
	"log"
	"net/http"
	"photoserver/packageObjects"
	"photoserver/packageTools"
)

var NavData NavBarData

var IndexTemplate *template.Template
var NavTemplate *template.Template
var LoginTemplate *template.Template
var RegisterTemplate *template.Template
var UploadTemplate *template.Template
var GalleryTemplate *template.Template
var DiaShowTemplate *template.Template
var OrderTemplate *template.Template
var ImageTemplate *template.Template

type NavBarData struct {
	Username string
}

type IndexViewData struct {
	Greeting string
}

func InitTemplates() {
	var err error
	IndexTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/index.html")
	NavTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/nav.html")
	LoginTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/login.html")
	RegisterTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/register.html")
	UploadTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/upload.html")
	GalleryTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/gallery.html")
	DiaShowTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/diashow.html")
	OrderTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/order.html")
	ImageTemplate, err = template.ParseFiles(packageTools.GetWD() + "/static/template/image_com.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var cookie, _ = r.Cookie("csrftoken")

	var user *packageObjects.User
	var NavData NavBarData
	var greeting = "Herzlich Willkommen auf auf dem Photo-Server. Registriere dich oder logge dich ein! Testuser: user= hneemann password= 123456"

	if cookie != nil {
		user = packageObjects.GetUserByToken(cookie.Value)
		greeting = "Herzlich Willkommen auf dem Photo-Server, " + user.Username
		NavData = NavBarData{Username: user.Username}
	}

	IndexViewData := IndexViewData{
		Greeting: greeting,
	}

	err := NavTemplate.Execute(w, NavData)
	err = IndexTemplate.Execute(w, IndexViewData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
