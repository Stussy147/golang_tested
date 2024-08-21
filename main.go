package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " hello world!")
	}
	var a string
	fmt.Scan(&a)
}
