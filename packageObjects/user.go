/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageObjects

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"photoserver/packageTools"
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Token    string `json:"token"`
}

var users *[]User

// returns users
func GetAllUsers() *[]User {
	return users
}

// sets users
func SetAllUsers(usersParam *[]User) {
	users = usersParam
}

// reads all users from json file. parsing to struct array
func readUsers() {
	userData, err := ioutil.ReadFile(packageTools.GetWD() + "/static/data/users.json")
	if err != nil {
		f, _ := os.Create(packageTools.GetWD() + "/static/data/users.json")
		f.WriteString("[]")
		f.Close()
		userData, _ = ioutil.ReadFile(packageTools.GetWD() + "/static/data/users.json")
	}

	err = json.Unmarshal(userData, &users)
	if err != nil {
		panic(err)
	}
}

// saves users to json file.
func saveUsers() {
	usersJson, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(packageTools.GetWD() + "static/data/users.json", usersJson, 0644)
	if err != nil {
		panic(err)
	}
}

// function to check password. use hash and "salting" for password check.
func CheckPassword(username string, password string) (bool, string) {
	readUsers()

	user := GetUserByUsername(username)
	if user != nil {
		hashedInputPassword := packageTools.HashSHA(user.Salt + password)

		if hashedInputPassword == user.Password {
			return true, user.Token
		}
	}

	return false, ""
}

// Creates a user. create struct and hash password (with generated salt) and add user to users.json file
func CreateUser(username string, password string) *User {
	readUsers()

	if GetUserByUsername(username) != nil {
		return nil
	}

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

	currentUsers := *GetAllUsers()
	currentUsers = append(currentUsers, user)

	SetAllUsers(&currentUsers)
	saveUsers()

	return &user
}

// Creates a random session token. 24 characters
func createSessionToken() string {
	token := packageTools.CreateRandomString()[0:24]

	return token
}

// check if user exists: check if a username in list matches
func UserExists(username string) bool {
	readUsers()
	for _, user := range *GetAllUsers() {
		if user.Username == username {
			return true
		}
	}

	return false
}

// Get user by token by matching user list
func GetUserByToken(token string) *User {
	readUsers()
	for _, user := range *GetAllUsers() {
		if user.Token == token {
			return &user
		}
	}

	return nil
}

// Get user by username by matching user list
func GetUserByUsername(username string) *User {
	for _, user := range *GetAllUsers() {
		if user.Username == username {
			return &user
		}
	}

	return nil
}
