package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Webrequests in GoLang.")

	res, err := http.Get(url)
	check(err)

	defer res.Body.Close()

	dataBytes, err := ioutil.ReadAll(res.Body)

	check(err)

	fmt.Println(string(dataBytes))

	fmt.Println("Ended")
}
