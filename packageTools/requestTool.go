/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// https://gist.github.com/mattetti/5914158
func SendFileUploadRequest(uri string, path string, username string) {
	req, err := createFileUploadRequest(uri, path, username)
	if err == nil {
		// ignore bad certificate at transport
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		// create client for upload
		client := &http.Client{Transport: tr}
		// do request
		response, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		} else {
			// get response
			body := &bytes.Buffer{}
			_, err := body.ReadFrom(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			response.Body.Close()
		}
	}
}

// function which creates a Post Request for the backend
func createFileUploadRequest(uri string, path string, username string) (*http.Request, error) {
	lowerPath := strings.ToLower(path)
	// check if path ends with jpg or jpeg
	if strings.HasSuffix(lowerPath, ".jpg") || strings.HasSuffix(lowerPath, ".jpeg") {
		// open file
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// get byte array
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)

		// write username value
		_ = writer.WriteField("username", username)
		err = writer.Close()
		if err != nil {
			return nil, err
		}

		// create new PostRequest
		req, err := http.NewRequest("POST", uri, body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		return req, nil
	}
	// return error if no jpg or jpeg detected wurde
	return nil, errors.New("No JPEG detected")
}

// function to check if host is reachable
func CheckHost(host string) bool {
	_, err := net.DialTimeout("tcp", host, time.Duration(5)*time.Second)
	if err != nil {
		return false
	}
	return true
}
