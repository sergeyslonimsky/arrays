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
		name     string
		arr      []int
		callback func(int, int) bool
		want     []int
	}{
		{
			name: "filter even numbers",
			arr:  []int{1, 2, 3, 4, 5, 6},
			callback: func(i, v int) bool {
				return v%2 == 0
			},
			want: []int{2, 4, 6},
		},
		{
			name: "filter by index",
			arr:  []int{10, 20, 30, 40},
			callback: func(i, v int) bool {
				return i > 1
			},
			want: []int{30, 40},
		},
		{
			name: "no matches",
			arr:  []int{1, 3, 5},
			callback: func(i, v int) bool {
				return v%2 == 0
			},
			want: []int{},
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(i, v int) bool {
				return true
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayFilter(tt.arr, tt.callback)

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
		name     string
		arr      []int
		callback func(int) bool
		want     bool
	}{
		{
			name: "all elements match",
			arr:  []int{2, 4, 6, 8},
			callback: func(v int) bool {
				return v%2 == 0
			},
			want: true,
		},
		{
			name: "not all elements match",
			arr:  []int{2, 3, 6, 8},
			callback: func(v int) bool {
				return v%2 == 0
			},
			want: false,
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(v int) bool {
				return v > 10
			},
			want: true,
		},
		{
			name: "single element matching",
			arr:  []int{5},
			callback: func(v int) bool {
				return v > 0
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.ArrayEvery(tt.arr, tt.callback)

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
				{name: "Alice", age: 35}, // Last Alice wins
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
		callback  func(int, int) bool
		wantValue int
		wantFound bool
	}{
		{
			name: "find first matching element",
			arr:  []int{1, 2, 3, 4, 5},
			callback: func(i, v int) bool {
				return v > 2
			},
			wantValue: 3,
			wantFound: true,
		},
		{
			name: "no matching element",
			arr:  []int{1, 2, 3},
			callback: func(i, v int) bool {
				return v > 10
			},
			wantValue: 0,
			wantFound: false,
		},
		{
			name: "find by index",
			arr:  []int{10, 20, 30},
			callback: func(i, v int) bool {
				return i == 1
			},
			wantValue: 20,
			wantFound: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(i, v int) bool {
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

			gotValue, gotFound := arrays.ArrayFind(tt.arr, tt.callback)

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
		callback  func(int, int) bool
		wantIndex int
		wantFound bool
	}{
		{
			name: "find first matching index",
			arr:  []int{1, 2, 3, 4, 5},
			callback: func(i, v int) bool {
				return v > 2
			},
			wantIndex: 2,
			wantFound: true,
		},
		{
			name: "no matching element",
			arr:  []int{1, 2, 3},
			callback: func(i, v int) bool {
				return v > 10
			},
			wantIndex: -1,
			wantFound: false,
		},
		{
			name: "find first element",
			arr:  []int{10, 20, 30},
			callback: func(i, v int) bool {
				return v == 10
			},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "empty array",
			arr:  []int{},
			callback: func(i, v int) bool {
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

			gotIndex, gotFound := arrays.ArrayFindIndex(tt.arr, tt.callback)

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
