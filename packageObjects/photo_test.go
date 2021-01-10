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
	assert.Equal(t, 3, len(*GetAllPhotosByUser("test")))
}

func TestPhotosPageAmount(t *testing.T) {
	assert.Equal(t, 1, GetPhotoPageAmount("test"))
}

func TestPhotosForPage(t *testing.T) {
	assert.Equal(t, 3, len(*GetPhotosForPage("test", 1)))
}

func TestGetPhotoByUserAndHash(t *testing.T) {
	photos := GetAllPhotosByUser("test")
	photo1 := GetPhotoByUserAndHash(photos, "fb891262c98f9725b54a613c6f1cbfb8c701bca2d386a55cd0d1b4966180549d")

	assert.Equal(t, false, photo1 == nil)
	assert.Equal(t, "DSC_0053.JPG", photo1.Name)

	photo2 := GetPhotoByUserAndHash(photos, "abc")

	assert.Equal(t, true, photo2 == nil)
}