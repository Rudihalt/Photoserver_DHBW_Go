/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang

func CreateZipFile(files []string, username string) {
	zipFile, err := os.Create("./static/orders/" + username + ".zip")
	if err != nil {
		log.Println(err)
	}
	defer zipFile.Close()

	writer := zip.NewWriter(zipFile)
	for _, file := range files {
		addFile(writer, file)
	}

	if err != nil {
		log.Println(err)
	}

	// Make sure to check the error on Close.
	err = writer.Close()
	if err != nil {
		log.Println(err)
	}
}

func addFile(w *zip.Writer, path string) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	// Add some files to the archive.
	file := strings.Split(path, "\\")
	file = strings.Split(file[len(file)-1], "/")
	f, err := w.Create(file[len(file)-1])
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(dat)
	if err != nil {
		log.Println(err)
	}
}

func ZipWriter() {
	baseFolder := "/Users/tom/Desktop/testing/"

	// Get a Buffer to Write To
	outFile, err := os.Create(`/Users/tom/Desktop/zip.zip`)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, baseFolder, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
