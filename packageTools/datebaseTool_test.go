/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"testing"
)

func TestPathExist(t *testing.T) {
	// check for not existing path
	assert.False(t, PathExist("Somewhere/in/your/directory/I/hope/this/path/doesnt/exist"), "Path does not exist")

	// get current workdirectory for a exist path
	dir, _ := os.Getwd()
	assert.True(t, PathExist(dir), "Path exists")
}

func TestCreateDirIfNotExists(t *testing.T) {
	// initialize a not existing directory
	dir, _ := os.Getwd()
	dir = path.Join(dir, "test_directory")
	CreateDirIfNotExists(dir)
	// check if direcotry exist
	exist := PathExist(dir)
	assert.True(t, exist)
	// remove directory
	err := os.Remove(dir)
	if err != nil {
		log.Println("Could not delete directory")
	}
}

func TestGetPublicDir(t *testing.T) {
	wd, _ := os.Getwd()
	expected := path.Join(wd, "s787IK")
	SetPublicDir(expected)
	actual := GetPublicDir()

	assert.Equal(t, expected, actual)

	err := os.RemoveAll(expected)
	if err != nil {
		log.Println("Could not delete directory")
	}
}

func TestGetImageFolder(t *testing.T) {
	wd, _ := os.Getwd()
	public := path.Join(wd, "s787IK")
	SetPublicDir(public)
	expected := path.Join(public, "images")
	actual := GetImageFolder()

	expected += "/"
	assert.Equal(t, expected, actual)

	err := os.RemoveAll(public)
	if err != nil {
		log.Println("Could not delete directory")
	}
}

func TestGetOrderFolder(t *testing.T) {
	wd, _ := os.Getwd()
	public := path.Join(wd, "s787IK")
	SetPublicDir(public)
	expected := path.Join(public, "orders")
	actual := GetOrderFolder()

	expected += "/"
	assert.Equal(t, expected, actual)

	err := os.RemoveAll(public)
	if err != nil {
		log.Println("Could not delete directory")
	}
}

func TestGetDataFolder(t *testing.T) {
	wd, _ := os.Getwd()
	public := path.Join(wd, "s787IK")
	SetPublicDir(public)
	expected := path.Join(public, "data")
	actual := GetDataFolder()

	expected += "/"
	assert.Equal(t, expected, actual)

	err := os.RemoveAll(public)
	if err != nil {
		log.Println("Could not delete directory")
	}
}
