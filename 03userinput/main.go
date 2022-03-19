package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	message := "Welcome to Go!"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")

	input, _ := reader.ReadString('\n')

	inputToString, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inputToString)
		fmt.Println("Input + 1 is ", inputToString+1)
	}
}
