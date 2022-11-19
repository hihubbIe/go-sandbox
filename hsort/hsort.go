package hsort

var Quicksort string = "quicksort"

func Sort(sort string, A []int) []int {
	B := make([]int, len(A))
	copy(B, A)
	switch sort {
	case Quicksort:
		_quicksort(B, 0, len(B)-1)
		return B
	default:
		panic("Unknown sort !")
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func _quicksort(A []int, lo int, hi int) {
	if lo >= hi || lo < 0 {
		return
	}
	p := _partition(A, lo, hi)
	_quicksort(A, lo, p-1)
	_quicksort(A, p+1, hi)
	return
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
