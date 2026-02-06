// Package expect provides simple test assertions.
package expect

import (
	"cmp"
	"errors"
	"maps"
	"reflect"
	"slices"
	"strings"
)

var errorType = reflect.TypeFor[error]()

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
	if !isNil(value) {
		failMismatch(t, nil, value)
	}
}

// NotNil asserts that value is not nil.
func NotNil(t T, value any) {
	t.Helper()
	if isNil(value) {
		failMatch(t, value)
	}
}

// isNil checks if a value is nil, handling typed nil pointers correctly.
// In Go, a typed nil pointer like (*T)(nil) wrapped in an interface is not
// equal to nil because the interface contains type information.
func isNil(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func:
		return v.IsNil()
	}
	return false
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

// ErrorIs asserts that err matches target using errors.Is.
func ErrorIs(t T, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf("expected error %v to match %v", err, target)
	}
}

// NotErrorIs asserts that err does not match target using errors.Is.
func NotErrorIs(t T, err, target error) {
	t.Helper()
	if errors.Is(err, target) {
		t.Errorf("expected error %v not to match %v", err, target)
	}
}

// ErrorAs asserts that err matches target using errors.As.
func ErrorAs(t T, err error, target any) {
	t.Helper()
	v := reflect.ValueOf(target)
	if target == nil || v.Kind() != reflect.Pointer || v.IsNil() {
		t.Errorf("expected target to be a non-nil pointer, got %T", target)
		return
	}

	typeToMatch := v.Elem().Type()
	if !typeToMatch.Implements(errorType) && typeToMatch.Kind() != reflect.Interface {
		t.Errorf("expected target to point to an error or interface type, got %T", target)
		return
	}

	if !errors.As(err, target) {
		t.Errorf("expected error %v to match target type %v", err, typeToMatch)
	}
}

// EqualSlice asserts that expected and actual slices are equal.
func EqualSlice[S ~[]E, E comparable](t T, expected, actual S) {
	t.Helper()
	if !slices.Equal(expected, actual) {
		failMismatch(t, expected, actual)
	}
}

// NotEqualSlice asserts that unexpected and actual slices are not equal.
func NotEqualSlice[S ~[]E, E comparable](t T, unexpected, actual S) {
	t.Helper()
	if slices.Equal(unexpected, actual) {
		failMatch(t, actual)
	}
}

// ContainsSlice asserts that values contains item.
func ContainsSlice[S ~[]E, E comparable](t T, values S, item E) {
	t.Helper()
	if !slices.Contains(values, item) {
		t.Errorf("expected %v to contain %v", values, item)
	}
}

// NotContainsSlice asserts that values does not contain item.
func NotContainsSlice[S ~[]E, E comparable](t T, values S, item E) {
	t.Helper()
	if slices.Contains(values, item) {
		t.Errorf("expected %v not to contain %v", values, item)
	}
}

// ContainsString asserts that s contains substr.
func ContainsString(t T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected %q to contain %q", s, substr)
	}
}

// NotContainsString asserts that s does not contain substr.
func NotContainsString(t T, s, substr string) {
	t.Helper()
	if strings.Contains(s, substr) {
		t.Errorf("expected %q not to contain %q", s, substr)
	}
}

// Len asserts that value has the expected length.
func Len(t T, value any, expected int) {
	t.Helper()
	actual, ok := valueLen(value)
	if !ok {
		t.Errorf("expected value with length, got %T", value)
		return
	}
	if actual != expected {
		t.Errorf("expected length %d, got %d", expected, actual)
	}
}

// Empty asserts that value is empty.
func Empty(t T, value any) {
	t.Helper()
	if isNil(value) {
		return
	}

	length, ok := valueLen(value)
	if !ok {
		t.Errorf("expected empty value, got unsupported type %T", value)
		return
	}

	if length != 0 {
		t.Errorf("expected empty value, got %v", value)
	}
}

// NotEmpty asserts that value is not empty.
func NotEmpty(t T, value any) {
	t.Helper()
	if isNil(value) {
		t.Errorf("expected non-empty value, got %v", value)
		return
	}

	length, ok := valueLen(value)
	if !ok {
		t.Errorf("expected non-empty value, got unsupported type %T", value)
		return
	}

	if length == 0 {
		t.Errorf("expected non-empty value, got %v", value)
	}
}

// EqualMap asserts that expected and actual maps are equal.
func EqualMap[M ~map[K]V, K, V comparable](t T, expected, actual M) {
	t.Helper()
	if !maps.Equal(expected, actual) {
		failMismatch(t, expected, actual)
	}
}

// NotEqualMap asserts that unexpected and actual maps are not equal.
func NotEqualMap[M ~map[K]V, K, V comparable](t T, unexpected, actual M) {
	t.Helper()
	if maps.Equal(unexpected, actual) {
		failMatch(t, actual)
	}
}

// ContainsMapKey asserts that m contains key.
func ContainsMapKey[M ~map[K]V, K comparable, V any](t T, m M, key K) {
	t.Helper()
	if _, ok := m[key]; !ok {
		t.Errorf("expected map %v to contain key %v", m, key)
	}
}

func valueLen(value any) (int, bool) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len(), true
	default:
		return 0, false
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
