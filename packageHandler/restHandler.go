package packageHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func RESTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if r.Method == "POST" {
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println("Error Retrieving the File")
			log.Println(err)
		}
		// read datetime data
		datetime := r.FormValue("datetime")
		log.Println("datetime", datetime)
		defer file.Close()

		// posted stuff (hot glue)
		log.Printf("Uploaded File: %+v\n", handler.Filename)
		log.Printf("File Size: %+v\n", handler.Size)
		log.Printf("MIME Header: %+v\n", handler.Header)

		path, _ := os.Getwd()
		path += "/static/images"
		f, err := os.Create(filepath.Join(path, handler.Filename))
		if err != nil {
			log.Println(f, "was successfully created")
		}
		defer f.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}
		f.Write(fileBytes)
	}
}
