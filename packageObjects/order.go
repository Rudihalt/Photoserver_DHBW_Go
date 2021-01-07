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

func createNewOrder(userID int, albumIDs []int, photoIDs []int) *Order {
	currentTime := time.Now()
	timeFormatted := currentTime.Format("2006.01.02 15:04:05")

	randInt := getNewOrderID()

	var order = Order{
		ID: randInt,
		UserID: userID,
		Date: timeFormatted,
		AlbumIDs: albumIDs,
		PhotoIDs: photoIDs,
	}

	return &order
}

func getNewOrderID() int {
	randInt := rand.Intn(100)
	return randInt
}