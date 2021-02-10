package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	floatingNumberREGEX = regexp.MustCompile(`^(\d+)\.?(\d*)$`)
)

// GetIntegerFromFloatingNumber Parse a FloatingNumber from a string and obtain the integer part
func GetIntegerFromFloatingNumber(number string) (int64, error) {
	matched := floatingNumberREGEX.FindStringSubmatch(number)

	if len(matched) == 0 {
		return 0, fmt.Errorf("Error parsing: %s is not a Floating Number", number)
	}

	return strconv.ParseInt(matched[1], 10, 64)
}

func truncateTask() {
	number, err := readTextFromConsole("floating point number: ", false)
	//number := "123.45"
	//var err error

	if err != nil {
		log.Fatal("Error reading text from console: ", err)
	} else {
		var integer, err = GetIntegerFromFloatingNumber(strings.Trim(number, "\n"))

		if err != nil {
			log.Fatal("Error parsing FloatingNumber: ", err)
		} else {
			fmt.Println(fmt.Sprintf("Integer parsed: %d", integer))
		}
	}
}