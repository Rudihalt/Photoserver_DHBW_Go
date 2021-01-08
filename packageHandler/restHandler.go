package packageHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func RESTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
	}
	defer file.Close()

	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)

	path, _ := os.Getwd()
	path += "/static/images"
	f, err := os.Create(filepath.Join(path, strings.Replace(handler.Filename, "p1", "p4", 1)))
	if err != nil {
		log.Println(f, "was successfully created")
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	f.Write(fileBytes)

	responseString := "<html><body>Test</body></html>"
	w.Write([]byte(responseString))
}
