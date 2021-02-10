package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)


func readTextFromConsole(message string, withoutBreakLine bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)
	text, err := reader.ReadString('\n')

	if withoutBreakLine {
		strings.Trim(text, " \n")
	}
	return text, err
}

func main(){
	text, err := readTextFromConsole("please enter the text: ", false)

	if err != nil {
		fmt.Printf("Error reading text from console: %s", err)
		log.Fatal("Error reading text from console: ", err)
	} else {
		text = strings.ToLower(text)
		fmt.Printf(text)
		ianRegex := regexp.MustCompile(`\b(\w*iaa*?n\w*)\b`)
		matched := ianRegex.FindStringSubmatch(text)

		if len(matched) == 0 {
			fmt.Printf("Not Found!")
		} else {
			fmt.Printf("Found!")
		}
	}
}

// alternative implementation
//func main() {
//	var input string
//
//	fmt.Printf("Enter string: ")
//	fmt.Scan(&input)
//	input = strings.ToLower(input)
//	fmt.Printf("string is %s.\n", input)
//
//	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
//		fmt.Println("Found!")
//	} else {
//		fmt.Println("Not Found!")
//	}
//}
