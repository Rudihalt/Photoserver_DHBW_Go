/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"os"
	"strings"
)

var dataFolder string

// checks if path or file exist
func PathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// creates the directory if it does not exist
func CreateDirIfNotExists(path string) bool {
	if !PathExist(path) {
		os.MkdirAll(path, os.ModeDir)
		return true
	}
	return false
}

// creates at start of the program all necessary directories
func CreateNecessaryDirs() {
	_ = GetImageFolder()
	_ = GetOrderFolder()
	_ = GetDataFolder()
}

// sets the public/home directory
func SetPublicDir(data string) {
	dataFolder = data
}

// function to get the public directory
func GetPublicDir() string {
	retStr := dataFolder
	CreateDirIfNotExists(retStr)
	return retStr
}

// get the image directory
func GetImageFolder() string {
	retStr := dataFolder + "/images/"
	CreateDirIfNotExists(retStr)
	return retStr
}

// get the order directory
func GetOrderFolder() string {
	retStr := dataFolder + "/orders/"
	CreateDirIfNotExists(retStr)
	return retStr
}

// get the data directory
func GetDataFolder() string {
	retStr := dataFolder + "/data/"
	CreateDirIfNotExists(retStr)
	return retStr
}

func GetWD() string {
	// https://stackoverflow.com/questions/14249217/how-do-i-know-im-running-within-go-test

	// Important for tests! if Path contains package (packageTools, ppackageObjects, packageHandler),
	// then the test files are running.

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if strings.Contains(wd, "package") {
		wd = wd + "/.."
	}

	return wd
}
