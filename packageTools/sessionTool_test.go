/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestHashSHA(t *testing.T) {
	// create three hashes
	same := "test"
	different := "tset"

	hash1 := HashSHA(same)
	hash2 := HashSHA(same)
	hash3 := HashSHA(different)

	assert.Equal(t, hash1, hash2)
	assert.NotEqual(t, hash1, hash3)
	assert.NotEqual(t, same, hash1)
}

func TestHashSHAFile(t *testing.T) {
	// create files
	same := "test.json"
	different := "tset.json"

	sameFile, _ := os.Create(same)
	sameFile.Close()
	differentFile, _ := os.Create(different)
	differentFile.Close()

	hash1 := HashSHA(sameFile.Name())
	hash2 := HashSHA(sameFile.Name())
	hash3 := HashSHA(differentFile.Name())

	assert.Equal(t, hash1, hash2)
	assert.NotEqual(t, hash1, hash3)
	assert.NotEqual(t, hash1, sameFile.Name())

	// remove created files
	err := os.Remove(sameFile.Name())
	if err != nil {
		log.Println("Could not delete file")
	}
	err = os.Remove(differentFile.Name())
	if err != nil {
		log.Println("Could not delete file")
	}
}

func TestCreateSalt(t *testing.T) {
	// create two salts
	salt1 := CreateSalt()
	salt2 := CreateSalt()

	// salt not the same and salt length = 8
	assert.NotEqual(t, salt1, salt2)
	assert.Equal(t, len(salt1), 8)
}

func TestGetRandomInt(t *testing.T) {
	// create two integers
	i1 := GetRandomInt()
	i2 := GetRandomInt()

	assert.NotEqual(t, i1, i2)
}

func TestCreateRandomString(t *testing.T) {
	// create two strings
	s1 := CreateRandomString()
	s2 := CreateRandomString()

	assert.NotEqual(t, s1, s2)
}
