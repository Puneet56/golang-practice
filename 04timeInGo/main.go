package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime)

	formattedTime := presentTime.Format("01-02-2006 Monday")

	fmt.Println(formattedTime)

	myTime := time.Date(2022, time.January, 25, 4, 4, 1, 0, time.Local)

	fmt.Println(myTime)

}
