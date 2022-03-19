package main

import "fmt"

func main() {
	fmt.Println("Functions in GoLang")
	result := add(2, 3)

	fmt.Println("result is:", result)

	fmt.Println("Pro result:", proAdder(2, 3, 4, 2, 5, 2, 5, 5))

}

func add(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
