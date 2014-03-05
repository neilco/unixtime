package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Usage() {
	fmt.Println("Usage: ", os.Args[0], "[datetime]\n")
	fmt.Println("Converts a date or date & time into a UNIX timestamp\n")
	fmt.Println("datetime: The date or date & time to parse. Can be one of the following formats:\n")
	fmt.Println("\t  now")
	fmt.Println("\t  YYYY-MM-DD")
	fmt.Println("\t  YYYY-MM-DD HH:MM")
	fmt.Println("\t  YYYY-MM-DD HH:MM:SS")
	fmt.Println("\t  YYYY-MM-DD HH:MM:SS.SSS")
	fmt.Println("\t  YYYY-MM-DDTHH:MMZ")
	fmt.Println("\t  YYYY-MM-DDTHH:MM:SSZ")
	fmt.Println("\t  YYYY-MM-DDTHH:MM:SS.SSSZ")
	fmt.Println("\nIf no argument is supplied, \"now\" is assumed.\n")
}

func main() {
	var t time.Time
	var err error

	dateString := strings.Join(os.Args[1:], " ")

	if dateString == "-h" || dateString == "--help" || dateString == "help" {
		Usage()
		os.Exit(0)
	} else if dateString == "now" || dateString == "" {
		t = time.Now()
	} else {
		layouts := map[int]string{
			10: "2006-01-02",
			16: "2006-01-02 15:04",
			17: "2006-01-02T15:04Z",
			19: "2006-01-02 15:04:05",
			20: "2006-01-02T15:04:05Z",
			23: "2006-01-02 15:04:05.000",
			24: "2006-01-02T15:04:05.000Z",
		}

		layout := layouts[len(dateString)]

		t, err = time.Parse(layout, string(dateString))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(t.Unix())
}
