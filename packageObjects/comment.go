/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
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
	"time"
)

type Comment struct {
	Comment string `json:"comment"`
	Date    string `json:"date"`
	Hash    string `json:"hash"`
}

func GetAllCommentsByUser(username string) *[]Comment {
	var comments []Comment
	var commentsFile []byte

	commentsFile, err := ioutil.ReadFile("static/data/comments_" + username + ".json")

	if err != nil {
		fmt.Println("Neue Datei anlegen: comments_" + username + ".json")
	}

	err = json.Unmarshal(commentsFile, &comments)

	if err != nil {
		// panic(err)
	}

	return &comments
}

func saveComments(username string, comments *[]Comment) {
	commentJson, err := json.MarshalIndent(comments, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/data/comments_"+username+".json", commentJson, 0644)
	if err != nil {
		panic(err)
	}
}

func AddComment(username string, hash string, commentStr string) *Comment {
	currentComments := *GetAllCommentsByUser(username)

	currentTime := time.Now()
	timeFormatted := currentTime.Format("2006.01.02 15:04:05")

	comment := Comment{
		Comment: commentStr,
		Hash:    hash,
		Date:    timeFormatted,
	}

	currentComments = append(currentComments, comment)

	saveComments(username, &currentComments)

	return &comment
}

func FilterAllCommentsByHash(comments *[]Comment, hash string) *[]Comment {
	var hashComments []Comment
	for _, comment := range *comments {
		if comment.Hash == hash {
			hashComments = append(hashComments, comment)
		}
	}

	if len(hashComments) == 0 {
		return nil
	}

	return &hashComments
}

func GetPhotoByIDX(id int) {
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

// https://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang

var ImageTemplateX string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body><h1>Image-Test</h1><img src="data:image/jpg;base64,{{.Image}}"></body>`

// Writeimagewithtemplate encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func WriteImageWithTemplateX(w http.ResponseWriter, img *image.Image) {

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

func GetImageByNameX(fileName string) *image.Image {
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
