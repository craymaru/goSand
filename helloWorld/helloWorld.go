package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init.")
	fmt.Println(time.Now())
	fmt.Println(user.Current())
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