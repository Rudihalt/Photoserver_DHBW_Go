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
		// check if the contentType is correct
		contentType := handler.Header.Get("Content-Type")
		if contentType != "application/octet-stream" {
			if contentType != "image/jpeg" {
				http.Redirect(w, r, "/upload", http.StatusSeeOther)
				log.Println("No correct jpeg format found:", contentType)
				return
			}
		}

		// get username from post
		username := r.FormValue("username")

		// set the path of the image
		path, _ := os.Getwd()
		path += "/static/images"

		// create file
		filePath := filepath.Join(path, handler.Filename)
		f, err := os.Create(filePath)
		if err != nil {
			log.Println(f, "was successfully created")
		}

		// read the file to a byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}

		// get the datetime from the exif header
		date, err := packageTools.GetDateTime(fileBytes)
		if err != nil {
			log.Println(err)
		}
		// write the file
		f.Write(fileBytes)
		e := f.Close()
		if e != nil {
			fmt.Println("Error Closing file")
		} else {
			fmt.Println("Successfully Closing file")
		}

		// create hash for file
		shaFile := packageTools.HashSHAFile(filePath)
		fmt.Println("shaFile: " + shaFile)

		// rename the file
		e = os.Rename(filePath, filepath.Join(path, shaFile+".jpg"))
		if e != nil {
			fmt.Println(e)
		}

		// ouput to save
		log.Println("username:", username)
		log.Println("Uploaded File:", handler.Filename)
		log.Println("date:", date)

		// save the photo to the user
		photo := packageObjects.SavePhoto(handler.Filename, username, "/images/"+shaFile+".jpg", date)
		if photo == nil {
			log.Println("File could not be uploaded! File already exists?")
		} else {
			log.Println("File successfully uploaded!")
		}

		// redirect to the gallery view
		http.Redirect(w, r, "/gallery", http.StatusSeeOther)
	}
}
