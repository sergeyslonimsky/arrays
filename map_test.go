package arrays_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sergeyslonimsky/arrays"
)

func TestMapWalk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      map[string]int
		callback func(string, int) string
		want     []string
	}{
		{
			name: "convert map to slice of strings",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			callback: func(k string, v int) string {
				return fmt.Sprintf("%s:%d", k, v)
			},
			want: []string{"a:1", "b:2", "c:3"},
		},
		{
			name: "transform values only",
			arr:  map[string]int{"x": 10, "y": 20},
			callback: func(k string, v int) string {
				return fmt.Sprintf("%d", v*2)
			},
			want: []string{"20", "40"},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			callback: func(k string, v int) string {
				return fmt.Sprintf("%s:%d", k, v)
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapWalk(tt.arr, tt.callback)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			// Sort both slices since map iteration order is not guaranteed
			sort.Strings(got)
			sort.Strings(tt.want)

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestMapForEach(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  map[string]int
		want []string
	}{
		{
			name: "iterate through map",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			want: []string{"a:1", "b:2", "c:3"},
		},
		{
			name: "single element map",
			arr:  map[string]int{"x": 10},
			want: []string{"x:10"},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := []string{}
			arrays.MapForEach(tt.arr, func(k string, v int) {
				got = append(got, fmt.Sprintf("%s:%d", k, v))
			})

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			// Sort both slices since map iteration order is not guaranteed
			sort.Strings(got)
			sort.Strings(tt.want)

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestMapFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       map[string]int
		predicate func(string, int) bool
		want      map[string]int
	}{
		{
			name: "filter values greater than 2",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(k string, v int) bool {
				return v > 2
			},
			want: map[string]int{"c": 3, "d": 4},
		},
		{
			name: "filter by key",
			arr:  map[string]int{"apple": 1, "banana": 2, "apricot": 3},
			predicate: func(k string, v int) bool {
				return k[0] == 'a'
			},
			want: map[string]int{"apple": 1, "apricot": 3},
		},
		{
			name: "no matches",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool {
				return v > 10
			},
			want: map[string]int{},
		},
		{
			name: "all matches",
			arr:  map[string]int{"a": 1, "b": 2},
			predicate: func(k string, v int) bool {
				return true
			},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			predicate: func(k string, v int) bool {
				return true
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapFilter(tt.arr, tt.predicate)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %s not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %s: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  map[string]int
		want []string
	}{
		{
			name: "get all keys",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			want: []string{"a", "b", "c"},
		},
		{
			name: "single key",
			arr:  map[string]int{"x": 10},
			want: []string{"x"},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			want: []string{},
		},
		{
			name: "numeric keys",
			arr:  map[string]int{"1": 100, "2": 200, "3": 300},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapKeys(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			// Sort both slices since map iteration order is not guaranteed
			sort.Strings(got)
			sort.Strings(tt.want)

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestMapInvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  map[string]int
		want map[int]string
	}{
		{
			name: "swap keys and values",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			want: map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			name: "single element",
			arr:  map[string]int{"x": 10},
			want: map[int]string{10: "x"},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			want: map[int]string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapInvert(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %d not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %d: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      map[string]int
		callback func(string, int) (string, int)
		want     map[string]int
	}{
		{
			name: "double all values",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			callback: func(k string, v int) (string, int) {
				return k, v * 2
			},
			want: map[string]int{"a": 2, "b": 4, "c": 6},
		},
		{
			name: "uppercase keys",
			arr:  map[string]int{"a": 1, "b": 2},
			callback: func(k string, v int) (string, int) {
				return k + k, v
			},
			want: map[string]int{"aa": 1, "bb": 2},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			callback: func(k string, v int) (string, int) {
				return k, v
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapMap(tt.arr, tt.callback)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %s not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %s: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapReduce(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      map[string]int
		initial  int
		callback func(int, string, int) int
		want     int
	}{
		{
			name:    "sum all values",
			arr:     map[string]int{"a": 1, "b": 2, "c": 3},
			initial: 0,
			callback: func(acc int, k string, v int) int {
				return acc + v
			},
			want: 6,
		},
		{
			name:    "sum with initial value",
			arr:     map[string]int{"a": 1, "b": 2},
			initial: 10,
			callback: func(acc int, k string, v int) int {
				return acc + v
			},
			want: 13,
		},
		{
			name:    "empty map returns initial",
			arr:     map[string]int{},
			initial: 42,
			callback: func(acc int, k string, v int) int {
				return acc + v
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapReduce(tt.arr, tt.initial, tt.callback)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapMerge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		maps []map[string]int
		want map[string]int
	}{
		{
			name: "merge two non-overlapping maps",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"c": 3, "d": 4},
			},
			want: map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
		},
		{
			name: "last map wins on conflict",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 99, "c": 3},
			},
			want: map[string]int{"a": 1, "b": 99, "c": 3},
		},
		{
			name: "merge three maps",
			maps: []map[string]int{
				{"a": 1},
				{"b": 2},
				{"c": 3},
			},
			want: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "empty maps",
			maps: []map[string]int{{}, {}},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapMerge(tt.maps...)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %s not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %s: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapMergeWith(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		maps    []map[string]int
		resolve func(int, int) int
		want    map[string]int
	}{
		{
			name: "sum on conflict",
			maps: []map[string]int{
				{"a": 1, "b": 2},
				{"b": 3, "c": 4},
			},
			resolve: func(existing, incoming int) int {
				return existing + incoming
			},
			want: map[string]int{"a": 1, "b": 5, "c": 4},
		},
		{
			name: "keep max on conflict",
			maps: []map[string]int{
				{"a": 10, "b": 2},
				{"a": 5, "b": 8},
			},
			resolve: func(existing, incoming int) int {
				if existing > incoming {
					return existing
				}
				return incoming
			},
			want: map[string]int{"a": 10, "b": 8},
		},
		{
			name: "no conflicts",
			maps: []map[string]int{
				{"a": 1},
				{"b": 2},
			},
			resolve: func(existing, incoming int) int {
				return existing + incoming
			},
			want: map[string]int{"a": 1, "b": 2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapMergeWith(tt.resolve, tt.maps...)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %s not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %s: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapPick(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  map[string]int
		keys []string
		want map[string]int
	}{
		{
			name: "pick existing keys",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			keys: []string{"a", "c"},
			want: map[string]int{"a": 1, "c": 3},
		},
		{
			name: "ignore missing keys",
			arr:  map[string]int{"a": 1, "b": 2},
			keys: []string{"a", "z"},
			want: map[string]int{"a": 1},
		},
		{
			name: "empty keys",
			arr:  map[string]int{"a": 1, "b": 2},
			keys: []string{},
			want: map[string]int{},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			keys: []string{"a"},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapPick(tt.arr, tt.keys)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %s not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %s: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapOmit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  map[string]int
		keys []string
		want map[string]int
	}{
		{
			name: "omit existing keys",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			keys: []string{"a", "c"},
			want: map[string]int{"b": 2},
		},
		{
			name: "omit missing keys changes nothing",
			arr:  map[string]int{"a": 1, "b": 2},
			keys: []string{"z"},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "omit all keys",
			arr:  map[string]int{"a": 1, "b": 2},
			keys: []string{"a", "b"},
			want: map[string]int{},
		},
		{
			name: "empty keys",
			arr:  map[string]int{"a": 1},
			keys: []string{},
			want: map[string]int{"a": 1},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapOmit(tt.arr, tt.keys)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %s not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %s: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapGroupBy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		slice []string
		keyFn func(string) int
		want  map[int][]string
	}{
		{
			name:  "group by string length",
			slice: []string{"a", "bb", "cc", "ddd"},
			keyFn: func(s string) int { return len(s) },
			want:  map[int][]string{1: {"a"}, 2: {"bb", "cc"}, 3: {"ddd"}},
		},
		{
			name:  "all in one group",
			slice: []string{"a", "b", "c"},
			keyFn: func(s string) int { return 0 },
			want:  map[int][]string{0: {"a", "b", "c"}},
		},
		{
			name:  "empty slice",
			slice: []string{},
			keyFn: func(s string) int { return len(s) },
			want:  map[int][]string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapGroupBy(tt.slice, tt.keyFn)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, wantGroup := range tt.want {
				gotGroup, ok := got[k]
				if !ok {
					t.Errorf("expected key %d not found in result", k)
					continue
				}
				if len(gotGroup) != len(wantGroup) {
					t.Errorf("for key %d: got group length %d, want %d", k, len(gotGroup), len(wantGroup))
					continue
				}
				sort.Strings(gotGroup)
				sort.Strings(wantGroup)
				for i, v := range wantGroup {
					if gotGroup[i] != v {
						t.Errorf("for key %d index %d: got %v, want %v", k, i, gotGroup[i], v)
					}
				}
			}
		})
	}
}

func TestMapFromSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		slice []string
		keyFn func(string) int
		want  map[int]string
	}{
		{
			name:  "index by string length",
			slice: []string{"a", "bb", "ccc"},
			keyFn: func(s string) int { return len(s) },
			want:  map[int]string{1: "a", 2: "bb", 3: "ccc"},
		},
		{
			name:  "single element",
			slice: []string{"hello"},
			keyFn: func(s string) int { return len(s) },
			want:  map[int]string{5: "hello"},
		},
		{
			name:  "empty slice",
			slice: []string{},
			keyFn: func(s string) int { return len(s) },
			want:  map[int]string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapFromSlice(tt.slice, tt.keyFn)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			for k, v := range tt.want {
				if gotValue, ok := got[k]; !ok {
					t.Errorf("expected key %d not found in result", k)
				} else if gotValue != v {
					t.Errorf("for key %d: got %v, want %v", k, gotValue, v)
				}
			}
		})
	}
}

func TestMapAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       map[string]int
		predicate func(string, int) bool
		want      bool
	}{
		{
			name:      "at least one matches",
			arr:       map[string]int{"a": 1, "b": 20, "c": 3},
			predicate: func(k string, v int) bool { return v > 10 },
			want:      true,
		},
		{
			name:      "none match",
			arr:       map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return v > 10 },
			want:      false,
		},
		{
			name:      "empty map returns false",
			arr:       map[string]int{},
			predicate: func(k string, v int) bool { return true },
			want:      false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapAny(tt.arr, tt.predicate)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       map[string]int
		predicate func(string, int) bool
		want      bool
	}{
		{
			name:      "all match",
			arr:       map[string]int{"a": 1, "b": 2, "c": 3},
			predicate: func(k string, v int) bool { return v > 0 },
			want:      true,
		},
		{
			name:      "one does not match",
			arr:       map[string]int{"a": 1, "b": -1, "c": 3},
			predicate: func(k string, v int) bool { return v > 0 },
			want:      false,
		},
		{
			name:      "empty map returns true",
			arr:       map[string]int{},
			predicate: func(k string, v int) bool { return false },
			want:      true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapAll(tt.arr, tt.predicate)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  map[string]int
		want []int
	}{
		{
			name: "get all values",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			want: []int{1, 2, 3},
		},
		{
			name: "single value",
			arr:  map[string]int{"x": 10},
			want: []int{10},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			want: []int{},
		},
		{
			name: "duplicate values",
			arr:  map[string]int{"a": 5, "b": 5, "c": 10},
			want: []int{5, 5, 10},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapValues(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			// Sort both slices since map iteration order is not guaranteed
			sort.Ints(got)
			sort.Ints(tt.want)

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}
