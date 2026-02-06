package expect

import (
	"errors"
	"fmt"
	"testing"
)

type mockT struct {
	failed bool
}

type customError struct{}

func (e *customError) Error() string {
	return "custom"
}

func (m *mockT) Helper() {}

func (m *mockT) Errorf(format string, args ...any) {
	m.failed = true
}

func TestEqual(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		Equal(m, 1, 1)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		Equal(m, 1, 2)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotEqual(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotEqual(m, 1, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotEqual(m, 1, 1)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestLess(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		Less(m, 1, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		Less(m, 2, 1)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestLessOrEqual(t *testing.T) {
	t.Run("pass less", func(t *testing.T) {
		m := &mockT{}
		LessOrEqual(m, 1, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass equal", func(t *testing.T) {
		m := &mockT{}
		LessOrEqual(m, 2, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		LessOrEqual(m, 3, 2)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestGreater(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		Greater(m, 2, 1)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		Greater(m, 1, 2)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestGreaterOrEqual(t *testing.T) {
	t.Run("pass greater", func(t *testing.T) {
		m := &mockT{}
		GreaterOrEqual(m, 2, 1)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass equal", func(t *testing.T) {
		m := &mockT{}
		GreaterOrEqual(m, 2, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		GreaterOrEqual(m, 1, 2)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestTrue(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		True(m, true)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		True(m, false)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestFalse(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		False(m, false)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		False(m, true)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNil(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		Nil(m, nil)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass typed nil pointer", func(t *testing.T) {
		m := &mockT{}
		var p *int = nil
		Nil(m, p)
		if m.failed {
			t.Error("expected pass for typed nil pointer")
		}
	})

	t.Run("pass nil slice", func(t *testing.T) {
		m := &mockT{}
		var s []int = nil
		Nil(m, s)
		if m.failed {
			t.Error("expected pass for nil slice")
		}
	})

	t.Run("pass nil map", func(t *testing.T) {
		m := &mockT{}
		var mp map[string]int = nil
		Nil(m, mp)
		if m.failed {
			t.Error("expected pass for nil map")
		}
	})

	t.Run("pass nil chan", func(t *testing.T) {
		m := &mockT{}
		var ch chan int = nil
		Nil(m, ch)
		if m.failed {
			t.Error("expected pass for nil chan")
		}
	})

	t.Run("pass nil func", func(t *testing.T) {
		m := &mockT{}
		var f func() = nil
		Nil(m, f)
		if m.failed {
			t.Error("expected pass for nil func")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		Nil(m, "not nil")
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail non-nil pointer", func(t *testing.T) {
		m := &mockT{}
		x := 42
		Nil(m, &x)
		if !m.failed {
			t.Error("expected fail for non-nil pointer")
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotNil(m, "not nil")
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass non-nil pointer", func(t *testing.T) {
		m := &mockT{}
		x := 42
		NotNil(m, &x)
		if m.failed {
			t.Error("expected pass for non-nil pointer")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotNil(m, nil)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail typed nil pointer", func(t *testing.T) {
		m := &mockT{}
		var p *int = nil
		NotNil(m, p)
		if !m.failed {
			t.Error("expected fail for typed nil pointer")
		}
	})

	t.Run("fail nil slice", func(t *testing.T) {
		m := &mockT{}
		var s []int = nil
		NotNil(m, s)
		if !m.failed {
			t.Error("expected fail for nil slice")
		}
	})

	t.Run("fail nil map", func(t *testing.T) {
		m := &mockT{}
		var mp map[string]int = nil
		NotNil(m, mp)
		if !m.failed {
			t.Error("expected fail for nil map")
		}
	})
}

func TestError(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		Error(m, errors.New("test"))
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		Error(m, nil)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNoError(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NoError(m, nil)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NoError(m, errors.New("test"))
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestErrorIs(t *testing.T) {
	target := errors.New("target")

	t.Run("pass direct match", func(t *testing.T) {
		m := &mockT{}
		ErrorIs(m, target, target)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass wrapped match", func(t *testing.T) {
		m := &mockT{}
		err := fmt.Errorf("wrapped: %w", target)
		ErrorIs(m, err, target)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		ErrorIs(m, errors.New("different"), target)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotErrorIs(t *testing.T) {
	target := errors.New("target")

	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotErrorIs(m, errors.New("different"), target)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail direct match", func(t *testing.T) {
		m := &mockT{}
		NotErrorIs(m, target, target)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail wrapped match", func(t *testing.T) {
		m := &mockT{}
		err := fmt.Errorf("wrapped: %w", target)
		NotErrorIs(m, err, target)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestErrorAs(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		wrapped := fmt.Errorf("wrapped: %w", &customError{})
		var target *customError
		ErrorAs(m, wrapped, &target)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail no match", func(t *testing.T) {
		m := &mockT{}
		var target *customError
		ErrorAs(m, errors.New("different"), &target)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail nil target", func(t *testing.T) {
		m := &mockT{}
		ErrorAs(m, errors.New("test"), nil)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail non-pointer target", func(t *testing.T) {
		m := &mockT{}
		var target *customError
		ErrorAs(m, errors.New("test"), target)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail pointer to non-error", func(t *testing.T) {
		m := &mockT{}
		value := 1
		ErrorAs(m, errors.New("test"), &value)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestEqualSlice(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		EqualSlice(m, []int{1, 2, 3}, []int{1, 2, 3})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		EqualSlice(m, []int{1, 2, 3}, []int{1, 2, 4})
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotEqualSlice(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotEqualSlice(m, []int{1, 2, 3}, []int{1, 2, 4})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotEqualSlice(m, []int{1, 2, 3}, []int{1, 2, 3})
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestContainsSlice(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		ContainsSlice(m, []int{1, 2, 3}, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		ContainsSlice(m, []int{1, 2, 3}, 4)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotContainsSlice(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotContainsSlice(m, []int{1, 2, 3}, 4)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotContainsSlice(m, []int{1, 2, 3}, 2)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestContainsString(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		ContainsString(m, "hello world", "world")
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		ContainsString(m, "hello world", "golang")
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotContainsString(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotContainsString(m, "hello world", "golang")
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotContainsString(m, "hello world", "world")
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestLen(t *testing.T) {
	t.Run("pass string", func(t *testing.T) {
		m := &mockT{}
		Len(m, "abc", 3)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass slice", func(t *testing.T) {
		m := &mockT{}
		Len(m, []int{1, 2, 3}, 3)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass map", func(t *testing.T) {
		m := &mockT{}
		Len(m, map[string]int{"a": 1, "b": 2}, 2)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail mismatch", func(t *testing.T) {
		m := &mockT{}
		Len(m, []int{1, 2, 3}, 2)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail unsupported", func(t *testing.T) {
		m := &mockT{}
		Len(m, 10, 1)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Run("pass nil", func(t *testing.T) {
		m := &mockT{}
		Empty(m, nil)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass empty string", func(t *testing.T) {
		m := &mockT{}
		Empty(m, "")
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass nil slice", func(t *testing.T) {
		m := &mockT{}
		var values []int
		Empty(m, values)
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass empty map", func(t *testing.T) {
		m := &mockT{}
		Empty(m, map[string]int{})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail non-empty", func(t *testing.T) {
		m := &mockT{}
		Empty(m, "abc")
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail unsupported", func(t *testing.T) {
		m := &mockT{}
		Empty(m, 10)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotEmpty(t *testing.T) {
	t.Run("pass string", func(t *testing.T) {
		m := &mockT{}
		NotEmpty(m, "abc")
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass slice", func(t *testing.T) {
		m := &mockT{}
		NotEmpty(m, []int{1})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("pass map", func(t *testing.T) {
		m := &mockT{}
		NotEmpty(m, map[string]int{"a": 1})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail empty", func(t *testing.T) {
		m := &mockT{}
		NotEmpty(m, "")
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail nil", func(t *testing.T) {
		m := &mockT{}
		NotEmpty(m, nil)
		if !m.failed {
			t.Error("expected fail")
		}
	})

	t.Run("fail unsupported", func(t *testing.T) {
		m := &mockT{}
		NotEmpty(m, 10)
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestEqualMap(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		EqualMap(m, map[string]int{"a": 1}, map[string]int{"a": 1})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		EqualMap(m, map[string]int{"a": 1}, map[string]int{"a": 2})
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestNotEqualMap(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		NotEqualMap(m, map[string]int{"a": 1}, map[string]int{"a": 2})
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotEqualMap(m, map[string]int{"a": 1}, map[string]int{"a": 1})
		if !m.failed {
			t.Error("expected fail")
		}
	})
}

func TestContainsMapKey(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		m := &mockT{}
		ContainsMapKey(m, map[string]int{"a": 1}, "a")
		if m.failed {
			t.Error("expected pass")
		}
	})

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		ContainsMapKey(m, map[string]int{"a": 1}, "b")
		if !m.failed {
			t.Error("expected fail")
		}
	})
}
