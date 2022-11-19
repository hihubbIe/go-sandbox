package functional

func Map[T any, R any](array []T, fun func(T) R) []R {
	res := make([]R, len(array))
	for i := 0; i < len(array); i++ {
		res[i] = fun(array[i])
	}
	return res
}
