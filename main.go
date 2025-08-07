package main

import "fmt"

func main() {
	var a1 [3]int
	var a2 = [3]int{10, 20, 30}
	a3 := [...]int{10, 20}
	fmt.Printf("%v %v %v\n", a1, a2, a3)
	fmt.Printf("%v %v\n", len(a3), cap(a3))
	fmt.Printf("%T %T\n", a2, a3)

	var s1 []int
	s2 := []int{}
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
}
