package expect

import (
	"errors"
	"testing"
)

type mockT struct {
	failed bool
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
