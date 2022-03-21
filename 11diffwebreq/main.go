package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Different web requests in GoLang")
	PerformGetRequest()
	PostJsonRequest()
	PostFormData()
}

func PerformGetRequest() {
	const url = "http://localhost:3000/get"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var response strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteLength, _ := response.Write(content)

	fmt.Println(byteLength)

	fmt.Println("response is:", response.String())

}

func PostJsonRequest() {
	const url = "http://localhost:3000/post"

	body := strings.NewReader(`
	{
		"name":"puneet"
	}
	`)
	res, err := http.Post(url, "application/json", body)

	check(err)
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)

	check(err)

	var response strings.Builder
	byteCount, _ := response.Write(content)

	fmt.Println(byteCount)
	fmt.Println("response: ", response.String())
}

func PostFormData() {
	const myUrl = "http://localhost:3000/postform"

	body := url.Values{}
	body.Add("hello", "world")

	res, err := http.PostForm(myUrl, body)

	check(err)
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)

	check(err)

	var response strings.Builder
	byteCount, _ := response.Write(content)

	fmt.Println(byteCount)
	fmt.Println("response: ", response.String())
}
