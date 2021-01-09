/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageHandler

import (
	"log"
	"net/http"
	"photoserver/packageObjects"
	"strconv"
)

type OrderViewData struct {
	OrderElementsData []OrderElementData
}

type OrderElementData struct {
	Name string
	ImagePath string
	Amount int
	Format string
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		user := packageObjects.GetUserByToken(cookie.Value)

		q := r.URL.Query()
		deleteAll := q.Get("deleteAll")
		order := q.Get("order")
		deleteOne := q.Get("delete")

		if order == "1" {
			orderProcess() // TODO: NOCH IMPLEMENTIEREN
		}

		if deleteAll == "1" {
			packageObjects.DeleteFullOrder(user.Username)
		}

		if len(deleteOne) > 0 {
			deleteId, err := strconv.Atoi(deleteOne)
			if err != nil {
				log.Println("Id is not type of number in order delete. id=" + strconv.Itoa(deleteId))
			}

			/*allPhotos := packageObjects.GetAllPhotosByUser(user.Username)
			photo := packageObjects.GetPhotoByUserAndHash(allPhotos, deleteOne)

			if photo == nil {
				log.Println("Wrong Hash to delete from Order. Hash: " + deleteOne)
			}*/

			packageObjects.DeleteOrderElementByHash(user.Username, deleteId)
		}

		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)

		allPhotos := packageObjects.GetAllPhotosByUser(user.Username)

		var orderElementData []OrderElementData
		orderElements := packageObjects.GetAllOrderElementsByUser(user.Username)

		for _, orderElement := range *orderElements {
			photo := *packageObjects.GetPhotoByUserAndHash(allPhotos, orderElement.Hash)

			temp := OrderElementData{
				Name: photo.Name,
				ImagePath: photo.Path,
				Amount: orderElement.Amount,
				Format: orderElement.Format,
			}

			orderElementData = append(orderElementData, temp)
		}

		orderViewData := OrderViewData{
			OrderElementsData: orderElementData,
		}

		err = OrderTemplate.Execute(w, orderViewData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func orderProcess() {

}
