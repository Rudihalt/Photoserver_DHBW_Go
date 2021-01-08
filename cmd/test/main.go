package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"photoserver/packageTools"
	"regexp"
	"time"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./static/images")))

	fs := http.FileServer(http.Dir("./static/images"))
	http.Handle("/images/", http.StripPrefix("/images", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testing() {
	fmt.Println(ReadExifFromFile("static/images/p2.jpg"))
	fmt.Println(GetDateObjectFromString(ReadExifFromFile("static/images/p2.jpg")).Minute)


	packageTools.InitCache(2)
	cache := packageTools.GetGlobalCache()

	cache.InitLru(2)
	cache.Put(2, "a")
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(1))
	cache.Put(1, "b")
	cache.Put(1, "c")
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	cache.Put(8, "d")
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(8))

	userPtr := createUser("testuser", "123456")
	if userPtr == nil {
		log.Println("User already exist!")
	}
	SavePhoto("photo.jpg", userPtr.Username, "ABCDEF", "2020:10:29 13:34:25")
}

type Date struct {
	Format string
	Year string
	Month string
	Day string
	Hour string
	Minute string
	Second string
}

func GetDateObjectFromString(input string) Date {
	date := Date{
		Format: input,
		Year: input[0:4],
		Month: input[5:7],
		Day: input[8:10],
		Hour: input[11:13],
		Minute: input[14:16],
		Second: input[17:19],
	}
	return date
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

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	Token string `json:"token"`
	Photos []string `photos:"Photos"`
}

var users []User

func GetAllUsers() *[]User {
	return &users
}

func readUsers() {
	userData, err := ioutil.ReadFile( "static/data/users.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(userData, &users)
	if err != nil {
		panic(err)
	}
}

func saveUsers() {
	usersJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/data/users.json", usersJson, 0644)
	if err != nil {
		panic(err)
	}
}

func checkPassword(username string, password string) (bool, string) {
	readUsers()

	user := getUserByUsername(username)
	hashedInputPassword := packageTools.HashSHA(user.Salt + password)

	if hashedInputPassword == user.Password {
		return true, user.Token
	}

	return false, ""
}

func addPhotoToUser(username string, photoHash string) {
	user := getUserByUsername(username)
	userPhotos := user.Photos

	fmt.Println(user.Photos)

	newUserPhotos := append(userPhotos, photoHash)
	user.Photos = newUserPhotos

	saveUsers()
}

func createUser(username string, password string) *User {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000000000)
	salt := packageTools.CreateSalt()
	passwordHash := packageTools.HashSHA(salt + password)
	token := createSessionToken()

	user := User{
		Id: id,
		Username: username,
		Password: passwordHash,
		Salt: salt,
		Token: token,
	}

	readUsers()

	if userExists(username) {
		return nil
	}

	users = append(users, user)
	saveUsers()

	return &user
}

func createSessionToken() string {
	token := ""
	for i := 1; i < 5; i++ {
		token += packageTools.CreateSalt()
	}

	return token
}

func userExists(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}

	return false
}

func getUserByToken(token string) *User {
	for _, user := range users {
		if user.Token == token {
			return &user
		}
	}

	return nil
}

func getUserByUsername(username string) *User {
	for _, user := range users {
		if user.Username == username {
			return &user
		}
	}

	return nil
}




type Photo struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Hash string `json:"hash"`
	Encoded string `json:encoded`
	exifDate string `json:exifdate`
}

func GetPhotoByHash(hash string) *Photo {
	var photo Photo
	photoFile, err := ioutil.ReadFile( "static/data/photo_" + hash + ".json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(photoFile, &photo)
	if err != nil {
		panic(err)
	}

	return &photo
}

func SavePhoto(name string, username string, encoded string, exifdate string) *Photo {
	hash := packageTools.HashSHA(encoded)

	if _, err := os.Stat("static/data/photo_" + hash + ".json"); os.IsNotExist(err) == false {
		return nil
	}

	photo := Photo {
		Name: name,
		Username: username,
		Hash: hash,
		Encoded: encoded,
		exifDate: exifdate,
	}

	photoJson, err := json.MarshalIndent(photo, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/data/photo_" + hash + ".json", photoJson, 0644)
	if err != nil {
		panic(err)
	}

	addPhotoToUser(username, hash)

	return &photo
}