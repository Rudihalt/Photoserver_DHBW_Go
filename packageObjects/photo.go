package packageObjects

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
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
	Name     string `json:"name"`
	Path     string `json:"name"`
	Hash     string `json:"hash"`
	Date     string `json:"date"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Comment string `json:"comment"`
	Date    string `json:"date"`
}

func GetAllPhotosByUser(username string) *[]Photo {
	var photos []Photo
	photosFile, err := ioutil.ReadFile("static/data/photos_" + username + ".json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(photosFile, &photos)

	if err != nil {
		panic(err)
	}

	return &photos
}

func getPhotosForPage(username string, page int) *[]Photo{
	photos := *GetAllPhotosByUser(username)
	total := len(photos)

	if total == 0 {
		return nil
	}

	photosPerPage := 3

	if (total / photosPerPage) + 1 > page {
		page = total / photosPerPage
	}

	start := page * photosPerPage
	end := start + photosPerPage

	if end > total {
		end = total
	}

	part := photos[start:end]

	return &part
}

func GetCommentsFromPhoto(photo *Photo) *[]Comment {
	return &photo.Comments
}

func GetPhotoByUserAndHash(username string, hash string) *Photo {
	photos := *GetAllPhotosByUser(username)

	for _, photo := range photos {
		if photo.Hash == hash {
			return &photo
		}
	}

	return nil
}

func SavePhoto(name string, username string, path string, encoded string, date string) *Photo {
	hash := packageTools.HashSHA(encoded)

	if GetPhotoByUserAndHash(username, hash) != nil {
		return nil
	}

	photo := Photo {
		Name:     name,
		Path:     path,
		Hash:     hash,
		Date:     date,
	}

	photoJson, err := json.MarshalIndent(photo, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/data/photo_"+hash+".json", photoJson, 0644)
	if err != nil {
		panic(err)
	}

	addPhotoToUser(username, hash)

	return &photo
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