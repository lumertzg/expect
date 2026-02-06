package examples

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lumertzg/expect"
)

func TestEqual(t *testing.T) {
	expect.Equal(t, 10, 5+5)
	expect.Equal(t, "hello", "hello")
	expect.Equal(t, true, true)
}

func TestNotEqual(t *testing.T) {
	expect.NotEqual(t, 10, 20)
	expect.NotEqual(t, "hello", "world")
}

func TestComparisons(t *testing.T) {
	expect.Less(t, 1, 2)
	expect.LessOrEqual(t, 1, 2)
	expect.LessOrEqual(t, 2, 2)
	expect.Greater(t, 2, 1)
	expect.GreaterOrEqual(t, 2, 1)
	expect.GreaterOrEqual(t, 2, 2)
}

func TestBooleans(t *testing.T) {
	expect.True(t, 10 > 5)
	expect.False(t, 10 < 5)
}

func TestNil(t *testing.T) {
	expect.Nil(t, nil)

	value := 42
	expect.NotNil(t, &value)
}

func TestErrors(t *testing.T) {
	baseErr := errors.New("something went wrong")
	wrappedErr := fmt.Errorf("request failed: %w", baseErr)
	var targetErr *customErr

	expect.Error(t, errors.New("something went wrong"))
	expect.NoError(t, nil)
	expect.ErrorIs(t, wrappedErr, baseErr)
	expect.NotErrorIs(t, errors.New("other"), baseErr)
	expect.ErrorAs(t, fmt.Errorf("wrapped custom: %w", &customErr{}), &targetErr)
}

func TestSlices(t *testing.T) {
	expect.EqualSlice(t, []int{1, 2, 3}, []int{1, 2, 3})
	expect.NotEqualSlice(t, []int{1, 2, 3}, []int{4, 5, 6})
	expect.ContainsSlice(t, []int{1, 2, 3}, 2)
	expect.NotContainsSlice(t, []int{1, 2, 3}, 4)
	expect.EqualSlice(t, []string{"a", "b"}, []string{"a", "b"})
}

func TestContainsString(t *testing.T) {
	expect.ContainsString(t, "hello world", "world")
	expect.NotContainsString(t, "hello world", "golang")
}

func TestMaps(t *testing.T) {
	expect.EqualMap(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	expect.NotEqualMap(t, map[string]int{"a": 1}, map[string]int{"a": 2})
	expect.ContainsMapKey(t, map[string]int{"a": 1}, "a")
}

func TestLengthAndEmpty(t *testing.T) {
	expect.Len(t, []int{1, 2, 3}, 3)
	expect.Empty(t, "")
	expect.NotEmpty(t, "hello")
}

type customErr struct{}

func (e *customErr) Error() string {
	return "custom"
}
