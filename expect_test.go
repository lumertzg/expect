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

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		Nil(m, "not nil")
		if !m.failed {
			t.Error("expected fail")
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

	t.Run("fail", func(t *testing.T) {
		m := &mockT{}
		NotNil(m, nil)
		if !m.failed {
			t.Error("expected fail")
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
