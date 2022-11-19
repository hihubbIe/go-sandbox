package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"hsort"
)

func main() {
	if len(os.Args) < 2 {
		panic("No argument passed")
	}
	sample_size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	trace := false
	if len(os.Args) > 2 && os.Args[2] == "trace" {
		trace = true
	}

	A := make([]int, sample_size)
	for i := 0; i < sample_size; i++ {
		A[i] = rand.Intn(sample_size)
	}
	if trace {
		fmt.Println(A)
	}
	fmt.Println("Initial array built")

	B := hsort.Sort(hsort.Quicksort, A)
	if trace {
		fmt.Println(B)
	}
	fmt.Println("Initial array sorted")
	for i := 0; i < sample_size-1; i++ {
		if B[i] > B[i+1] {
			panic(fmt.Sprintf("Value at index %d(%d) higher than next value at index %d(%d)", i, B[i], i+1, B[i+1]))
		}
	}
	fmt.Println("Array was correctly sorted !")
}
