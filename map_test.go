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
		name     string
		arr      map[string]int
		callback func(string, int) bool
		want     map[string]int
	}{
		{
			name: "filter values greater than 2",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			callback: func(k string, v int) bool {
				return v > 2
			},
			want: map[string]int{"c": 3, "d": 4},
		},
		{
			name: "filter by key",
			arr:  map[string]int{"apple": 1, "banana": 2, "apricot": 3},
			callback: func(k string, v int) bool {
				return k[0] == 'a'
			},
			want: map[string]int{"apple": 1, "apricot": 3},
		},
		{
			name: "no matches",
			arr:  map[string]int{"a": 1, "b": 2, "c": 3},
			callback: func(k string, v int) bool {
				return v > 10
			},
			want: map[string]int{},
		},
		{
			name: "all matches",
			arr:  map[string]int{"a": 1, "b": 2},
			callback: func(k string, v int) bool {
				return true
			},
			want: map[string]int{"a": 1, "b": 2},
		},
		{
			name: "empty map",
			arr:  map[string]int{},
			callback: func(k string, v int) bool {
				return true
			},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := arrays.MapFilter(tt.arr, tt.callback)

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
