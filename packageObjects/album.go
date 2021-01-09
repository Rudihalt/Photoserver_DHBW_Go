/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageObjects

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Album struct {
	Name string `json:"comment"`
	Date string `json:"date"`
	Hash string `json:"hash"`
}

type AlbumElement struct {
	Hash string `json:"hash"`
}

func GetAllAlbumElementsByUser(username string) *[]OrderElement {
	var orderElements []OrderElement
	var orderElementsFile []byte

	orderElementsFile, err := ioutil.ReadFile("static/data/order_" + username + ".json")

	err = json.Unmarshal(orderElementsFile, &orderElements)

	if err != nil {
		// panic(err)
	}

	return &orderElements
}

func saveAlbumElements(username string, orderElements *[]OrderElement) {
	orderElementsJson, err := json.MarshalIndent(orderElements, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/data/order_"+username+".json", orderElementsJson, 0644)
	if err != nil {
		panic(err)
	}
}

func AddAlbumElement(username string, hash string, amount int, format string) *OrderElement {
	if format != "3x4" && format != "16x9" && format != "1x2" {
		return nil
	}

	currentOrderElements := *GetAllOrderElementsByUser(username)

	var orderElement = OrderElement{
		Hash:   hash,
		Amount: amount,
		Format: format,
	}

	currentOrderElements = append(currentOrderElements, orderElement)

	saveOrderElements(username, &currentOrderElements)

	return &orderElement
}

func deleteAlbumElementByHash(username string, hash string) {
	var newOrderElements []OrderElement
	currentOrderElements := *GetAllOrderElementsByUser(username)

	for _, orderElement := range currentOrderElements {
		if orderElement.Hash != hash {
			newOrderElements = append(newOrderElements, orderElement)
		}
	}

	saveOrderElements(username, &newOrderElements)
}

func deleteFullAlbum(username string) {
	err := os.Remove("static/data/order_" + username + ".json")
	if err != nil {
		panic(err)
	}
}
