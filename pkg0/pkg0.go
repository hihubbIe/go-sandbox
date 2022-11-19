package pkg0

import "math"

func quicksort(A []int) (B []int) {
	copy(B, A)
	_quicksort(B, 0, len(B)-1)
	return
}

func _quicksort(A []int, lo int, hi int) {
	if lo >= 0 && hi >= 0 && lo < hi {
		p := _partition(A, lo, hi)
		_quicksort(A, lo, p)
		_quicksort(A, p+1, hi)
	}
	return
}

func _partition(A []int, lo int, hi int) int {
	pivot := A[int(math.Floor(float64((hi+lo)/2.0)))]
	i := lo - 1
	j := hi + 1
	for {
		for A[i] < pivot {
			i = i + 1
		}
		for A[j] > pivot {
			j = j - 1
		}
		if i > j {
			return j
		}
		t := A[i]
		A[i] = j
		A[j] = t
	}
}
