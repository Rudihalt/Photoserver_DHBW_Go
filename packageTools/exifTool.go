package packageTools

import (
	"bytes"
	"errors"
	"log"
	"regexp"
	"time"
)

// b := bytearray after ioutil.ReadFile(FILE)
func GetDateTime(b []byte) (date string, err error) {
	// check if exif header exist
	if (bytes.Equal(b[0:4], []byte{0xff, 0xd8, 0xff, 0xe1}) || bytes.Equal(b[0:4], []byte{0xff, 0xd8, 0xff, 0xe0})) &&
		(string(b[6:10]) == "Exif" || string(b[6:10]) == "JFIF") {

		s := string(b)
		// 1999:03:20 o 2020:12:31
		dateRegex := regexp.MustCompile("([12]\\d{3}:(0[1-9]|1[0-2]):(0[1-9]|[12]\\d|3[01]))")
		date = dateRegex.FindString(s)
		if date == "" {
			currentTime := time.Now()
			date = currentTime.Format("2006:01:02")
			log.Println("No date found -> date = ", date)
		}
		return date, nil
	}
	return "", errors.New("No Exif-Header found")
}
