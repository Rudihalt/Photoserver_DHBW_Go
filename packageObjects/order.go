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
	"photoserver/packageTools"
)

type OrderElement struct {
	ID int        `json:"id"`
	Hash   string `json:"hash"`
	Amount int    `json:"amount"`
	Format string `json:"format"`
}

func GetAllOrderElementsByUser(username string) *[]OrderElement {
	var orderElements []OrderElement
	var orderElementsFile []byte

	orderElementsFile, err := ioutil.ReadFile(packageTools.GetWD() + "/static/data/order_" + username + ".json")

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

	err = ioutil.WriteFile(packageTools.GetWD() + "/static/data/order_"+username+".json", orderElementsJson, 0644)
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
		ID: packageTools.GetRandomInt(),
		Hash:   hash,
		Amount: amount,
		Format: format,
	}

	currentOrderElements = append(currentOrderElements, orderElement)

	saveOrderElements(username, &currentOrderElements)

	return &orderElement
}

func DeleteOrderElementByHash(username string, id int) {
	var newOrderElements []OrderElement
	currentOrderElements := *GetAllOrderElementsByUser(username)

	for _, orderElement := range currentOrderElements {
		if orderElement.ID != id {
			newOrderElements = append(newOrderElements, orderElement)
		}
	}

	saveOrderElements(username, &newOrderElements)
}


func DeleteFullOrder(username string) {
	err := os.Remove(packageTools.GetWD() + "/static/data/order_" + username + ".json")
	if err != nil {
		panic(err)
	}
}
