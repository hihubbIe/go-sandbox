package hsort

var Quicksort string = "quicksort"
var Bubblesort string = "bubblesort"

func Sort(sort string, A []int) (B []int) {
	B = make([]int, len(A))
	copy(B, A)
	switch sort {
	case Quicksort:
		_quicksort(B, 0, len(B)-1)
		return
	case Bubblesort:
		_bubblesort(B)
		return
	default:
		panic("Unknown sort !")
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func _bubblesort(A []int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)-i-1; j++ {
			if A[j] > A[j+1] {
				swap(&A[j], &A[j+1])
			}
		}
	}
}

func _quicksort(A []int, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := _partition(A, lo, hi)
	_quicksort(A, lo, p-1)
	_quicksort(A, p+1, hi)
}

func _partition(A []int, lo int, hi int) int {
	pivot := A[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if A[j] <= pivot {
			i++
			swap(&A[i], &A[j])
		}
	}
	i++
	swap(&A[i], &A[hi])
	return i
}
