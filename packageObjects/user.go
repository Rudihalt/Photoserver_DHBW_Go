package packageObjects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"photoserver/packageTools"
	"time"
)

type User struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Salt     string   `json:"salt"`
	Token    string   `json:"token"`
	Photos   []string `photos:"Photos"`
}

var users []User

func GetAllUsers() *[]User {
	return &users
}

func readUsers() {
	userData, err := ioutil.ReadFile("static/data/users.json")
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

func CheckPassword(username string, password string) (bool, string) {
	readUsers()

	user := getUserByUsername(username)
	if user != nil {
		hashedInputPassword := packageTools.HashSHA(user.Salt + password)

		if hashedInputPassword == user.Password {
			return true, user.Token
		}
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

func CreateUser(username string, password string) *User {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000000000)
	salt := packageTools.CreateSalt()
	passwordHash := packageTools.HashSHA(salt + password)
	token := createSessionToken()

	user := User{
		Id:       id,
		Username: username,
		Password: passwordHash,
		Salt:     salt,
		Token:    token,
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

func UserExists(username string) bool {
	readUsers()
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}

	return false
}

func GetUserByToken(token string) *User {
	readUsers()
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
