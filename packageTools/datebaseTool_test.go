/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageTools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPathExist(t *testing.T) {
	assert.False(t, PathExist("Somewhere/in/your/directory"), "Path does not exist")
	assert.True(t, PathExist("C:/Users"), "Path exists")
}
