/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageObjects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"photoserver/packageTools"
)

type Photo struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Hash string `json:"hash"`
	Date string `json:"date"`
}

func GetAllPhotosByUser(username string) *[]Photo {
	var photos []Photo
	var photosFile []byte

	photosFile, err := ioutil.ReadFile(packageTools.GetWD() + "/static/data/photos_" + username + ".json")

	if err != nil {
		fmt.Println("Neue Datei anlegen: photos_" + username + ".json")
	}

	err = json.Unmarshal(photosFile, &photos)

	if err != nil {
		// panic(err)
	}

	return &photos
}

func GetPhotoPageAmount(username string) int {
	photos := *GetAllPhotosByUser(username)
	total := len(photos)
	if total == 0 {
		return 0
	}
	amount := total / 9
	if total == 9 {
		return 1
	}
	return amount + 1
}

func GetPhotosForPage(username string, page int) *[]Photo {
	page--
	photos := *GetAllPhotosByUser(username)
	total := len(photos)

	if total == 0 {
		return nil
	}

	photosPerPage := 9

	if (total/photosPerPage)+1 < page {
		page = total / photosPerPage
	}

	start := page * photosPerPage
	if page == 0 {
		start = 0
	}
	end := start + photosPerPage

	if end > total {
		end = total
	}

	part := photos[start:end]

	return &part
}

func GetPhotoByUserAndHash(photos *[]Photo, hash string) *Photo {

	for _, photo := range *photos {
		if photo.Hash == hash {
			return &photo
		}
	}

	return nil
}

func SavePhoto(name string, username string, path string, date string) *Photo {
	dir := packageTools.GetWD()
	hash := packageTools.HashSHAFile(dir + "/static" + path)

	fmt.Println("Hashing " + name + " Path: " + path + " Hash: " + hash)

	currentPhotos := *GetAllPhotosByUser(username)

	if GetPhotoByUserAndHash(&currentPhotos, hash) != nil {
		return nil
	}

	photo := Photo{
		Name: name,
		Path: path,
		Hash: hash,
		Date: date,
	}

	currentPhotos = append(currentPhotos, photo)

	savePhotos(username, &currentPhotos)

	return &photo
}

func savePhotos(username string, photos *[]Photo) {
	photoJson, err := json.MarshalIndent(photos, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(packageTools.GetWD() + "/static/data/photos_"+username+".json", photoJson, 0644)
	if err != nil {
		panic(err)
	}
}

