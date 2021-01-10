/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"archive/zip"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"testing"
)

func TestCreateZipFile(t *testing.T) {
	// initialize test parameter
	dir, _ := os.Getwd()
	SetPublicDir(path.Join(dir, "static"))
	files := []string{"test1.json", "test2.json"}
	username := "test"

	var zipItems []ZipItem
	for _, file := range files {
		p := path.Join(GetPublicDir(), file)
		f, _ := os.Create(p)
		item := ZipItem{Name: file, Path: "/" + file, Format: "1x2", Amount: 3}
		zipItems = append(zipItems, item)
		f.Close()
	}

	zipFile, err := CreateZipFile(zipItems, username)
	zipLocation := path.Join(GetOrderFolder(), username+".zip")
	assert.Equal(t, zipLocation, zipFile)
	assert.Nil(t, err)
	// delete all created files
	err = os.RemoveAll(GetPublicDir())
	if err != nil {
		log.Println(err)
	}
}

func TestAddFile(t *testing.T) {
	// initialize test parameter
	files := []string{"test1.json", "test2.json", "test3.json"}
	username := "test"
	zipFile, err := os.Create(GetOrderFolder() + username + ".zip")
	if err != nil {
		log.Println(err)

	}
	defer zipFile.Close()
	// add two files to zip
	writer := zip.NewWriter(zipFile)
	for _, file := range files {
		p := path.Join(GetPublicDir(), file)
		f, _ := os.Create(p)
		addFile(writer, "./static/"+file, file)
		f.Close()
	}
	writer.Close()
	zf, err := zip.OpenReader(zipFile.Name())
	err = zf.Close()
	if err != nil {
		log.Println(err)
	}
	amount := len(zf.File)
	assert.Equal(t, 3, amount)
	zipFile.Close()
	// delete all created files
	err = os.RemoveAll(GetPublicDir())
	if err != nil {
		log.Println(err)
	}
}
