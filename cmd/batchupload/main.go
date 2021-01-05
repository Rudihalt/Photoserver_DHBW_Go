package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Templates: https://www.calhoun.io/intro-to-templates-p2-actions/

	var address = flag.String("address", "localhost", "Server-Address")
	var port = flag.Int("port", 4443, "Port")
	var dataFolder = flag.String("data", "", "Data-Folder")
	
	flag.Parse()

	log.Println("----- BATCH-UPLOAD -----")
	log.Println()
	log.Println("Folgende Parameter werden verwendet: Port: " + strconv.Itoa(*port) + " Address: " + *address + " Data-Folder: " + *dataFolder)
	log.Println()

	if *dataFolder == "" {
		log.Println("Folgende Eingabeparameter übergeben:")
		log.Println("-address [Adresse] -> Adresse des Photoserver")
		log.Println("-port [Port] -> Port des Photoserver.")
		log.Println("-data [Adresse] -> Adresse ")
		log.Println()
		log.Println("-> Der Server ist nur über https erreichbar. Sonst ist ein Batch-Upload nicht möglich")

		os.Exit(0)
	}

	// https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
	url := "https://" + *address + "/api:" + strconv.Itoa(*port)
	fmt.Println("URL:>", url)

	var jsonStr = []byte(
		`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body:", string(body))

}
