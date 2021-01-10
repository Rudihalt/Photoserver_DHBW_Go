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

func HashSHA(str string) string {
	var bytes = []byte(str)

	var hash = sha256.New()
	hash.Write(bytes)
	var code = hash.Sum(nil)
	var hashedString = hex.EncodeToString(code)

	return hashedString
}

func HashSHAFile(filePath string) string {
	bytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Datei nicht gefunden!")
	}

	var hash = sha256.New()
	hash.Write(bytes)
	var code = hash.Sum(nil)
	var hashedString = hex.EncodeToString(code)

	return hashedString
}

func CreateSalt() string {
	return CreateRandomString()[0:8]
}

func GetRandomInt() int {
	return rand.Int()
}

func CreateRandomString() string {
	var strRandInd = strconv.Itoa(randInt)
	var salt = HashSHA(strRandInd)

	return salt
}

func Init() {
	rand.Seed(time.Now().UnixNano())
}
