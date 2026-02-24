package arrays

// MapWalk transforms each key-value pair of m by applying callback and returns
// the results as a slice. The order of elements in the returned slice is not guaranteed.
func MapWalk[I comparable, K, T any](arr map[I]K, callback func(key I, value K) T) []T {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		r = append(r, callback(i, v))
	}

	return r
}

// MapForEach calls callback for each key-value pair in m.
// It is intended for side effects; use MapWalk to collect results.
func MapForEach[I comparable, K any](arr map[I]K, callback func(key I, value K)) {
	for i, v := range arr {
		callback(i, v)
	}
}

// MapFilter returns a new map containing only the key-value pairs for which
// callback returns true.
func MapFilter[I comparable, K any](arr map[I]K, predicate func(key I, value K) bool) map[I]K {
	r := make(map[I]K, len(arr))

	for i, v := range arr {
		if predicate(i, v) {
			r[i] = v
		}
	}

	return r
}

// MapKeys returns a slice of all keys in m.
// The order of keys is not guaranteed.
func MapKeys[I comparable, K any](arr map[I]K) []I {
	r := make([]I, 0, len(arr))

	for k := range arr {
		r = append(r, k)
	}

	return r
}

// MapValues returns a slice of all values in m.
// The order of values is not guaranteed.
func MapValues[I comparable, K any](arr map[I]K) []K {
	r := make([]K, 0, len(arr))

	for _, k := range arr {
		r = append(r, k)
	}

	return r
}

// MapInvert returns a new map with keys and values swapped.
// If multiple keys in m map to the same value, only one will appear in the result;
// which one is non-deterministic due to map iteration order.
func MapInvert[K, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}

	return result
}

// MapMap transforms m by applying callback to each key-value pair and returns
// a new map built from the returned key-value pairs.
// If multiple pairs produce the same key, the last one wins.
func MapMap[K comparable, V any, K2 comparable, V2 any](m map[K]V, callback func(K, V) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))
	for k, v := range m {
		k2, v2 := callback(k, v)
		result[k2] = v2
	}

	return result
}

// MapReduce reduces m to a single value by applying callback to each key-value pair.
// initial is used as the starting accumulator value.
func MapReduce[K comparable, V, R any](m map[K]V, initial R, callback func(acc R, key K, value V) R) R {
	acc := initial
	for k, v := range m {
		acc = callback(acc, k, v)
	}

	return acc
}

// MapMerge merges multiple maps into a single map.
// If the same key exists in multiple maps, the value from the last map wins.
func MapMerge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// MapMergeWith merges multiple maps using resolve to handle key conflicts.
// When the same key exists in multiple maps, resolve is called with the existing
// and incoming values to determine the final value.
func MapMergeWith[K comparable, V any](resolve func(existing, incoming V) V, maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			if existing, ok := result[k]; ok {
				result[k] = resolve(existing, v)
			} else {
				result[k] = v
			}
		}
	}

	return result
}

// MapPick returns a new map containing only the entries whose keys are in keys.
// Keys not present in m are silently ignored.
func MapPick[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V, len(keys))
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}

	return result
}

// MapOmit returns a new map with all entries from m except those whose keys are in keys.
func MapOmit[K comparable, V any](m map[K]V, keys []K) map[K]V {
	excluded := make(map[K]struct{}, len(keys))
	for _, k := range keys {
		excluded[k] = struct{}{}
	}

	result := make(map[K]V, len(m))
	for k, v := range m {
		if _, skip := excluded[k]; !skip {
			result[k] = v
		}
	}

	return result
}

// MapGroupBy groups slice elements into a map by the key returned by keyFn.
// Elements with the same key are collected into a slice.
func MapGroupBy[K comparable, V any](slice []V, keyFn func(V) K) map[K][]V {
	result := make(map[K][]V)
	for _, v := range slice {
		k := keyFn(v)
		result[k] = append(result[k], v)
	}

	return result
}

// MapFromSlice builds a map from slice using keyFn to derive the key for each element.
// If multiple elements produce the same key, the last one wins.
func MapFromSlice[K comparable, V any](slice []V, keyFn func(V) K) map[K]V {
	result := make(map[K]V, len(slice))
	for _, v := range slice {
		result[keyFn(v)] = v
	}

	return result
}

// MapAny reports whether any key-value pair in m satisfies predicate.
// Returns false for an empty map.
func MapAny[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for k, v := range m {
		if predicate(k, v) {
			return true
		}
	}

	return false
}

// MapAll reports whether all key-value pairs in m satisfy predicate.
// Returns true for an empty map.
func MapAll[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for k, v := range m {
		if !predicate(k, v) {
			return false
		}
	}

	return true
}
