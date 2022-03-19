package main

import "fmt"

func main() {
	var name string = "Gopher"
	fmt.Println("Hello", name)
	fmt.Printf("Type of variable is %T\n", name)

	var age int = 25
	fmt.Println("Age is", age)
	fmt.Printf("Type of variable is %T\n", age)

	var isCool bool = true
	fmt.Println("Is Gopher cool?", isCool)
	fmt.Printf("Type of variable is %T\n", isCool)

	var size float32 = 1.3
	fmt.Println("Size of Gopher is", size)
	fmt.Printf("Type of variable is %T\n", size)

	var gopher = "Gopher"
	fmt.Println("Gopher is", gopher)
	fmt.Printf("Type of variable is %T\n", gopher)

	newVariable := "New Variable"
	fmt.Println("New Variable is", newVariable)
	fmt.Printf("Type of variable is %T\n", newVariable)

}
