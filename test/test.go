package main

import (
	"fmt"
	"functional"
)

func main() {
	A := []int{1, 2, 3}
	B := functional.Map(A, func(i int) string { return fmt.Sprintf("\"%v\"", i) })
	fmt.Println(B)
}
