package main

import "fmt"

func main() {
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	a3 := [...]int{10, 20}
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	fmt.Printf("%v %v\n", len(a3), cap(a3))

}
