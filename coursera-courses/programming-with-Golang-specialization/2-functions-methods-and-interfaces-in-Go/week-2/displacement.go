package main

import (
	"fmt"
	"os"
)

func main() {
	a := GetInputOf("acceleration")
	v0 := GetInputOf("initial velocity")
	s0 := GetInputOf("displacement")

	fn := GenDisplaceFn(a, v0, s0)
	t := GetInputOf("time")

	fmt.Println("The displacement is", fn(t))
}

func GetInputOf(varName string) float64 {
	var num float64

	fmt.Printf("Please, enter the value of %s:\n", varName)

	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return num
}

// function that return a function 
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5 * a * t * t + (v0 * t) + s0
	}
}