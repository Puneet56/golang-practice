package main

import "fmt"

func main() {
	fmt.Println("Methods in GoLang")

	puneet := User{"Puneet Dhiman", "puneet@go.dev", true}

	fmt.Println(puneet)

	puneet.GetStatus()

	fmt.Println(puneet)
}

type User struct {
	Name   string
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("User Status:", u.Status)
}
