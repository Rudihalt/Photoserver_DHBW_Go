package packageObjects

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"photoserver/packageTools"
)

type Photo struct {
	Name string `json:name`
	userID int `json:userID`
	Encoded string `json:encoded`
	exifDate string `json:exifDate`
}

func addNewPhotoToDB(name string, userid int, encoded string, exifdate string) *Photo {
	photo := Photo{
		Name: name,
		userID: userid,
		Encoded: encoded,
		exifDate: exifdate,
	}




	return &photo
}

func GetPhotoByID(id int) {
	lruCache := packageTools.GetGlobalCache()
	var cache = *lruCache

	encoded := cache.Get(id)

	if encoded == "" {
		log.Println(encoded)

		encoded = getPhotoByIDDB(id)
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
		// Handle error
	}

	log.Println("fmtName: " + fmtName)

	return &img
}