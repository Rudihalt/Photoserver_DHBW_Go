/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageObjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserFind(t *testing.T) {
	// Check Find functions
	assert.Equal(t, "test", GetUserByToken("d003b47ba0d9ade5f482e7e0").Username)
	assert.Equal(t, "test", GetUserByUsername("test").Username)

	assert.NotEqual(t, "abc", GetUserByToken("d003b47ba0d9ade5f482e7e0").Username)
	assert.NotEqual(t, "abc", GetUserByUsername("test").Username)

}

func TestUserSize(t *testing.T) {
	// check length of user (predefined file)
	assert.Equal(t, 2, len(*GetAllUsers()))
}

func TestUserPassword(t *testing.T) {
	// check checkpassword function. use username and password
	ok, token := CheckPassword("test", "123456")
	assert.Equal(t, true, ok)
	assert.Equal(t, "d003b47ba0d9ade5f482e7e0", token)

	ok, token = CheckPassword("test", "1234567")
	assert.Equal(t, false, ok)
	assert.Equal(t, "", token)
}

func TestCreateUser(t *testing.T) {
	// test create a user who already exists
	user := CreateUser("test", "123456")
	assert.Equal(t, true, user == nil)
}

func TestSessionToken(t *testing.T) {
	// Test of creating a ssession token. Should be different every time
	token := createSessionToken()
	assert.Equal(t, 24, len(token))

	token2 := createSessionToken()
	assert.Equal(t, 24, len(token2))

	assert.Equal(t, false, token == token2)
}

func TestUserExists(t *testing.T) {
	// Test if User already exists
	exists := UserExists("test")
	assert.Equal(t, true, exists)

	exists2 := UserExists("test123")
	assert.Equal(t, false, exists2)
}