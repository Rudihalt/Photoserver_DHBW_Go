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

func TestGetAllOrderElementsByUserNum(t *testing.T) {
	assert.Equal(t, 3, len(*GetAllOrderElementsByUser("test")))
	assert.Equal(t, 0, len(*GetAllOrderElementsByUser("test1")))

	order := *GetAllOrderElementsByUser("test")
	assert.Equal(t, 1, order[0].Amount)
	assert.Equal(t, "3x4", order[0].Format)

	assert.Equal(t, 2, order[1].Amount)
	assert.Equal(t, "16x9", order[1].Format)
}

func TestAddOrderElementAndDelete(t *testing.T) {
	assert.Equal(t, 3, len(*GetAllOrderElementsByUser("test")))

	orderElement1 := AddOrderElement("test", "hash1", 2, "3x5")

	assert.Equal(t, true, orderElement1 == nil)

	orderElement2 := AddOrderElement("test", "hash2", 2, "3x4")
	assert.Equal(t, false, orderElement2 == nil)
	assert.Equal(t, "3x4",  orderElement2.Format)
	assert.Equal(t, 2,  orderElement2.Amount)

	orderElement3 := AddOrderElement("test", "hash3", 2, "1x2")
	assert.Equal(t, false, orderElement3 == nil)
	assert.Equal(t, "1x2",  orderElement3.Format)
	assert.Equal(t, 2,  orderElement3.Amount)

	assert.Equal(t, 5, len(*GetAllOrderElementsByUser("test")))

	DeleteOrderElementByHash("test", orderElement2.ID)
	DeleteOrderElementByHash("test", orderElement3.ID)

	assert.Equal(t, 3, len(*GetAllOrderElementsByUser("test")))
}