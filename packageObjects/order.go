package packageObjects

import (
	"math/rand"
	"time"
)

type Order struct {
	ID int
	UserID int
	Date string
	AlbumIDs []int
	PhotoIDs []int
}
