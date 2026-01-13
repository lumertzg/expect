// Package expect provides simple test assertions.
package expect

import (
	"cmp"
	"maps"
	"slices"
)

// T is the interface required for test assertions. Both *testing.T and *testing.B satisfy it.
type T interface {
	Helper()
	Errorf(format string, args ...any)
}

// Equal asserts that expected and actual are equal.
func Equal[V comparable](t T, expected, actual V) {
	t.Helper()
	if expected != actual {
		failMismatch(t, expected, actual)
	}
}

// NotEqual asserts that unexpected and actual are not equal.
func NotEqual[V comparable](t T, unexpected, actual V) {
	t.Helper()
	if unexpected == actual {
		failMatch(t, actual)
	}
}

// Less asserts that a < b.
func Less[V cmp.Ordered](t T, a, b V) {
	t.Helper()
	if a >= b {
		failCompare(t, a, "<", b)
	}
}

// LessOrEqual asserts that a <= b.
func LessOrEqual[V cmp.Ordered](t T, a, b V) {
	t.Helper()
	if a > b {
		failCompare(t, a, "<=", b)
	}
}

// Greater asserts that a > b.
func Greater[V cmp.Ordered](t T, a, b V) {
	t.Helper()
	if a <= b {
		failCompare(t, a, ">", b)
	}
}

// GreaterOrEqual asserts that a >= b.
func GreaterOrEqual[V cmp.Ordered](t T, a, b V) {
	t.Helper()
	if a < b {
		failCompare(t, a, ">=", b)
	}
}

// True asserts that value is true.
func True(t T, value bool) {
	t.Helper()
	if !value {
		failMismatch(t, true, false)
	}
}

// False asserts that value is false.
func False(t T, value bool) {
	t.Helper()
	if value {
		failMismatch(t, false, true)
	}
}

// Nil asserts that value is nil.
func Nil(t T, value any) {
	t.Helper()
	if value != nil {
		failMismatch(t, nil, value)
	}
}

// NotNil asserts that value is not nil.
func NotNil(t T, value any) {
	t.Helper()
	if value == nil {
		failMatch(t, value)
	}
}

// Error asserts that err is not nil.
func Error(t T, err error) {
	t.Helper()
	if err == nil {
		failMismatch(t, "error", nil)
	}
}

// NoError asserts that err is nil.
func NoError(t T, err error) {
	t.Helper()
	if err != nil {
		failMismatch(t, nil, err)
	}
}

// EqualSlice asserts that expected and actual slices are equal.
func EqualSlice[V comparable](t T, expected, actual []V) {
	t.Helper()
	if !slices.Equal(expected, actual) {
		failMismatch(t, expected, actual)
	}
}

// NotEqualSlice asserts that unexpected and actual slices are not equal.
func NotEqualSlice[V comparable](t T, unexpected, actual []V) {
	t.Helper()
	if slices.Equal(unexpected, actual) {
		failMatch(t, actual)
	}
}

// EqualMap asserts that expected and actual maps are equal.
func EqualMap[K, V comparable](t T, expected, actual map[K]V) {
	t.Helper()
	if !maps.Equal(expected, actual) {
		failMismatch(t, expected, actual)
	}
}

// NotEqualMap asserts that unexpected and actual maps are not equal.
func NotEqualMap[K, V comparable](t T, unexpected, actual map[K]V) {
	t.Helper()
	if maps.Equal(unexpected, actual) {
		failMatch(t, actual)
	}
}

func failMismatch(t T, expected, actual any) {
	t.Helper()
	t.Errorf("expected %v, got %v", expected, actual)
}

func failMatch(t T, value any) {
	t.Helper()
	t.Errorf("expected different values, got equal: %v", value)
}

func failCompare(t T, a any, op string, b any) {
	t.Helper()
	t.Errorf("expected %v %s %v", a, op, b)
}
