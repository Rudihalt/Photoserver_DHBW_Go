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

func CreateDirIfNotExists(path string) bool {
	if PathExist(path) {
		os.Mkdir(path, os.ModeDir)
		return true
	}
	return false
}

func SetDataFolder(data string) {
	dataFolder = data
}

func getImageFolder() string {
	retStr := dataFolder + "/images/"
	CreateDirIfNotExists(retStr)
	return retStr
}

func getAlbumsFolder() string {
	retStr := dataFolder + "/albums/"
	return retStr
}

func getOrderFolder() string {
	retStr := dataFolder + "/order/"
	return retStr
}

func getUserFile() string {
	retStr := dataFolder + "/images/user.json"
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
