package main

import "fmt"

func main() {

	var (
		i int = 1
		f64 float64 = 99.725
		s string = "text"
		b bool = true
	)
	fmt.Println(i, f64, s, b)

	// Short Variable Declarations

	sf64 := 999.98
	fmt.Println(sf64)
	fmt.Printf("%T\n", sf64)

}
