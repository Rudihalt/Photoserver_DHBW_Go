package packageHandler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type PhotoTest struct {
	Name string `json:name`
	EXIFDate string `json:exifdate`
	Encoded string `json:encoded`
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	// q := r.URL.Query()
	// name := q.Get("name")
	// if len(name) == 0 {
	// 	name = "World"
	// }
	// responseString := "<html><body>Index Page<br>Hello " + name + "</body></html>"
	// w.Write([]byte(responseString)) // unbedingt Templates verwenden!

	tempFile := uploadFile(w, r)

	tempFileName := tempFile.Name()


	f, err := os.Open(tempFileName)
	if err != nil {
		log.Println("Error opening the file: " + tempFileName)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		// Handle error
	}

	log.Println("fmtName: " + fmtName)

	buffer := new(bytes.Buffer)

	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}

	var str = base64.StdEncoding.EncodeToString(buffer.Bytes())

	myPhoto := PhotoTest{
		Name: tempFileName,
		EXIFDate: ReadExifFromFile(tempFileName),
		Encoded: str,
	}

	log.Println("Name: " + myPhoto.Name + " ExifDate: " + myPhoto.EXIFDate + " Encoded: " + myPhoto.Encoded)
}

func uploadFile(w http.ResponseWriter, r *http.Request) *os.File {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20) // 10MB max

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)


	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()


	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")

	return tempFile
}

func ReadExifFromFile(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var checkString = string(f)[0:1000]

	re := regexp.MustCompile(`[0-9]{4}:[0-9]{2}:[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}`)

	return re.FindString(checkString)
}


