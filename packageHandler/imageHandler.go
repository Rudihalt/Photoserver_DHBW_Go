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

type ImageShowData struct {
	Hash     string
	Photo    packageObjects.Photo
	Comments []packageObjects.Comment
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// get token from cookie
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		// if no cookie is set redirect to login site
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// get user from cookie token and set navigation bar with username
		user := packageObjects.GetUserByToken(cookie.Value)
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)

		q := r.URL.Query()
		imageGet := q.Get("image")

		// if no photo in get parameter redirect to gallery
		if imageGet == "" {
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		}

		allPhotos := packageObjects.GetAllPhotosByUser(user.Username)
		photo := packageObjects.GetPhotoByUserAndHash(allPhotos, imageGet)

		// if photo not exist redirect to gallery
		if photo == nil {
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		}

		// post requests
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				log.Fatalln(err)
				return
			}
			// get comment value from form
			comment := r.FormValue("comment")
			log.Println("comment:", comment)

			// if comment is not empty and something is written
			// add comment to the image
			if comment != "" && len(comment) > 0 {
				addComment := packageObjects.AddComment(user.Username, imageGet, comment)

				if addComment == nil {
					log.Println("Error adding comment")
				} else {
					log.Println("Successfully adding comment")
				}
			}

			// order values:
			// amount -> how often the image will be downloaded
			orderAmount := r.FormValue("orderAmount")
			log.Println("orderAmount:", orderAmount)

			// format -> the format of the image
			orderFormat := r.FormValue("orderFormat")
			log.Println("orderFormat:", orderFormat)

			// check if the order amount is not empty
			if orderAmount != "" && len(orderAmount) > 0 {

				amount, err := strconv.Atoi(orderAmount)
				if err != nil {
					log.Println("Error Amount no number")
				}

				orderElement := packageObjects.AddOrderElement(user.Username, imageGet, amount, orderFormat)

				if orderElement == nil {
					log.Println("Error adding order")
				} else {
					log.Println("Successfully adding order")
				}
			}
		}

		// get all comments for user and filter them
		allComments := packageObjects.GetAllCommentsByUser(user.Username)
		comments := packageObjects.FilterAllCommentsByHash(allComments, imageGet)

		// create comment data which will be send with the imageShowData object to the image template
		var dataComments []packageObjects.Comment
		if comments != nil {
			dataComments = *comments
		}

		imageShowData := ImageShowData{
			Hash:     photo.Hash,
			Photo:    *photo,
			Comments: dataComments,
		}

		// hand over the ImageShowData to the tetmplate
		err = ImageTemplate.Execute(w, imageShowData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
