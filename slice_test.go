package arrays_test

import (
	"fmt"
	"testing"

	"github.com/sergeyslonimsky/arrays"
)

func TestArrayMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []int
		callback func(int, int) string
		want     []string
	}{
		{
			name: "converting int to string with keys",
			arr:  []int{1, 2, 3, 4, 5},
			callback: func(i, v int) string {
				return fmt.Sprintf("%d%d", i, v)
			},
			want: []string{"01", "12", "23", "34", "45"},
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(i, v int) string {
				return fmt.Sprintf("%d", v)
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayMap(tt.arr, tt.callback)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestArrayMapErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		callback  func(int, int) (string, error)
		want      []string
		wantError bool
	}{
		{
			name: "successful conversion",
			arr:  []int{1, 2, 3},
			callback: func(i, v int) (string, error) {
				return fmt.Sprintf("%d%d", i, v), nil
			},
			want:      []string{"01", "12", "23"},
			wantError: false,
		},
		{
			name: "callback returns error",
			arr:  []int{1, 2, 3},
			callback: func(i, v int) (string, error) {
				if v == 2 {
					return "", fmt.Errorf("error at value %d", v)
				}
				return fmt.Sprintf("%d", v), nil
			},
			want:      nil,
			wantError: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(i, v int) (string, error) {
				return fmt.Sprintf("%d", v), nil
			},
			want:      []string{},
			wantError: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := arrays.ArrayMapErr(tt.arr, tt.callback)

			if (err != nil) != tt.wantError {
				t.Errorf("ArrayMapErr() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				if len(got) != len(tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}

				for i, v := range tt.want {
					if got[i] != v {
						t.Errorf("got %v, want %v", got[i], v)
					}
				}
			}
		})
	}
}

func TestArrayForEach(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []string
	}{
		{
			name: "iterate through array",
			arr:  []int{1, 2, 3},
			want: []string{"0:1", "1:2", "2:3"},
		},
		{
			name: "empty array",
			arr:  []int{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := []string{}
			arrays.ArrayForEach(tt.arr, func(i, v int) {
				got = append(got, fmt.Sprintf("%d:%d", i, v))
			})

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestArrayFilter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		predicate func(int, int) bool
		want      []int
	}{
		{
			name: "filter even numbers",
			arr:  []int{1, 2, 3, 4, 5, 6},
			predicate: func(i, v int) bool {
				return v%2 == 0
			},
			want: []int{2, 4, 6},
		},
		{
			name: "filter by index",
			arr:  []int{10, 20, 30, 40},
			predicate: func(i, v int) bool {
				return i > 1
			},
			want: []int{30, 40},
		},
		{
			name: "no matches",
			arr:  []int{1, 3, 5},
			predicate: func(i, v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
		{
			name: "empty array",
			arr:  []int{},
			predicate: func(i, v int) bool {
				return true
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayFilter(tt.arr, tt.predicate)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestArrayConcat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arrs [][]int
		want []int
	}{
		{
			name: "concatenate two arrays",
			arrs: [][]int{{1, 2}, {3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "concatenate multiple arrays",
			arrs: [][]int{{1}, {2, 3}, {4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "concatenate with empty arrays",
			arrs: [][]int{{1, 2}, {}, {3}},
			want: []int{1, 2, 3},
		},
		{
			name: "single array",
			arrs: [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayConcat(tt.arrs...)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestArrayEvery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		predicate func(int) bool
		want      bool
	}{
		{
			name: "all elements match",
			arr:  []int{2, 4, 6, 8},
			predicate: func(v int) bool {
				return v%2 == 0
			},
			want: true,
		},
		{
			name: "not all elements match",
			arr:  []int{2, 3, 6, 8},
			predicate: func(v int) bool {
				return v%2 == 0
			},
			want: false,
		},
		{
			name: "empty array",
			arr:  []int{},
			predicate: func(v int) bool {
				return v > 10
			},
			want: true,
		},
		{
			name: "single element matching",
			arr:  []int{5},
			predicate: func(v int) bool {
				return v > 0
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayEvery(tt.arr, tt.predicate)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUniq(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "remove duplicates",
			arr:  []int{1, 2, 2, 3, 3, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "no duplicates",
			arr:  []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "all duplicates",
			arr:  []int{5, 5, 5, 5},
			want: []int{5},
		},
		{
			name: "empty array",
			arr:  []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayUniq(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			// Create a map for comparison since order may vary
			gotMap := make(map[int]bool)
			for _, v := range got {
				gotMap[v] = true
			}

			for _, v := range tt.want {
				if !gotMap[v] {
					t.Errorf("expected value %d not found in result", v)
				}
			}
		})
	}
}

func TestArrayHashUniq(t *testing.T) {
	t.Parallel()

	type person struct {
		name string
		age  int
	}

	tests := []struct {
		name     string
		arr      []person
		hashFunc func(person) string
		want     []person
	}{
		{
			name: "unique by name",
			arr: []person{
				{name: "Alice", age: 30},
				{name: "Bob", age: 25},
				{name: "Alice", age: 35},
			},
			hashFunc: func(p person) string {
				return p.name
			},
			want: []person{
				{name: "Alice", age: 30}, // first Alice wins
				{name: "Bob", age: 25},
			},
		},
		{
			name:     "empty array",
			arr:      []person{},
			hashFunc: func(p person) string { return p.name },
			want:     []person{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayHashUniq(tt.arr, tt.hashFunc)

			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want length %d", len(got), len(tt.want))
			}

			// Create a map for comparison since order may vary
			gotMap := make(map[string]person)
			for _, v := range got {
				gotMap[tt.hashFunc(v)] = v
			}

			for _, v := range tt.want {
				key := tt.hashFunc(v)
				if _, ok := gotMap[key]; !ok {
					t.Errorf("expected key %s not found in result", key)
				}
			}
		})
	}
}

func TestArrayFind(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		predicate func(int, int) bool
		wantValue int
		wantFound bool
	}{
		{
			name: "find first matching element",
			arr:  []int{1, 2, 3, 4, 5},
			predicate: func(i, v int) bool {
				return v > 2
			},
			wantValue: 3,
			wantFound: true,
		},
		{
			name: "no matching element",
			arr:  []int{1, 2, 3},
			predicate: func(i, v int) bool {
				return v > 10
			},
			wantValue: 0,
			wantFound: false,
		},
		{
			name: "find by index",
			arr:  []int{10, 20, 30},
			predicate: func(i, v int) bool {
				return i == 1
			},
			wantValue: 20,
			wantFound: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			predicate: func(i, v int) bool {
				return true
			},
			wantValue: 0,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotValue, gotFound := arrays.ArrayFind(tt.arr, tt.predicate)

			if gotFound != tt.wantFound {
				t.Errorf("got found %v, want found %v", gotFound, tt.wantFound)
			}

			if gotValue != tt.wantValue {
				t.Errorf("got value %v, want value %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestArrayFindIndex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		predicate func(int, int) bool
		wantIndex int
		wantFound bool
	}{
		{
			name: "find first matching index",
			arr:  []int{1, 2, 3, 4, 5},
			predicate: func(i, v int) bool {
				return v > 2
			},
			wantIndex: 2,
			wantFound: true,
		},
		{
			name: "no matching element",
			arr:  []int{1, 2, 3},
			predicate: func(i, v int) bool {
				return v > 10
			},
			wantIndex: -1,
			wantFound: false,
		},
		{
			name: "find first element",
			arr:  []int{10, 20, 30},
			predicate: func(i, v int) bool {
				return v == 10
			},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			predicate: func(i, v int) bool {
				return true
			},
			wantIndex: -1,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotIndex, gotFound := arrays.ArrayFindIndex(tt.arr, tt.predicate)

			if gotFound != tt.wantFound {
				t.Errorf("got found %v, want found %v", gotFound, tt.wantFound)
			}

			if gotIndex != tt.wantIndex {
				t.Errorf("got index %v, want index %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestArrayReverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "reverse array",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "reverse single element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "reverse two elements",
			arr:  []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "empty array",
			arr:  []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayReverse(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestArrayContains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		elem int
		want bool
	}{
		{
			name: "element exists",
			arr:  []int{1, 2, 3, 4, 5},
			elem: 3,
			want: true,
		},
		{
			name: "element does not exist",
			arr:  []int{1, 2, 3, 4, 5},
			elem: 6,
			want: false,
		},
		{
			name: "first element",
			arr:  []int{1, 2, 3},
			elem: 1,
			want: true,
		},
		{
			name: "last element",
			arr:  []int{1, 2, 3},
			elem: 3,
			want: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			elem: 1,
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayContains(tt.arr, tt.elem)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayProcess(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []int
		callback func(int) string
		want     []string
	}{
		{
			name: "convert int to string",
			arr:  []int{1, 2, 3, 4, 5},
			callback: func(v int) string {
				return fmt.Sprintf("num-%d", v)
			},
			want: []string{"num-1", "num-2", "num-3", "num-4", "num-5"},
		},
		{
			name: "double values",
			arr:  []int{1, 2, 3},
			callback: func(v int) string {
				return fmt.Sprintf("%d", v*2)
			},
			want: []string{"2", "4", "6"},
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(v int) string {
				return fmt.Sprintf("%d", v)
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayProcess(tt.arr, tt.callback)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}

func TestArrayAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		predicate func(int) bool
		want      bool
	}{
		{
			name: "at least one matches",
			arr:  []int{1, 2, 3, 4},
			predicate: func(v int) bool {
				return v > 3
			},
			want: true,
		},
		{
			name: "none match",
			arr:  []int{1, 2, 3},
			predicate: func(v int) bool {
				return v > 10
			},
			want: false,
		},
		{
			name: "empty array returns false",
			arr:  []int{},
			predicate: func(v int) bool {
				return true
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayAny(tt.arr, tt.predicate)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayReduce(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []int
		initial  int
		callback func(int, int) int
		want     int
	}{
		{
			name:    "sum all elements",
			arr:     []int{1, 2, 3, 4, 5},
			initial: 0,
			callback: func(acc, v int) int {
				return acc + v
			},
			want: 15,
		},
		{
			name:    "sum with initial value",
			arr:     []int{1, 2, 3},
			initial: 10,
			callback: func(acc, v int) int {
				return acc + v
			},
			want: 16,
		},
		{
			name:    "empty array returns initial",
			arr:     []int{},
			initial: 42,
			callback: func(acc, v int) int {
				return acc + v
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayReduce(tt.arr, tt.initial, tt.callback)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayFlatten(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  [][]int
		want []int
	}{
		{
			name: "flatten nested slices",
			arr:  [][]int{{1, 2}, {3, 4}, {5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "flatten with empty inner slices",
			arr:  [][]int{{1}, {}, {2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "single nested slice",
			arr:  [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "empty outer slice",
			arr:  [][]int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayFlatten(tt.arr)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("index %d: got %v, want %v", i, got[i], v)
				}
			}
		})
	}
}

func TestArrayChunk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arr  []int
		size int
		want [][]int
	}{
		{
			name: "evenly divisible",
			arr:  []int{1, 2, 3, 4, 5, 6},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "last chunk smaller",
			arr:  []int{1, 2, 3, 4, 5},
			size: 2,
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "chunk size equals length",
			arr:  []int{1, 2, 3},
			size: 3,
			want: [][]int{{1, 2, 3}},
		},
		{
			name: "chunk size larger than length",
			arr:  []int{1, 2},
			size: 5,
			want: [][]int{{1, 2}},
		},
		{
			name: "empty array",
			arr:  []int{},
			size: 3,
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayChunk(tt.arr, tt.size)

			if len(got) != len(tt.want) {
				t.Errorf("got %v chunks, want %v chunks", len(got), len(tt.want))
				return
			}

			for i, wantChunk := range tt.want {
				if len(got[i]) != len(wantChunk) {
					t.Errorf("chunk %d: got length %d, want %d", i, len(got[i]), len(wantChunk))
					continue
				}
				for j, v := range wantChunk {
					if got[i][j] != v {
						t.Errorf("chunk %d index %d: got %v, want %v", i, j, got[i][j], v)
					}
				}
			}
		})
	}
}

func TestArrayDifference(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "elements in a but not b",
			a:    []int{1, 2, 3, 4, 5},
			b:    []int{3, 4},
			want: []int{1, 2, 5},
		},
		{
			name: "no difference",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: []int{},
		},
		{
			name: "b is empty",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "a is empty",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayDifference(tt.a, tt.b)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
				return
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("index %d: got %v, want %v", i, got[i], v)
				}
			}
		})
	}
}

func TestArrayIntersection(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "common elements",
			a:    []int{1, 2, 3, 4, 5},
			b:    []int{3, 4, 6},
			want: []int{3, 4},
		},
		{
			name: "no common elements",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: []int{},
		},
		{
			name: "all common",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "a is empty",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayIntersection(tt.a, tt.b)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
				return
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("index %d: got %v, want %v", i, got[i], v)
				}
			}
		})
	}
}

func TestArrayProcessErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		arr       []int
		callback  func(int) (string, error)
		want      []string
		wantError bool
	}{
		{
			name: "successful processing",
			arr:  []int{1, 2, 3},
			callback: func(v int) (string, error) {
				return fmt.Sprintf("num-%d", v), nil
			},
			want:      []string{"num-1", "num-2", "num-3"},
			wantError: false,
		},
		{
			name: "callback returns error",
			arr:  []int{1, 2, 3},
			callback: func(v int) (string, error) {
				if v == 2 {
					return "", fmt.Errorf("error at value %d", v)
				}
				return fmt.Sprintf("%d", v), nil
			},
			want:      nil,
			wantError: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(v int) (string, error) {
				return fmt.Sprintf("%d", v), nil
			},
			want:      []string{},
			wantError: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := arrays.ArrayProcessErr(tt.arr, tt.callback)

			if (err != nil) != tt.wantError {
				t.Errorf("ArrayProcessErr() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if !tt.wantError {
				if len(got) != len(tt.want) {
					t.Errorf("got %v, want %v", got, tt.want)
				}

				for i, v := range tt.want {
					if got[i] != v {
						t.Errorf("got %v, want %v", got[i], v)
					}
				}
			}
		})
	}
}
