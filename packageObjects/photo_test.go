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

func TestPhotosSize(t *testing.T) {
	// Check length of photos. predefined file
	assert.Equal(t, 4, len(*GetAllPhotosByUser("test")))
	savePhotos("test", GetAllPhotosByUser("test"))
	photo := SavePhoto("file.jpg", "test", "/static/file.jpg", "date")
	assert.Equal(t, true, photo == nil)
}

func TestPhotosPageAmount(t *testing.T) {
	// check GetPhotoPageAmount
	assert.Equal(t, 1, GetPhotoPageAmount("test"))
}

func TestPhotosForPage(t *testing.T) {
	// check GetPhotosForPage
	assert.Equal(t, 4, len(*GetPhotosForPage("test", 1)))
}

func TestGetPhotoByUserAndHash(t *testing.T) {
	// Get Photo by user and hash. Use predefined json file
	photos := GetAllPhotosByUser("test")
	photo1 := GetPhotoByUserAndHash(photos, "fb891262c98f9725b54a613c6f1cbfb8c701bca2d386a55cd0d1b4966180549d")

	assert.Equal(t, false, photo1 == nil)
	assert.Equal(t, "DSC_0053.JPG", photo1.Name)

	photo2 := GetPhotoByUserAndHash(photos, "abc")

	assert.Equal(t, true, photo2 == nil)
}