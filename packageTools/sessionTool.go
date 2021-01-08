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

type User struct {
	Username string
	Password string
	Salt     string
}

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

func GetUser(username string) User {
	return User{
		Username: "User",
		Salt:     "SALT",
		Password: "123456",
	}
}

func SaveUser(user User) {

}

func checkPassword(username string, passwordInput string, passwordDatabase string) bool {
	//var user = GetUser(username)
	// if user == nil {
	//	fmt.Println("No User")
	//}

	return true
}
