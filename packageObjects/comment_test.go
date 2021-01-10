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

func TestCommentsSize(t *testing.T) {
	assert.Equal(t, 4, len(*GetAllCommentsByUser("test")))
	assert.Equal(t, 0, len(*GetAllCommentsByUser("test1")))
}

func TestFilterCommentByHash(t *testing.T) {
	allComments := GetAllCommentsByUser("test")

	comments := FilterAllCommentsByHash(allComments, "fb891262c98f9725b54a613c6f1cbfb8c701bca2d386a55cd0d1b4966180549d")

	assert.Equal(t, true, comments != nil)

	assert.Equal(t, "Gutes Bild", (*comments)[0].Comment)
	assert.Equal(t, "2021.01.10 16:34:37", (*comments)[0].Date)

	assert.Equal(t, "Super Bild", (*comments)[1].Comment)
	assert.Equal(t, "2021.01.10 16:34:49", (*comments)[1].Date)

	assert.Equal(t, true, FilterAllCommentsByHash(allComments, "abc") == nil)
}