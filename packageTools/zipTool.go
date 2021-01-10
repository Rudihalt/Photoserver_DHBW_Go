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

func CreateZipFile(files []ZipItem, username string) error {
	zipFile, err := os.Create("./static/orders/" + username + ".zip")
	if err != nil {
		log.Println(err)
		return err
	}
	defer zipFile.Close()

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

	// Make sure to check the error on Close.
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func addFile(w *zip.Writer, path string, name string) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	// Add some files to the archive.
	f, err := w.Create(name)
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(dat)
	if err != nil {
		log.Println(err)
	}
}
