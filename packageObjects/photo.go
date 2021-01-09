package packageObjects

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	photosFile, err := ioutil.ReadFile("static/data/photos_" + username + ".json")

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

	start := page*photosPerPage - 1
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
	dir, _ := os.Getwd()
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

	err = ioutil.WriteFile("static/data/photos_"+username+".json", photoJson, 0644)
	if err != nil {
		panic(err)
	}
}

func GetPhotoByUserAndHash2(username string, hash string) {
	lruCache := packageTools.GetGlobalCache()
	var cache = *lruCache

	encoded := cache.Get(2)

	if encoded == "" {
		log.Println(encoded)

		encoded = getPhotoByIDDB(2)
	}

	// todo: load other info from json file an

	log.Println(encoded)
}

func getPhotoByIDDB(id int) string {
	return "not implemented"
}

// https://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang

var ImageTemplate string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body><h1>Image-Test</h1><img src="data:image/jpg;base64,{{.Image}}"></body>`

// Writeimagewithtemplate encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func WriteImageWithTemplate(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

func GetImageByName(fileName string) *image.Image {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening the file: " + fileName)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	log.Println("fmtName: " + fmtName)

	return &img
}
