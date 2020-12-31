package main

import (
	"flag"
	"log"
	"net/http"
	"photoserver/packageHandler"
	"strconv"
)

func main() {
	port := *flag.Int("port", 8080, "Port")
	certificates := *flag.String("certificates", "static/ssl", "SSL-Certificates")
	flag.Parse()

	log.Printf("Port: " + strconv.Itoa(port) + " Certificates: " + certificates)

	http.HandleFunc("/", packageHandler.IndexHandler)
	http.HandleFunc("/login", packageHandler.LoginHandler)
	http.HandleFunc("/register", packageHandler.RegisterHandler)
	http.HandleFunc("/my", packageHandler.MyHandler)
	http.HandleFunc("/diashow", packageHandler.DiashowHandler)
	http.HandleFunc("/test", packageHandler.TestHandler)

	// log.Fatalln(http.ListenAndServe(":8080", nil))
	// https://stackoverflow.com/questions/10175812/how-to-create-a-self-signed-certificate-with-openssl
	// Command: openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
	// https://serverfault.com/questions/366372/is-it-possible-to-generate-rsa-key-without-pass-phrase
	// No Passphrase for testing project. Use -nodes for no DES encryption for private key
	// Result Command: openssl req -nodes -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
	log.Fatalln(http.ListenAndServeTLS(":4443", certificates + "/cert.pem", certificates + "/key.pem", nil))

}
