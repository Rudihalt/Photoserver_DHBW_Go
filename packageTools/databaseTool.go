package packageTools

import "os"

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
		return true;
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
	retStr :=  dataFolder + "/albums/"
	return retStr
}

func getOrderFolder() string {
	retStr :=  dataFolder + "/order/"
	return retStr
}

func getUserFile() string {
	retStr :=  dataFolder + "/images/user.json"
	return retStr
}