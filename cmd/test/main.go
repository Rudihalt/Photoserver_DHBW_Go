package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"photoserver/packageObjects"
	"photoserver/packageTools"
	"regexp"
)

func main() {

	packageObjects.SavePhoto("p1.jpg", "yannis", "static/images/p1.jpg", "ABCDEF", "2020:10:29 12:45:23")
	packageObjects.SavePhoto("p1.jpg", "yannis", "static/images/p1.jpg", "ABCDEFG", "2020:10:29 12:45:23")
	packageObjects.SavePhoto("p1.jpg", "yannis", "static/images/p1.jpg", "ABCDEFGH", "2020:10:29 12:45:23")

	packageObjects.AddComment("yannis", "hash1", "Das ist ein Test")
	packageObjects.AddComment("yannis", "hash2", "Das ist ein Test")



	/*fmt.Println(ReadExifFromFile("static/images/p3.jpg"))

	userPtr := packageObjects.CreateUser("x2", "123456")
	if userPtr == nil {
		log.Println("User already exist!")
	}
	packageObjects.SavePhoto("photo.jpg", userPtr.Username, "ABCDEFG", "2020:10:29 13:34:25")*/



	// path, _ := os.Getwd()
	// path += "/static/images/p1.jpg"

	// SendFileUploadRequest("https://localhost:4443/api", path)

	// http.Handle("/", http.FileServer(http.Dir("./static/images")))

	//fs := http.FileServer(http.Dir("./static/images"))
	//http.Handle("/images/", http.StripPrefix("/images", fs))

	//log.Fatal(http.ListenAndServe(":8080", nil))


}

// https://gist.github.com/mattetti/5914158
func SendFileUploadRequest(uri string, path string) {
	// get date from exif header image = path
	req, err := createFileUploadRequest(uri, path, "date")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		response.Body.Close()
	}
}

func createFileUploadRequest(uri string, path string, date string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	_ = writer.WriteField("datetime", date)
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	log.Println(req)
	return req, err
}

func testing() {
	fmt.Println(ReadExifFromFile("static/images/p2.jpg"))
	fmt.Println(GetDateObjectFromString(ReadExifFromFile("static/images/p2.jpg")).Minute)

	packageTools.InitCache(2)
	cache := packageTools.GetGlobalCache()

	cache.InitLru(2)
	cache.Put(2, "a")
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))
	cache.Put(1, "b")
	cache.Put(1, "c")
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	cache.Put(8, "d")
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(8))

	userPtr := packageObjects.CreateUser("testuser", "123456")
	if userPtr == nil {
		log.Println("User already exist!")
	}
	//packageObjects.SavePhoto("photo.jpg", userPtr.Username, "ABCDEF", "2020:10:29 13:34:25")
}

type Date struct {
	Format string
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
}

func GetDateObjectFromString(input string) Date {
	date := Date{
		Format: input,
		Year:   input[0:4],
		Month:  input[5:7],
		Day:    input[8:10],
		Hour:   input[11:13],
		Minute: input[14:16],
		Second: input[17:19],
	}
	return date
}

func ReadExifFromFile(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var checkString = string(f) //[0:1000]

	re := regexp.MustCompile(`[0-9]{4}:[0-9]{2}:[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}`)

	return re.FindString(checkString)
}




