package packageHandler

import (
	"html/template"
	"log"
	"net/http"
	"photoserver/packageObjects"
)

var NavData NavBarData

var IndexTemplate *template.Template
var NavTemplate *template.Template
var LoginTemplate *template.Template
var RegisterTemplate *template.Template
var UploadTemplate *template.Template
var GalleryTemplate *template.Template
var DiashowTemplate *template.Template
var OrderTemplate *template.Template
var ImageTemplate *template.Template
var AlbumTemplate *template.Template

type NavBarData struct {
	Username string
}

type IndexViewData struct {
	Greeting string
}

func InitTemplates() {
	var err error
	IndexTemplate, err = template.ParseFiles("static/template/index.html")
	NavTemplate, err = template.ParseFiles("static/template/nav.html")
	LoginTemplate, err = template.ParseFiles("static/template/login.html")
	RegisterTemplate, err = template.ParseFiles("static/template/register.html")
	UploadTemplate, err = template.ParseFiles("static/template/upload.html")
	GalleryTemplate, err = template.ParseFiles("static/template/paging.html")
	DiashowTemplate, err = template.ParseFiles("static/template/diashow.html")
	OrderTemplate, err = template.ParseFiles("static/template/order.html")
	ImageTemplate, err = template.ParseFiles("static/template/image_com.html")
	AlbumTemplate, err = template.ParseFiles("static/template/album.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var cookie, _ = r.Cookie("csrftoken")

	var user *packageObjects.User
	var NavData NavBarData
	var greeting = "Herzlich Willkommen auf auf dem Photo-Server. Registriere dich oder logge dich ein!"

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
