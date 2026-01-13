package examples

import (
	"errors"
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
	expect.Error(t, errors.New("something went wrong"))
	expect.NoError(t, nil)
}

func TestSlices(t *testing.T) {
	expect.EqualSlice(t, []int{1, 2, 3}, []int{1, 2, 3})
	expect.NotEqualSlice(t, []int{1, 2, 3}, []int{4, 5, 6})

	expect.EqualSlice(t, []string{"a", "b"}, []string{"a", "b"})
}

func TestMaps(t *testing.T) {
	expect.EqualMap(t, map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1, "b": 2})
	expect.NotEqualMap(t, map[string]int{"a": 1}, map[string]int{"a": 2})
}
