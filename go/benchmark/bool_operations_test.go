package main

import "testing"

func TestCompare(t *testing.T) {
	testcases := []struct {
		name     string
		fun      func(compare bool) bool
		input    bool
		expected bool
	}{
		{"short true", compareShort, true, false},
		{"short false", compareShort, true, false},
		{"long true", compareLong, true, false},
		{"long false", compareLong, true, false},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.fun(tc.input)
			if tc.expected != res {
				t.Errorf("expected %t but got %t", tc.expected, res)
			}
		})
	}
}

func BenchmarkCompare(b *testing.B) {
	testcases := []struct {
		name string
		fun  func(compare bool) bool
	}{
		{"short", compareShort},
		{"long", compareLong},
	}

	for _, tc := range testcases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.fun(true)
				tc.fun(false)
			}
		})
	}
}

func compareShort(compare bool) bool {
	return !compare
}

func compareLong(compare bool) bool {
	return compare == false
}
