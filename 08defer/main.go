package main

import "fmt"

func main() {
	fmt.Println("Defer in golang")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")

	fmt.Println("Hello")
	deferLoop()

}

func deferLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("value is:", i)
	}
}
