package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Working with Files")
	content := "This line is written in file using GoLang"

	file, err := os.Create("file.txt")
	check(err)

	length, err := io.WriteString(file, content)
	check(err)

	fmt.Println(length)
	defer file.Close()

	readFile(file.Name())

}

func readFile(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)

	fmt.Println("File content:", string(file))

}
