package packageObjects

type Album struct {
	Name string `json:"comment"`
	Date    string `json:"date"`
	Hash	string `json:"hash"`
}

/*
func GetAllAlbumsAlbumsAlbumsByUser(username string) *[]Comment {
	var comments []Comment
	var commentsFile []byte

	commentsFile, err := ioutil.ReadFile("static/data/comments_" + username + ".json")

	if err != nil {
		fmt.Println("Neue Datei anlegen: comments_" + username + ".json")
	}

	err = json.Unmarshal(commentsFile, &comments)

	if err != nil {
		// panic(err)
	}

	return &comments
}

func GetCommentByUserAndHash(comments *[]Comment, hash string) *Comment {


}

func saveComments(username string, comments *[]Comment) {

}

func AddComment(username string, hash string, commentStr string) {

}

func FilterAllCommentsByHash(comments *[]Comment, hash string) *[]Comment {

}
*/