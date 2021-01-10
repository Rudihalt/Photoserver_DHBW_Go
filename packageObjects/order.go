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

// Returns Pointer of List of Order Elements. Reads Elements from json File and parses to struct
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

// saves List of Orders to the corresponding user in json file
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

// Adds a Order Element by creating a struct, appending to the current list and save. Additionally returning the new OrderElement
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

// Delete Order Element by id. Create new List of all elements without adding the one which matches
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

// Delete all -> Delete whole order json file for the user
func DeleteFullOrder(username string) {
	err := os.Remove(packageTools.GetWD() + "/static/data/order_" + username + ".json")
	if err != nil {
		panic(err)
	}
}
