package packageObjects

type Album struct {
	Name string "json:name"
	Photos []Photo "json:photos"
}