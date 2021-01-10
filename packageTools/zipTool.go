/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type ZipItem struct {
	Name   string
	Path   string
	Format string
	Amount int
}

// https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang

// function to create a zip archive from the zipitem which are
// data from the users order
func CreateZipFile(files []ZipItem, username string) (string, error) {
	// create zip archive
	zipFile, err := os.Create(GetOrderFolder() + username + ".zip")
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer zipFile.Close()

	// add all items to zip to the zip
	writer := zip.NewWriter(zipFile)
	for _, file := range files {
		for i := 0; i < file.Amount; i++ {
			name := file.Name
			name = strings.Replace(name, ".", "-"+file.Format+".", 1)
			if i != 0 {
				name = strings.Replace(name, ".", "-"+strconv.Itoa(i)+".", 1)
			}
			addFile(writer, "./static"+file.Path, name)
		}
	}

	// close the writer
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return zipFile.Name(), nil
}

// function to add a file to the current created zip archive
func addFile(w *zip.Writer, path string, name string) {
	// reads the data from
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	// write files to the zip archive
	f, err := w.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(dat)
	if err != nil {
		log.Println(err)
	}
}
