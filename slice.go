package arrays

import "fmt"

// ArrayMap transforms each element of arr by applying callback and returns the results as a new slice.
// callback receives the element's index and value.
func ArrayMap[I, T any](arr []I, callback func(index int, value I) T) []T {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		r = append(r, callback(i, v))
	}

	return r
}

// ArrayMapErr is like ArrayMap but callback may return an error.
// Returns the first error encountered, wrapped with the element's index for context.
func ArrayMapErr[I, T any](arr []I, callback func(index int, value I) (T, error)) ([]T, error) {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		res, err := callback(i, v)
		if err != nil {
			return nil, fmt.Errorf("index %d: %w", i, err)
		}

		r = append(r, res)
	}

	return r, nil
}

// ArrayProcess transforms each element of arr by applying callback and returns the results as a new slice.
// Unlike ArrayMap, callback receives only the element value without the index,
// which allows passing existing converter functions directly without wrapping them in a closure.
func ArrayProcess[I, T any](arr []I, callback func(value I) T) []T {
	r := make([]T, 0, len(arr))

	for _, v := range arr {
		r = append(r, callback(v))
	}

	return r
}

// ArrayProcessErr is like ArrayProcess but callback may return an error.
// Returns the first error encountered.
// Use this when your converter function returns an error and does not need the element index.
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

// ArrayForEach calls callback for each element of arr.
// callback receives the element's index and value.
// It is intended for side effects; use ArrayMap to collect results.
func ArrayForEach[I any](arr []I, callback func(index int, value I)) {
	for i, v := range arr {
		callback(i, v)
	}
}

// ArrayFilter returns a new slice containing only elements for which predicate returns true.
// predicate receives the element's index and value.
func ArrayFilter[I any](arr []I, predicate func(index int, value I) bool) []I {
	r := make([]I, 0, len(arr))

	for i, v := range arr {
		if predicate(i, v) {
			r = append(r, v)
		}
	}

	return r
}

// ArrayConcat returns a new slice containing all elements from all provided slices, in order.
func ArrayConcat[I any](arrs ...[]I) []I {
	total := 0
	for _, a := range arrs {
		total += len(a)
	}

	r := make([]I, 0, total)
	for _, a := range arrs {
		r = append(r, a...)
	}

	return r
}

// ArrayEvery reports whether all elements in arr satisfy predicate.
// Returns true for an empty slice.
func ArrayEvery[I any](arr []I, predicate func(value I) bool) bool {
	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}

	return true
}

// ArrayAny reports whether any element in arr satisfies predicate.
// Returns false for an empty slice.
func ArrayAny[I any](arr []I, predicate func(value I) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}

	return false
}

// ArrayUniq returns a new slice with duplicate elements removed, preserving the order of first occurrence.
// Elements must be comparable.
// For non-comparable values use ArrayHashUniq.
func ArrayUniq[I comparable](arr []I) []I {
	seen := make(map[I]struct{}, len(arr))
	r := make([]I, 0, len(arr))

	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			r = append(r, v)
		}
	}

	return r
}

// ArrayHashUniq returns a new slice with duplicate elements removed, preserving the order of first occurrence.
// Uniqueness is determined by hashFn, which must return a unique string for each unique element.
// For comparable values prefer ArrayUniq.
func ArrayHashUniq[I any](arr []I, hashFn func(value I) string) []I {
	seen := make(map[string]struct{}, len(arr))
	r := make([]I, 0, len(arr))

	for _, v := range arr {
		h := hashFn(v)
		if _, ok := seen[h]; !ok {
			seen[h] = struct{}{}
			r = append(r, v)
		}
	}

	return r
}

// ArrayFind returns the first element in arr that satisfies predicate, along with true.
// If no element satisfies predicate, returns the zero value and false.
// predicate receives the element's index and value.
func ArrayFind[I any](arr []I, predicate func(index int, value I) bool) (I, bool) {
	for i, v := range arr {
		if predicate(i, v) {
			return v, true
		}
	}

	var zero I
	return zero, false
}

// ArrayFindIndex returns the index of the first element in arr that satisfies predicate, along with true.
// If no element satisfies predicate, returns -1 and false.
// predicate receives the element's index and value.
func ArrayFindIndex[I any](arr []I, predicate func(index int, value I) bool) (int, bool) {
	for i, v := range arr {
		if predicate(i, v) {
			return i, true
		}
	}

	return -1, false
}

// ArrayReverse returns a new slice with elements in reverse order.
// The original slice is not modified.
func ArrayReverse[I any](arr []I) []I {
	r := make([]I, len(arr))
	copy(r, arr)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}

// ArrayContains reports whether elem is present in arr.
func ArrayContains[I comparable](arr []I, elem I) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}

	return false
}

// ArrayReduce reduces arr to a single value by applying callback to each element.
// initial is used as the starting accumulator value.
func ArrayReduce[I, R any](arr []I, initial R, callback func(acc R, value I) R) R {
	acc := initial

	for _, v := range arr {
		acc = callback(acc, v)
	}

	return acc
}

// ArrayFlatten returns a new slice containing all elements from all nested slices, in order.
func ArrayFlatten[I any](arr [][]I) []I {
	return ArrayConcat(arr...)
}

// ArrayChunk splits arr into consecutive sub-slices of the given size.
// The last chunk may be smaller than size if the length of arr is not evenly divisible.
// Panics if size is less than 1.
func ArrayChunk[I any](arr []I, size int) [][]I {
	if size < 1 {
		panic("arrays: chunk size must be greater than 0")
	}

	r := make([][]I, 0, (len(arr)+size-1)/size)

	for len(arr) > 0 {
		end := size
		if end > len(arr) {
			end = len(arr)
		}

		r = append(r, arr[:end])
		arr = arr[end:]
	}

	return r
}

// ArrayDifference returns a new slice containing elements that are present in a but not in b.
// The order of elements follows a.
func ArrayDifference[I comparable](a, b []I) []I {
	exclude := make(map[I]struct{}, len(b))
	for _, v := range b {
		exclude[v] = struct{}{}
	}

	r := make([]I, 0)
	for _, v := range a {
		if _, ok := exclude[v]; !ok {
			r = append(r, v)
		}
	}

	return r
}

// ArrayIntersection returns a new slice containing elements present in both a and b.
// The order of elements follows a.
func ArrayIntersection[I comparable](a, b []I) []I {
	bSet := make(map[I]struct{}, len(b))
	for _, v := range b {
		bSet[v] = struct{}{}
	}

	r := make([]I, 0)
	for _, v := range a {
		if _, ok := bSet[v]; ok {
			r = append(r, v)
		}
	}

	return r
}
