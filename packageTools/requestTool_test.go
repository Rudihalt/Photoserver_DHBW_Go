package packageTools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckHostExist(t *testing.T) {
	// host exist
	host := "google.com:80"
	exist := CheckHost(host)
	assert.Equal(t, true, exist)
}

func TestCheckHostNotExist(t *testing.T) {
	// host does not exist
	host := "gibt-es.nicht:80"
	exist := CheckHost(host)
	assert.Equal(t, false, exist)
}

func TestCreateFileUploadRequest(t *testing.T) {
	_, err := createFileUploadRequest("localhost", "Test.png", "username")
	assert.Error(t, err, "No JPEG detected")
}
