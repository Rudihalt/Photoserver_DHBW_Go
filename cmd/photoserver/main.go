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
	"net/http"
	"photoserver/packageHandler"
	"photoserver/packageTools"
	"strconv"
)

func main() {
	// Templates: https://www.calhoun.io/intro-to-templates-p2-actions/

	var port = flag.Int("port", 4443, "Port")
	var certificates = flag.String("certificates", "static/ssl", "SSL-Certificates")
	var data = flag.String("data", "static/data", "Data-Directory")

	flag.Parse()

	log.Println("----- Photoserver -----")
	log.Println()
	log.Printf("Folgende Parameter werden verwendet: Port: " + strconv.Itoa(*port) + " Certificates: " + *certificates + " Data: " + *data)
	log.Println()

	// check if all necessary directories exist
	packageTools.CreateNecessaryDirs()

	// initialize random and all html templates
	packageTools.Init()
	packageHandler.InitTemplates()

	http.HandleFunc("/", packageHandler.IndexHandler)
	http.HandleFunc("/login", packageHandler.LoginHandler)
	http.HandleFunc("/register", packageHandler.RegisterHandler)
	http.HandleFunc("/logout", packageHandler.LogoutHandler)
	http.HandleFunc("/diashow", packageHandler.DiashowHandler)
	http.HandleFunc("/upload", packageHandler.UploadHandler)
	http.HandleFunc("/gallery", packageHandler.GalleryHandler)
	http.HandleFunc("/order", packageHandler.OrderHandler)
	http.HandleFunc("/image", packageHandler.ImageHandler)

	http.HandleFunc("/api", packageHandler.RESTHandler)

	fs := http.FileServer(http.Dir("./static/images"))
	http.Handle("/images/", http.StripPrefix("/images", fs))

	// https://stackoverflow.com/questions/25552107/how-to-serve-static-files-with-basic-authentication
	// Wer Lust hat, kann den FileServer vor fremden Zugriffen sichern. Heute wird das nichts mehr.

	// log.Fatalln(http.ListenAndServe(":8080", nil))
	// https://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl
	// Command: openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
	// https://serverfault.com/questions/366372/is-it-possible-to-generate-rsa-key-without-pass-phrase
	// No Passphrase for testing project. Use -nodes for no DES encryption for private key
	// Result Command: openssl req -nodes -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
	log.Fatalln(http.ListenAndServeTLS(":"+strconv.Itoa(*port), *certificates+"/cert.pem", *certificates+"/key.pem", nil))

}

func handleFileServer(dir, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(dir))
	realHandler := http.StripPrefix(prefix, fs).ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL)
		realHandler(w, req)
	}
}
