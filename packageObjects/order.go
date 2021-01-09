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

type OrderElement struct {
	Hash   string `json:"hash"`
	Amount int    `json:"amount"`
	Format string `json:"format"`
}

func GetAllOrderElementsByUser(username string) *[]OrderElement {
	var orderElements []OrderElement
	var orderElementsFile []byte

	orderElementsFile, err := ioutil.ReadFile("static/data/order_" + username + ".json")

	err = json.Unmarshal(orderElementsFile, &orderElements)

	if err != nil {
		// panic(err)
	}

	return &orderElements
}

func saveOrderElements(username string, orderElements *[]OrderElement) {
	orderElementsJson, err := json.MarshalIndent(orderElements, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("static/data/order_"+username+".json", orderElementsJson, 0644)
	if err != nil {
		panic(err)
	}
}

func AddOrderElement(username string, hash string, amount int, format string) *OrderElement {
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

func deleteOrderElementByHash(username string, hash string) {
	var newOrderElements []OrderElement
	currentOrderElements := *GetAllOrderElementsByUser(username)

	for _, orderElement := range currentOrderElements {
		if orderElement.Hash != hash {
			newOrderElements = append(newOrderElements, orderElement)
		}
	}

	saveOrderElements(username, &newOrderElements)
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func removeElementByIndex(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func deleteFullOrder(username string) {
	err := os.Remove("static/data/order_" + username + ".json")
	if err != nil {
		panic(err)
	}
}
