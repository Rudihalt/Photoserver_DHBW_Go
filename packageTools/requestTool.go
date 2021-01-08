package packageTools

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// https://gist.github.com/mattetti/5914158
func SendFileUploadRequest(uri string, path string) {
	// get date from exif header image = path
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	date, err := GetDateTime(b)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := createFileUploadRequest(uri, path, date)
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

// function to check if host is reachable
func CheckHost(host string) bool {
	_, err := net.DialTimeout("tcp", host, time.Duration(5)*time.Second)
	if err != nil {
		return false
	}
	return true
}
