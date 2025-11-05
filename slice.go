package arrays

import "fmt"

// ArrayMap creates a new array populated with the results of calling a provided function
// on every element in the calling array.
func ArrayMap[I, T any](arr []I, callback func(key int, value I) T) []T {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		r = append(r, callback(i, v))
	}

	return r
}

// ArrayMapErr creates a new array populated with the results of calling a provided function
// on every element in the calling array.
// Returns first error, if callback fails.
func ArrayMapErr[I, T any](arr []I, callback func(key int, value I) (T, error)) ([]T, error) {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		res, err := callback(i, v)
		if err != nil {
			return nil, fmt.Errorf("callback: %w", err)
		}

		r = append(r, res)
	}

	return r, nil
}

// ArrayForEach executes a provided function once for each array element.
func ArrayForEach[I any](arr []I, callback func(key int, value I)) {
	for i, v := range arr {
		callback(i, v)
	}
}

// ArrayFilter creates a copy of a given array,
// filtered down to just the elements from the given array that pass the test implemented by the provided function.
func ArrayFilter[I any](arr []I, callback func(key int, value I) bool) []I {
	r := make([]I, 0, len(arr))

	for i, v := range arr {
		if callback(i, v) {
			r = append(r, v)
		}
	}

	return r
}

// ArrayConcat creates a new array with all the elements of the provided arrays.
func ArrayConcat[I any](arr ...[]I) []I {
	r := make([]I, 0, len(arr[0]))

	for _, arrN := range arr {
		r = append(r, arrN...)
	}

	return r
}

// ArrayEvery tests whether all elements in the array pass the test implemented by the provided function.
func ArrayEvery[I any](arr []I, callback func(value I) bool) bool {
	if len(arr) == 0 {
		return true
	}

	for _, v := range arr {
		if !callback(v) {
			return false
		}
	}

	return true
}

// ArrayUniq creates a new array with all unique values from provided array.
// Provided array should contain comparable values.
// For non-comparable values use ArrayHashUniq.
func ArrayUniq[I comparable](arr []I) []I {
	r := make([]I, 0, len(arr))
	uniqMap := make(map[I]bool, len(arr))

	for _, v := range arr {
		uniqMap[v] = true
	}

	for v := range uniqMap {
		r = append(r, v)
	}

	return r
}

// ArrayHashUniq creates a new array with all unique values comparable by provided hashFunc.
// HashFunc should return unique string for every unique element.
func ArrayHashUniq[I any](arr []I, hashFunc func(value I) string) []I {
	r := make([]I, 0, len(arr))
	uniqMap := make(map[string]I, len(arr))

	for _, v := range arr {
		uniqMap[hashFunc(v)] = v
	}

	for _, v := range uniqMap {
		r = append(r, v)
	}

	return r
}

// ArrayFind returns the first element in the provided array that satisfies the provided testing function.
// If no values satisfy the testing function, empty value and false are returned.
func ArrayFind[I any](arr []I, callback func(key int, value I) bool) (I, bool) {
	for i, v := range arr {
		if callback(i, v) {
			return v, true
		}
	}

	return *new(I), false
}

// ArrayFindIndex returns the index of the first element in the provided array
// that satisfies the provided testing function.
// If no values satisfy the testing function, -1 and false are returned.
func ArrayFindIndex[I any](arr []I, callback func(key int, value I) bool) (int, bool) {
	for i, v := range arr {
		if callback(i, v) {
			return i, true
		}
	}

	return -1, false
}

// ArrayReverse reverses an array in place and returns the new array.
// The first array element now becoming the last, and the last array element becoming the first.
func ArrayReverse[I any](arr []I) []I {
	res := make([]I, len(arr))
	copy(res, arr)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

// ArrayContains returns bool if the element in the array, otherwise returns false.
func ArrayContains[I comparable](arr []I, elem I) bool {
	for _, i := range arr {
		if i == elem {
			return true
		}
	}

	return false
}

// ArrayProcess creates a new array populated with the results of calling a provided function
// on every element in the calling array.
func ArrayProcess[I, T any](arr []I, callback func(value I) T) []T {
	r := make([]T, 0, len(arr))

	for _, v := range arr {
		r = append(r, callback(v))
	}

	return r
}

// ArrayProcessErr creates a new array populated with the results of calling a provided function
// on every element in the calling array.
// Returns first error, if callback fails.
func ArrayProcessErr[I, T any](arr []I, callback func(value I) (T, error)) ([]T, error) {
	r := make([]T, 0, len(arr))

	for _, v := range arr {
		res, err := callback(v)
		if err != nil {
			return nil, fmt.Errorf("callback: %w", err)
		}

		r = append(r, res)
	}

	return r, nil
}
