package adventofcode

import "strconv"

func Map[I any, O any](vs []I, f func(I) O) []O {
	vsm := make([]O, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func parseInt(val string) int {
	if number, err := strconv.ParseInt(val, 10, 64); err == nil {
		return int(number)
	}
	return 0
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	result := 1
	for _, i := range integers {
		result = result * i / GCD(result, i)
	}
	return result
}

func AllEqual[T comparable](arr []T) bool {
	for _, v := range arr {
		if v != arr[0] {
			return false
		}
	}
	return true
}
