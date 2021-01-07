package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	fmt.Println(ReadExifFromFile("static/images/p2.jpg"))
	fmt.Println(GetDateObjectFromString(ReadExifFromFile("static/images/p2.jpg")).Minute)

}

type Date struct {
	Format string
	Year string
	Month string
	Day string
	Hour string
	Minute string
	Second string
}

func GetDateObjectFromString(input string) Date {
	date := Date{
		Format: input,
		Year: input[0:4],
		Month: input[5:7],
		Day: input[8:10],
		Hour: input[11:13],
		Minute: input[14:16],
		Second: input[17:19],
	}
	return date
}

func ReadExifFromFile(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var checkString = string(f)[0:1000]

	re := regexp.MustCompile(`[0-9]{4}:[0-9]{2}:[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}`)

	return re.FindString(checkString)
}