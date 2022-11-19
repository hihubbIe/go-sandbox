package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"hsort"
)

type Params struct {
	size  int
	sort  string
	btime int64
	bunit string
	trace bool
}

func processArgs() Params {
	params := Params{
		size:  1000000,
		sort:  hsort.Quicksort,
		btime: 1000000,
		bunit: "ms",
		trace: false,
	}

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "-n":
			val, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				panic(err)
			}
			params.size = val
			i++
		case "-s":
			val := os.Args[i+1]
			switch val {
			case "quicksort":
				params.sort = hsort.Quicksort
				break
			case "bubblesort":
				params.sort = hsort.Bubblesort
				break
			}
			i++
		case "-trace":
			params.trace = true
			break
		case "-t":
			val := os.Args[i+1]
			switch val {
			case "ns":
				params.btime = 1
				params.bunit = "ns"
				break
			case "mis":
				params.btime = 1000
				params.bunit = "μs"
				break
			case "ms":
				params.btime = 1000000
				params.bunit = "ms"
				break
			case "s":
				params.btime = 1000000000
				params.bunit = "s"
				break
			}
		}
	}
	return params
}

func main() {
	params := processArgs()

	fmt.Printf("Sorting an array of %d elements using %s\n", params.size, params.sort)

	A := make([]int, params.size)
	for i := 0; i < params.size; i++ {
		A[i] = rand.Intn(params.size)
	}

	// Sort array
	startTime := time.Now().UnixNano()
	B := hsort.Sort(params.sort, A)
	endTime := time.Now().UnixNano()
	fmt.Printf("Elapsed time: %v%s\n", (endTime-startTime)/params.btime, params.bunit)

	sorted := true
	for i := 0; i < params.size-1; i++ {
		if B[i] > B[i+1] {
			fmt.Printf("%s error: Value at index %d(%d) higher than next value at index %d(%d)", params.sort, i, B[i], i+1, B[i+1])
			sorted = false
			break
		}
	}
	if sorted {
		fmt.Printf("The array was correctly sorted using %s\n", params.sort)
	}
}
