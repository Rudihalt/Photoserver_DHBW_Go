package packageTools

import "os"

// checks if path or file exist
func PathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
