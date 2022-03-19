package main

import "fmt"

func main() {
	fmt.Println("Learnig about pointers")

	myNumber := 23

	myPointer := &myNumber

	fmt.Println("Value of my number pointer ", myPointer)
	fmt.Println("Value of my number using pointer ", *myPointer)

	modifiedPointer := *myPointer + 2

	fmt.Println("modifiedPointer is", modifiedPointer)

}
