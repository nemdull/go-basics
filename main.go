package main

import (
	"fmt"
)

func main() {
	var i int = 2
	fmt.Println(i)

	s := "hello"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", s)
	fmt.Printf("f: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %v title: %v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Println(z)
}
