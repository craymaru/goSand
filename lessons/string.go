package main

import (
	"fmt"
	"strings"
)

func main()  {
	var s string = "Hello"
	fmt.Println(s[0])
	fmt.Println(string(72))

	s = strings.Replace(s, "o", "o!", 1)
	fmt.Println(strings.Contains(s, "!"))

	fmt.Println(`Test
Something Strings`)
}
