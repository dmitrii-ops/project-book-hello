package main

import "fmt"

func hello() string {
	return "Hello"
}

func main() {
	println(fmt.Sprintf(hello()))
	// fmt.Println("Hello")
}
