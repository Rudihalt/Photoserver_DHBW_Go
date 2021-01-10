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
	"path"
	"testing"
)

func TestPathExist(t *testing.T) {
	// check for not existing path
	assert.False(t, PathExist("Somewhere/in/your/directory/I/hope/this/path/doesnt/exist"), "Path does not exist")

	// get current workdirectory for a exist path
	dir, _ := os.Getwd()
	assert.True(t, PathExist(dir), "Path exists")
}

func TestCreateDirIfNotExists(t *testing.T) {
	// initialize a not existing directory
	dir, _ := os.Getwd()
	dir = path.Join(dir, "test_directory")
	CreateDirIfNotExists(dir)
	// check if direcotry exist
	exist := PathExist(dir)
	assert.True(t, exist)
	// remove directory
	err := os.Remove(dir)
	if err != nil {
		log.Println("Could not delete directory")
	}
}
