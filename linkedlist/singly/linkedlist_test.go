package singly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
func TestList_AddNodeAtTail(t *testing.T) {
	l := NewList()

	l.AddNodeAtTail(Node{value: 1})
	t.Run("count with one value", func(t *testing.T) {
		expected := 1
		got := l.count
		assert.Equal(t, expected, got)
	})
	t.Run("String with one value", func(t *testing.T) {
		expected := "[1]"
		got := l.String()
		assert.Equal(t, expected, got)
	})

	l.AddNodeAtTail(Node{value: 2})
	l.AddNodeAtTail(Node{value: 3})
	t.Run("count with 3 values", func(t *testing.T) {
		expected := 3
		got := l.count
		assert.Equal(t, expected, got)
	})
	t.Run("String with 3 values", func(t *testing.T) {
		expected := "[1, 2, 3]"
		got := l.String()
		assert.Equal(t, expected, got)
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
func TestList_IsPresent(t *testing.T) {
	l := NewList()

	l.AddHead(2)
	l.AddHead(3)
	l.AddHead(14)
	l.AddHead(100)
	t.Run("2", func(t *testing.T) {
		expected := true
		got := l.IsPresent(2)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("3", func(t *testing.T) {
		expected := true
		got := l.IsPresent(3)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("100", func(t *testing.T) {
		expected := true
		got := l.IsPresent(100)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("-1", func(t *testing.T) {
		expected := false
		got := l.IsPresent(-1)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("5", func(t *testing.T) {
		expected := false
		got := l.IsPresent(5)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
	t.Run("222", func(t *testing.T) {
		expected := false
		got := l.IsPresent(222)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})
}
func TestList_RemoveHead(t *testing.T) {
	l := NewList()

	t.Run("empty list", func(t *testing.T) {
		_, err := l.RemoveHead()
		assertError(t, err, ErrEmptyList)
	})

	l.AddHead(2)
	t.Run("[2]", func(t *testing.T) {
		expected := 2
		got, err := l.RemoveHead()
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})

	})

	l.AddHead(3)
	l.AddHead(4)
	l.AddHead(5)
	t.Run("[5, 4, 3]", func(t *testing.T) {
		expected := 5
		got, err := l.RemoveHead()
		if got != expected {
			t.Errorf("got %d want %d", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[4, 3]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})
	})

}
func TestList_DeleteNode(t *testing.T) {
	l := NewList()
	delValue := 100

	t.Run("empty list", func(t *testing.T) {
		_, err := l.DeleteNode(delValue)
		assertError(t, err, ErrEmptyList)
	})

	l.AddHead(3)
	l.AddHead(4)
	l.AddHead(delValue)
	l.AddHead(5)
	t.Run("delete 100 from [5, 100, 4, 3]", func(t *testing.T) {
		expected := true
		got, err := l.DeleteNode(delValue)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[5, 4, 3]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})
	})

	l.AddTail(delValue)
	t.Run("delete 100 from [5, 4, 3, 100]", func(t *testing.T) {
		expected := true
		got, err := l.DeleteNode(delValue)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[5, 4, 3]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})
	})

	delValue = 200
	t.Run("delete 200 from [5, 4, 3]", func(t *testing.T) {
		expected := false
		got, err := l.DeleteNode(delValue)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[5, 4, 3]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})
	})

}
func TestList_DeleteNodes(t *testing.T) {
	l := NewList()
	delValue := 100

	t.Run("empty list", func(t *testing.T) {
		_, err := l.DeleteNode(delValue)
		assertError(t, err, ErrEmptyList)
	})

	l.AddHead(delValue)
	l.AddHead(3)
	l.AddHead(4)
	l.AddHead(delValue)
	l.AddHead(5)
	l.AddHead(delValue)
	t.Run("delete 100 from [100, 5, 100, 4, 3, 100]", func(t *testing.T) {
		expected := true
		got, err := l.DeleteNodes(delValue)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[5, 4, 3]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})
	})

	delValue = 200
	t.Run("delete 200 from [5, 4, 3]", func(t *testing.T) {
		expected := false
		got, err := l.DeleteNodes(delValue)
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
		assertError(t, err, nil)

		t.Run("String", func(t *testing.T) {
			expected := "[5, 4, 3]"
			got := l.String()
			if got != expected {
				t.Errorf("got %s want %s", got, expected)
			}
		})
	})

}
func TestList_FreeList(t *testing.T) {
	l := NewList()

	l.AddHead(3)
	l.AddHead(5)
	l.FreeList()

	t.Run("IsEmpty", func(t *testing.T) {
		expected := true
		got := l.IsEmpty()
		if got != expected {
			t.Errorf("got %t want %t", got, expected)
		}
	})

	t.Run("String", func(t *testing.T) {
		expected := "[]"
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	})
}
func TestList_Reverse(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)
	t.Run("[1, 2, 3, 4, 5]", func(t *testing.T) {
		expected := "[5, 4, 3, 2, 1]"
		l.Reverse()
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}

	})

	l.FreeList()
	l.AddHead(1)
	t.Run("[1]", func(t *testing.T) {
		expected := "[1]"
		l.Reverse()
		got := l.String()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}

	})
}
func TestList_ReverseRecurse(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)
	t.Run("[1, 2, 3, 4, 5]", func(t *testing.T) {
		expected := "[5, 4, 3, 2, 1]"
		l.ReverseRecurse()
		got := l.String()
		assert.Equal(t, expected, got)

	})

	l.FreeList()
	l.AddHead(1)
	t.Run("[1]", func(t *testing.T) {
		expected := "[1]"
		l.ReverseRecurse()
		got := l.String()
		assert.Equal(t, expected, got)
	})
}
func TestList_RemoveDuplicate(t *testing.T) {
	l := NewList()

	l.SortedInsert(1)
	l.SortedInsert(2)
	l.SortedInsert(2)
	l.SortedInsert(4)
	l.SortedInsert(5)
	l.SortedInsert(5)
	t.Run("[1, 2, 2, 4, 5, 5]", func(t *testing.T) {
		expected := "[1, 2, 4, 5]"
		l.RemoveDuplicate()
		got := l.String()
		assert.Equal(t, expected, got)

	})
}
func TestList_CopyListReversed(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)
	t.Run("[1, 2, 3, 4, 5]", func(t *testing.T) {
		expected := "[5, 4, 3, 2, 1]"
		got := l.CopyListReversed().String()
		assert.Equal(t, expected, got)

	})
}
func TestList_CopyList(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)
	t.Run("[1, 2, 3, 4, 5]", func(t *testing.T) {
		expected := "[1, 2, 3, 4, 5]"
		got := l.CopyList().String()
		assert.Equal(t, expected, got)

	})
}
func TestList_CompareList(t *testing.T) {
	l := NewList()
	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)

	t.Run("true", func(t *testing.T) {
		l2 := NewList()
		l2.AddTail(1)
		l2.AddTail(2)
		l2.AddTail(3)
		l2.AddTail(4)
		l2.AddTail(5)

		expected := true
		got := l.CompareList(l2)
		assert.Equal(t, expected, got)
	})
	t.Run("false", func(t *testing.T) {
		l2 := NewList()
		l2.AddTail(1)
		l2.AddTail(2)
		l2.AddTail(33)
		l2.AddTail(4)
		l2.AddTail(5)

		expected := false
		got := l.CompareList(l2)
		assert.Equal(t, expected, got)
	})
}
func TestList_FindLength(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)
	t.Run("[1, 2, 3, 4, 5]", func(t *testing.T) {
		expected := 5
		got := l.FindLength()
		assert.Equal(t, expected, got)

	})
}
func TestList_NthNodeFromBeginning(t *testing.T) {
	l := NewList()
	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)

	t.Run("[1, 2, 3, 4, 5], index=3", func(t *testing.T) {
		expected := 3
		got, err := l.NthNodeFromBeginning(3)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, got)

	})
	t.Run("[1, 2, 3, 4, 5], index=5", func(t *testing.T) {
		expected := 5
		got, err := l.NthNodeFromBeginning(5)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, got)

	})
	t.Run("[1, 2, 3, 4, 5], index=1", func(t *testing.T) {
		expected := 1
		got, err := l.NthNodeFromBeginning(1)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, got)

	})
	t.Run("[1, 2, 3, 4, 5], index=-1", func(t *testing.T) {
		expected := 0
		got, err := l.NthNodeFromBeginning(-1)
		assert.Equal(t, ErrIndexOutOfRange, err)
		assert.Equal(t, expected, got)
	})
	t.Run("[1, 2, 3, 4, 5], index=6", func(t *testing.T) {
		expected := 0
		got, err := l.NthNodeFromBeginning(6)
		assert.Equal(t, ErrIndexOutOfRange, err)
		assert.Equal(t, expected, got)
	})
}
func TestList_NthNodeFromEnd(t *testing.T) {
	l := NewList()
	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)

	t.Run("[1, 2, 3, 4, 5], index=3", func(t *testing.T) {
		expected := 3
		got, err := l.NthNodeFromEnd(3)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, got)

	})
	t.Run("[1, 2, 3, 4, 5], index=5", func(t *testing.T) {
		expected := 1
		got, err := l.NthNodeFromEnd(5)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, got)

	})
	t.Run("[1, 2, 3, 4, 5], index=1", func(t *testing.T) {
		expected := 5
		got, err := l.NthNodeFromEnd(1)
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, got)

	})
	t.Run("[1, 2, 3, 4, 5], index=-1", func(t *testing.T) {
		expected := 0
		got, err := l.NthNodeFromEnd(-1)
		assert.Equal(t, ErrIndexOutOfRange, err)
		assert.Equal(t, expected, got)
	})
	t.Run("[1, 2, 3, 4, 5], index=6", func(t *testing.T) {
		expected := 0
		got, err := l.NthNodeFromEnd(6)
		assert.Equal(t, ErrIndexOutOfRange, err)
		assert.Equal(t, expected, got)
	})
}
func TestList_LoopDetect(t *testing.T) {
	t.Run("no loop", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)
		l.AddTail(4)
		l.AddTail(5)

		expected := false
		got := l.LoopDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("loop. 2 nodes", func(t *testing.T) {
		l := NewList()
		n1 := new(Node)
		n2 := new(Node)
		n1.next = n2
		n2.next = n1
		l.head = n1
		l.count = 2

		expected := true
		got := l.LoopDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("circular loop", func(t *testing.T) {
		l := getCircularLoopList()

		expected := true
		got := l.LoopDetect()
		assert.Equal(t, expected, got)
	})

	t.Run("loop", func(t *testing.T) {
		l := getLoopList()

		expected := true
		got := l.LoopDetect()
		assert.Equal(t, expected, got)
	})
}
func TestList_ReverseListLoopDetect(t *testing.T) {
	t.Run("no loop", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)
		l.AddTail(4)
		l.AddTail(5)

		expected := false
		got := l.ReverseListLoopDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("loop. 2 nodes", func(t *testing.T) {
		l := NewList()
		n1 := new(Node)
		n2 := new(Node)
		n1.next = n2
		n2.next = n1
		l.head = n1
		l.count = 2

		expected := true
		got := l.ReverseListLoopDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("circular loop", func(t *testing.T) {
		l := getCircularLoopList()

		expected := true
		got := l.ReverseListLoopDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("loop", func(t *testing.T) {
		l := getLoopList()

		expected := true
		got := l.ReverseListLoopDetect()
		assert.Equal(t, expected, got)
	})
}
func TestList_LoopTypeDetect(t *testing.T) {
	t.Run("no loop", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)
		l.AddTail(4)
		l.AddTail(5)

		expected := 0
		got := l.LoopTypeDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("circular loop. 2 nodes", func(t *testing.T) {
		l := NewList()
		n1 := new(Node)
		n2 := new(Node)
		n1.next = n2
		n2.next = n1
		l.head = n1
		l.count = 2

		expected := 2
		got := l.LoopTypeDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("circular loop", func(t *testing.T) {
		l := getCircularLoopList()

		expected := 2
		got := l.LoopTypeDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("loop", func(t *testing.T) {
		l := getLoopList()

		expected := 1
		got := l.LoopTypeDetect()
		assert.Equal(t, expected, got)
	})
}
func TestList_LoopPointDetect(t *testing.T) {
	t.Run("no loop", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)
		l.AddTail(4)
		l.AddTail(5)

		var expected *Node
		got := l.LoopPointDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("circular loop. 2 nodes", func(t *testing.T) {
		l := NewList()
		n1 := new(Node)
		n2 := new(Node)
		n1.next = n2
		n2.next = n1
		l.head = n1
		l.count = 2

		expected := n2
		got := l.LoopPointDetect()
		assert.Equal(t, expected, got)
	})
	t.Run("circular loop", func(t *testing.T) {
		l := getCircularLoopList()

		expected := 1
		got := l.LoopPointDetect()
		assert.Equal(t, expected, got.value)
	})
	t.Run("loop", func(t *testing.T) {
		l := getLoopList()

		expected := 3
		got := l.LoopPointDetect()
		assert.Equal(t, expected, got.value)
	})
}

func getCircularLoopList() *List {
	l := NewList()
	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddTail(4)
	l.AddTail(5)
	l.AddNodeAtTail(Node{
		value: 6,
		next:  l.head,
	})
	return l
}
func getLoopList() *List {
	l := NewList()
	l.AddTail(1)
	l.AddTail(2)
	l.AddTail(3)
	l.AddNodeAtTail(Node{
		value: 4,
		next:  l.head,
	})
	l.AddHead(5)
	l.AddHead(6)
	return l
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
