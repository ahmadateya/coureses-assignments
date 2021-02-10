
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func first20Char(s string) string {
	runes := []rune(s)
	return string(runes[0:20])
}

func main() {
	var fileName string
	nameSli := make([]Name, 0)
	var nameObj Name

	fmt.Print("Enter file name: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// reading all the file at once
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		// fileScanner.Text() is a line of text
		// splitting the line to []string
		lineOfStrings := strings.Split(fileScanner.Text(), " ")

		if len(lineOfStrings[0]) > 20 {
			lineOfStrings[0] = first20Char(lineOfStrings[0])
		}
		if len(lineOfStrings[1]) > 20 {
			lineOfStrings[1] = first20Char(lineOfStrings[1])
		}

		nameObj.fname, nameObj.lname = lineOfStrings[0], lineOfStrings[1]
		nameSli = append(nameSli, nameObj)
	}

	for _, v := range nameSli {
		fmt.Println(v.fname, v.lname)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
}