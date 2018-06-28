package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		strNumber     string
		number        int64
		parseIntError error
	)

	fmt.Print("Enter a number if you want: ")
	fmt.Scan(&strNumber)

	number, parseIntError = strconv.ParseInt(strNumber, 10, 64)

	if parseIntError == nil {
		if number%2 == 0 {
			fmt.Println("Number is even!")
		} else {
			fmt.Println("Number is odd!")
		}
	} else {
		fmt.Println("Oops, look's like something went wrong :(")
	}

}
