/*
Matrikelnummern:
- 9122564
- 2227134
- 3886565
*/
package packageObjects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"photoserver/packageTools"
	"time"
)

type Comment struct {
	Comment string `json:"comment"`
	Date    string `json:"date"`
	Hash    string `json:"hash"`
}

func GetAllCommentsByUser(username string) *[]Comment {
	var comments []Comment
	var commentsFile []byte

	commentsFile, err := ioutil.ReadFile(packageTools.GetWD() + "/static/data/comments_" + username + ".json")

	if err != nil {
		fmt.Println("Neue Datei anlegen: comments_" + username + ".json")
	}

	err = json.Unmarshal(commentsFile, &comments)

	if err != nil {
		// panic(err)
	}

	return &comments
}

func saveComments(username string, comments *[]Comment) {
	commentJson, err := json.MarshalIndent(comments, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(packageTools.GetWD() + "static/data/comments_"+username+".json", commentJson, 0644)
	if err != nil {
		panic(err)
	}
}

func AddComment(username string, hash string, commentStr string) *Comment {
	currentComments := *GetAllCommentsByUser(username)

	currentTime := time.Now()
	timeFormatted := currentTime.Format("2006.01.02 15:04:05")

	comment := Comment{
		Comment: commentStr,
		Hash:    hash,
		Date:    timeFormatted,
	}

	currentComments = append(currentComments, comment)

	saveComments(username, &currentComments)

	return &comment
}

func FilterAllCommentsByHash(comments *[]Comment, hash string) *[]Comment {
	var hashComments []Comment
	for _, comment := range *comments {
		if comment.Hash == hash {
			hashComments = append(hashComments, comment)
		}
	}

	if len(hashComments) == 0 {
		return nil
	}

	return &hashComments
}
