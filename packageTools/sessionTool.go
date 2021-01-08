package packageTools

import (
	"crypto/sha256"
	"encoding/hex"
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

func CreateSalt() string {
	rand.Seed(time.Now().UnixNano())
	var randInt = rand.Intn(10000000)
	var strRandInd = strconv.Itoa(randInt)
	var salt = HashSHA(strRandInd)[0:8]

	return salt
}
