package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

type Animal struct {
	food, locomotion_method, sound string
}

func (a *Animal) eat() {
	fmt.Println(a.food)
}

func (a *Animal) move() {
	fmt.Println(a.locomotion_method)
}

func (a *Animal) speak() {
	fmt.Println(a.sound)
}

func takeInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">")
	scanner.Scan()
	
	fields := strings.Fields(scanner.Text())

	for len(fields) != 2 {
		fmt.Println("Invalid input, please input animal name and action.")
		fmt.Printf(">")
		scanner.Scan()
		fields := strings.Fields(scanner.Text())

		if len(fields) == 2 {
			break
		}
	}

	return strings.ToLower(fields[0]), strings.ToLower(fields[1])
}

func main() {
	animals := map[string]Animal{}
	actions := "eat move speak" 
	animalNames := "cow bird snake" 

	// initializing the animals
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}

	for true {
		animal, action := takeInput()

		a, _ := animals[animal]

		if !strings.Contains(actions, action) || !strings.Contains(animalNames, animal) {
			fmt.Println("Invalid action or animal name, please use one of these actions (eat, move, speak) with those animals (cow, bird, snake)")
			continue
		}

		switch (action) {
			case "eat" :
				a.eat()
			case "move" :
				a.move()
			default:
				a.speak()
		}
	}
}