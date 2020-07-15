package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	f64 := float64(i)
	fmt.Printf("%T %v %f\n", f64, f64, f64)

	var f64b float64 = 98.5
	ib := f64b
	fmt.Printf("%T %v %f\n", ib, ib, ib)

	var s string = "99"
	ic, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d\n", ic, ic, ic)
}
