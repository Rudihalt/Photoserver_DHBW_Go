package packageHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"photoserver/packageTools"
)

func RESTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.Method == "POST" {
		r.ParseMultipartForm(10 << 20)
		// posted stuff
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println("Error Retrieving the File")
			log.Println(err)
		}
		contentType := handler.Header.Get("Content-Type")
		if contentType != "application/octet-stream" {
			if contentType != "image/jpeg" {
				http.Redirect(w, r, "/upload", http.StatusSeeOther)
				log.Println("No correct jpeg format found:", contentType)
				return
			}
		}

		// read datetime data
		username := r.FormValue("username")
		defer file.Close()

		path, _ := os.Getwd()
		path += "/static/images"

		//TODO: CHECK IF PHOTO ALREADY EXIST

		f, err := os.Create(filepath.Join(path, handler.Filename))
		if err != nil {
			log.Println(f, "was successfully created")
		}
		defer f.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}

		date, err := packageTools.GetDateTime(fileBytes)
		if err != nil {
			log.Println(err)
		}
		f.Write(fileBytes)

		// ouput to save
		log.Println("username:", username)
		log.Println("Uploaded File:", handler.Filename)
		log.Println("date:", date)
		http.Redirect(w, r, "/my", http.StatusSeeOther)
	}
}
