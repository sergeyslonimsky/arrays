package arrays

func MapWalk[I comparable, K, T any](arr map[I]K, callback func(key I, value K) T) []T {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		r = append(r, callback(i, v))
	}

	return r
}

func MapForEach[I comparable, K any](arr map[I]K, callback func(key I, value K)) {
	for i, v := range arr {
		callback(i, v)
	}
}

func MapFilter[I comparable, K any](arr map[I]K, callback func(key I, value K) bool) map[I]K {
	r := make(map[I]K, len(arr))

	for i, v := range arr {
		if callback(i, v) {
			r[i] = v
		}
	}

	return r
}

func MapKeys[I comparable, K any](arr map[I]K) []I {
	r := make([]I, 0, len(arr))

	for k := range arr {
		r = append(r, k)
	}

	return r
}

func MapValues[I comparable, K any](arr map[I]K) []K {
	r := make([]K, 0, len(arr))

	for _, k := range arr {
		r = append(r, k)
	}

	return r
}
