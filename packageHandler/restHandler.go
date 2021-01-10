/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"photoserver/packageObjects"
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
		defer file.Close()

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


		path, _ := os.Getwd()
		path += "/static/images"

		//TODO: CHECK IF PHOTO ALREADY EXIST

		filePath := filepath.Join(path, handler.Filename)
		f, err := os.Create(filePath)
		if err != nil {
			log.Println(f, "was successfully created")
		}

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}

		date, err := packageTools.GetDateTime(fileBytes)
		if err != nil {
			log.Println(err)
		}
		f.Write(fileBytes)
		e := f.Close()
		if e != nil {
			fmt.Println("Error Closing file")
		} else {
			fmt.Println("Successfully Closing file")
		}

		shaFile := packageTools.HashSHAFile(filePath)
		fmt.Println("shaFile: " + shaFile)

		e = os.Rename(filePath, filepath.Join(path, shaFile + ".jpg"))
		if e != nil {
			fmt.Println(e)
		}

		// ouput to save
		log.Println("username:", username)
		log.Println("Uploaded File:", handler.Filename)
		log.Println("date:", date)

		photo := packageObjects.SavePhoto(handler.Filename, username, "/images/"+shaFile + ".jpg", date)
		if photo == nil {
			log.Println("File could not be uploaded! File already exists?")
		} else {
			log.Println("File successfully uploaded!")
		}

		http.Redirect(w, r, "/gallery", http.StatusSeeOther)
	}
}
