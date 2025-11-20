package list

import (
	"cmp"
)

type Node[t cmp.Ordered] struct {
	Val  t
	prev *Node[t]
	next *Node[t]
}

// return Next Node
func (n *Node[t]) Next() *Node[t] {
	if n == nil {
		return nil
	}
	return n.next
}

// return Previous Node
func (n *Node[t]) Prev() *Node[t] {
	if n == nil {
		return nil
	}
	return n.prev
}

type List[t cmp.Ordered] struct {
	start *Node[t]
	last  *Node[t]
	size  int
}

// Creates new List with vals and returns pointer to it.
func New[t cmp.Ordered](vals ...t) *List[t] {
	l := new(List[t])
	return l.Add(vals...)
}

// Adds vals to List and returns the List pointer.
func (l *List[t]) Add(vals ...t) *List[t] {
	if l == nil {
		var o t
		l = New(o)
	}
	for _, v := range vals {
		l.add(v)
	}
	return l
}

// Returns pointer to Node from List which has Val v (or nil).
func (l *List[t]) Find(v t) *Node[t] {
	if l == nil {
		return nil
	}
	for n := l.start; n != nil; n = n.next {
		if n.Val == v {
			return n
		}
	}
	return nil
}

// Returns pointer to last Node from List (or nil).
func (l *List[t]) Last() *Node[t] {
	if l == nil {
		return nil
	}
	return l.last
}

// Removes vals from List and returns the List pointer.
func (l *List[t]) Remove(vals ...t) *List[t] {
	if l == nil {
		var o t
		l = New(o)
		return l
	}
	for _, v := range vals {
		l.remove(v)
	}
	return l
}

// Returns size of the List.
func (l *List[t]) Size() int {
	if l == nil {
		return 0
	}
	return l.size
}

// Returns pointer to first Node from List (or nil).
func (l *List[t]) Start() *Node[t] {
	if l == nil {
		var o t
		l = New(o)
	}
	return l.start
}

// Returns sorted slice of List's Values (for compatibility and ease of use).
func (l *List[t]) Values() []t {
	if l == nil {
		return nil
	}

	vals := make([]t, l.Size())
	i := 0
	for v := l.Start(); v != nil; v = v.Next() {
		vals[i] = v.Val
		i++
	}

	return vals
}

func (l *List[t]) remove(v t) *List[t] {
	c := l.Find(v)
	if c == nil {
		return l
	}

	p := c.Prev()
	n := c.Next()

	if p != nil {
		p.next = n
	}

	if n != nil {
		n.prev = p
	}
	l.size--
	return l
}

func insert_between[t cmp.Ordered](p, c, n *Node[t]) {
	c.next = n
	c.prev = n.prev
	n.prev = c
	if c.prev != nil {
		c.prev.next = c
	}
}

func (l *List[t]) add(s t) *List[t] {
	if l == nil {
		l = new(List[t])
	}

	if l.start == nil {
		l.start = &Node[t]{Val: s}
		l.last = l.start
		l.size++
		return l
	}

	c := &Node[t]{Val: s}

	n := l.start

	p := new(Node[t])
	for p = n; n != nil; n = n.next {
		if n.Val == s {
			return l // don't add again
		}

		if n.Val < s {
			p = n
			continue
		}

		if n.Val > s {
			// this means that to_add should be added before tw
			insert_between(p, c, n)

			if c.prev == nil {
				l.start = c
			}

			if c.next == nil {
				l.last = c
			}

			l.size++
			return l
		}
	}

	// wasn't added yet.
	// just add to end
	p.next = c
	c.prev = p
	l.last = c
	l.size++
	return l
}
