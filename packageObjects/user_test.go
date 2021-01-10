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
	assert.Equal(t, "test", GetUserByToken("d003b47ba0d9ade5f482e7e0").Username)
	assert.Equal(t, "test", GetUserByUsername("test").Username)

	assert.NotEqual(t, "abc", GetUserByToken("d003b47ba0d9ade5f482e7e0").Username)
	assert.NotEqual(t, "abc", GetUserByUsername("test").Username)

}

func TestUserSize(t *testing.T) {
	assert.Equal(t, 2, len(*GetAllUsers()))
}

func TestUserPassword(t *testing.T) {
	ok, token := CheckPassword("test", "123456")
	assert.Equal(t, true, ok)
	assert.Equal(t, "d003b47ba0d9ade5f482e7e0", token)

	ok, token = CheckPassword("test", "1234567")
	assert.Equal(t, false, ok)
	assert.Equal(t, "", token)
}

func TestCreateUser(t *testing.T) {
	user := CreateUser("test", "123456")
	assert.Equal(t, true, user == nil)
}

func TestSessionToken(t *testing.T) {
	token := createSessionToken()
	assert.Equal(t, 24, len(token))

	token2 := createSessionToken()
	assert.Equal(t, 24, len(token2))

	assert.Equal(t, false, token == token2)
}

func TestUserExists(t *testing.T) {
	exists := UserExists("test")
	assert.Equal(t, true, exists)

	exists2 := UserExists("test123")
	assert.Equal(t, false, exists2)
}