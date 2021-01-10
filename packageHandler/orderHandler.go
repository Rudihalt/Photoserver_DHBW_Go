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
	"photoserver/packageTools"
	"strconv"
)

type OrderViewData struct {
	OrderElementsData []OrderElementData
	ZipPath           string
}

type OrderElementData struct {
	ID        int
	Name      string
	ImagePath string
	Amount    int
	Format    string
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// get token from cookie
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		// if no cookie is set redirect to login site
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// get user from cookie token and set navigation bar with username
		user := packageObjects.GetUserByToken(cookie.Value)

		// get all parameters from the get parameter
		q := r.URL.Query()
		deleteAll := q.Get("deleteAll")
		order := q.Get("order")
		deleteOne := q.Get("delete")

		// create new OrderViewData variable
		var orderViewData = OrderViewData{}
		if order == "1" {
			// if order was pressed run orderProcess function which creates
			// the zipFile
			path := orderProcess(user.Username)
			if path != "" {
				orderViewData.ZipPath = path
			}
		}

		if deleteAll == "1" {
			// if the button deleteAll was pressed remove All OrderElements from the Orderlist
			packageObjects.DeleteFullOrder(user.Username)
		}

		// if deleteOne is bigger than 0 a element is deleted from the orderlist
		if len(deleteOne) > 0 {
			// convert string to int
			deleteId, err := strconv.Atoi(deleteOne)
			if err != nil {
				log.Println("Id is not type of number in order delete. id=" + strconv.Itoa(deleteId))
			}

			/*allPhotos := packageObjects.GetAllPhotosByUser(user.Username)
			photo := packageObjects.GetPhotoByUserAndHash(allPhotos, deleteOne)

			if photo == nil {
				log.Println("Wrong Hash to delete from Order. Hash: " + deleteOne)
			}*/

			// delete the selected orderelement
			packageObjects.DeleteOrderElementByHash(user.Username, deleteId)
		}

		// create NavData to be hand over to the navbar
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)

		// get all photos from user
		allPhotos := packageObjects.GetAllPhotosByUser(user.Username)

		// get all orderelements to be shown on the page
		var orderElementData []OrderElementData
		orderElements := packageObjects.GetAllOrderElementsByUser(user.Username)

		// create the orderElements
		for _, orderElement := range *orderElements {
			photo := *packageObjects.GetPhotoByUserAndHash(allPhotos, orderElement.Hash)

			temp := OrderElementData{
				ID:        orderElement.ID,
				Name:      photo.Name,
				ImagePath: photo.Path,
				Amount:    orderElement.Amount,
				Format:    orderElement.Format,
			}

			orderElementData = append(orderElementData, temp)
		}
		orderViewData.OrderElementsData = orderElementData

		// hand over the OrderViewData to the template
		err = OrderTemplate.Execute(w, orderViewData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func orderProcess(username string) string {
	photos := packageObjects.GetAllPhotosByUser(username)
	order := packageObjects.GetAllOrderElementsByUser(username)

	var zipItems []packageTools.ZipItem

	log.Println(username, "places an order")
	for _, element := range *order {
		photo := packageObjects.GetPhotoByUserAndHash(photos, element.Hash)

		item := packageTools.ZipItem{Name: photo.Name, Path: photo.Path, Format: element.Format, Amount: element.Amount}
		log.Println("-", strconv.Itoa(item.Amount)+"x", item.Name, "in", item.Format)
		zipItems = append(zipItems, item)
	}
	zipFile, err := packageTools.CreateZipFile(zipItems, username)
	if err == nil {
		packageObjects.DeleteFullOrder(username)
		return zipFile
	}
	return ""
}
