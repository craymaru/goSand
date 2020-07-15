package main

import "fmt"

func init() {
	fmt.Println("Init.")
}

func hello() {
	fmt.Println("Hello world.")
}

func main() {
	hello()
	meow()
}

func meow() {
	fmt.Println("Meow", "Meow")
}