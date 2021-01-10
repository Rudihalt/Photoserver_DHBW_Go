/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

// https://medium.com/better-programming/a-short-guide-to-hashing-in-go-e8bb0173e97e
// https://gobyexample.com/sha1-hashes
// https://austingwalters.com/building-a-web-server-in-go-salting-passwords/

// hash function
func HashSHA(str string) string {
	var bytes = []byte(str)

	// create new hash for str
	var hash = sha256.New()
	hash.Write(bytes)
	var code = hash.Sum(nil)
	var hashedString = hex.EncodeToString(code)

	return hashedString
}

// get hash for file
func HashSHAFile(filePath string) string {
	// read file to bytes
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Datei nicht gefunden!")
	}

	// create new hash and create it for the file
	var hash = sha256.New()
	hash.Write(bytes)
	var code = hash.Sum(nil)
	var hashedString = hex.EncodeToString(code)

	return hashedString
}

// create a salt which is a 8 char long random string
func CreateSalt() string {
	return CreateRandomString()[0:8]
}

// get random integer
func GetRandomInt() int {
	return rand.Int()
}

// create a random string
func CreateRandomString() string {
	var randInt = rand.Intn(10000000)
	var strRandInd = strconv.Itoa(randInt)
	var salt = HashSHA(strRandInd)

	return salt
}

// initialize rand
func Init() {
	rand.Seed(time.Now().UnixNano())
}
