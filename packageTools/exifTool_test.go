package packageTools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDateTime(t *testing.T) {
	byteimg := GetByteImage()
	date, err := GetDateTime(byteimg)
	if assert.Nil(t, err) {
		assert.Equal(t, date, "2020:10:27")
	}

	b := []byte{0, 0, 0, 0, 0, 0, 0}
	date, err = GetDateTime(b)
	assert.Error(t, err, "No Exif-Header found")
}
