/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"photoserver/packageObjects"
	"photoserver/packageTools"
	"strconv"
	"strings"
)

func main() {
	// Templates: https://www.calhoun.io/intro-to-templates-p2-actions/

	var address = flag.String("address", "localhost", "Adresse des Photoserver")
	var port = flag.Int("port", 4443, "Port des Photoserver")
	var dataFolder = flag.String("data", "", "Pfad des Ordners mit den Photos")
	var username = flag.String("username", "", "Username")
	var password = flag.String("password", "", "Password")

	flag.Parse()

	log.Println("----- BATCH-UPLOAD -----")
	log.Println()
	log.Println("Folgende Parameter werden verwendet: Port: " + strconv.Itoa(*port) + " Address: " + *address + " Data-Folder: " + *dataFolder)
	log.Println()

	// check if host-address is reachable
	var host = *address
	if strings.Contains(host, "https://") || strings.Contains(*address, "http://") {
		host = strings.Replace(host, "https://", "", 1)
		host = strings.Replace(host, "http://", "", 1)
	}
	host += ":" + strconv.Itoa(*port)
	if !packageTools.CheckHost(host) {
		log.Fatalln("Could not connect to", host)
	}

	// check if directory exist
	if !packageTools.PathExist(*dataFolder) {
		log.Fatalln("Data-Folder not found")
	}

	// check if user exist
	if !packageObjects.UserExists(*username) {
		log.Fatalln("User does not exist")
	}

	//check if password is correct
	ok, _ := packageObjects.CheckPassword(*username, *password)
	if !ok {
		log.Fatalln("Password for username", *username, "is not correct")
	}

	host = "https://" + host + "/api"
	var files []string

	// send for each file a post request to the endpoint
	path := *dataFolder
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		files = append(files, p)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files[1:] {
		log.Println(file)
		packageTools.SendFileUploadRequest(host, file, *username)
	}
}
