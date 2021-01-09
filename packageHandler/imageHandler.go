package packageHandler

import (
	"log"
	"net/http"
	"photoserver/packageObjects"
)

type ImageShowData struct {
	Hash string
	Photo packageObjects.Photo
	Comments []packageObjects.Comment
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var cookie, _ = r.Cookie("csrftoken")
	if cookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		user := packageObjects.GetUserByToken(cookie.Value)
		NavData = NavBarData{Username: user.Username}
		err := NavTemplate.Execute(w, NavData)

		q := r.URL.Query()
		imageGet := q.Get("image")

		if imageGet == "" {
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		}

		allPhotos := packageObjects.GetAllPhotosByUser(user.Username)
		photo := packageObjects.GetPhotoByUserAndHash(allPhotos, imageGet)

		if photo == nil {
			http.Redirect(w, r, "/gallery", http.StatusSeeOther)
		}

		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				log.Fatalln(err)
				return
			}
			comment := r.FormValue("comment")
			log.Println("comment:", comment)

			if comment != "" && len(comment) > 0 {
				addComment := packageObjects.AddComment(user.Username, imageGet, comment)

				if addComment == nil {
					log.Println("Error adding comment")
				} else {
					log.Println("Successfully adding comment")
				}
			}
		}

		allComments := packageObjects.GetAllCommentsByUser(user.Username)
		comments := packageObjects.FilterAllCommentsByHash(allComments, imageGet)

		imageShowData := ImageShowData{
			Hash: photo.Hash,
			Photo: *photo,
			Comments: *comments,
		}

		err = ImageTemplate.Execute(w, imageShowData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}