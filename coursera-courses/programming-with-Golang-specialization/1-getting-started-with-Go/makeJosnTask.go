package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	 Name string `json:"Name"`
	 Address string `json:"Address"`
}

func main() {

	var user Person
	user.Name, _ = readTextFromConsole("Enter your name: ", true)
	user.Address, _ = readTextFromConsole("Enter your address: ", true)

	// Create JSON object from the struct
	jsonObj, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonObj))
}