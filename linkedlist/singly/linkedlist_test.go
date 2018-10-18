package singly

import "testing"

func TestNewList(t *testing.T) {
	l := NewList()

	t.Run("head and count", func(t *testing.T) {
		if l.head != nil || l.count != 0 {
			t.Errorf("Got head=%v count=%d expected head=nil count=0", l.head, l.count)
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		expected := true
		got := l.IsEmpty()
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})

	t.Run("String with zero value", func(t *testing.T) {
		expected := "[]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})
}
func TestList_AddHead(t *testing.T) {
	l := NewList()

	l.AddHead(1)
	t.Run("count with one value", func(t *testing.T) {
		expected := 1
		got := l.count
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("IsEmpty with one value", func(t *testing.T) {
		expected := false
		got := l.IsEmpty()
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("Size with one value", func(t *testing.T) {
		expected := 1
		got := l.Size()
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("String with one value", func(t *testing.T) {
		expected := "[1]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

	l.AddHead(2)
	l.AddHead(3)
	t.Run("count with 3 values", func(t *testing.T) {
		expected := 3
		got := l.count
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("Size with 3 values", func(t *testing.T) {
		expected := 3
		got := l.Size()
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("String with 3 values", func(t *testing.T) {
		expected := "[3, 2, 1]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

}
func TestList_AddTail(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	t.Run("count with one value", func(t *testing.T) {
		expected := 1
		got := l.count
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("IsEmpty with one value", func(t *testing.T) {
		expected := false
		got := l.IsEmpty()
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("Size with one value", func(t *testing.T) {
		expected := 1
		got := l.Size()
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("String with one value", func(t *testing.T) {
		expected := "[1]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

	l.AddTail(2)
	l.AddTail(3)
	t.Run("count with 3 values", func(t *testing.T) {
		expected := 3
		got := l.count
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("Size with 3 values", func(t *testing.T) {
		expected := 3
		got := l.Size()
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
	})
	t.Run("String with 3 values", func(t *testing.T) {
		expected := "[1, 2, 3]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})
}
func TestList_SortedInsert(t *testing.T) {
	l := NewList()

	l.SortedInsert(1)
	t.Run("1", func(t *testing.T) {
		expected := "[1]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

	l.SortedInsert(10)
	t.Run("10", func(t *testing.T) {
		expected := "[1, 10]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

	l.SortedInsert(0)
	t.Run("0", func(t *testing.T) {
		expected := "[0, 1, 10]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

	l.SortedInsert(5)
	t.Run("5", func(t *testing.T) {
		expected := "[0, 1, 5, 10]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})

	l.SortedInsert(3)
	t.Run("3", func(t *testing.T) {
		expected := "[0, 1, 3, 5, 10]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})
}
