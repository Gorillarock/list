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
	start   *Node[t]
	last    *Node[t]
	size    int
	indexed map[t]*Node[t]
}

// Creates new List with vals and returns pointer to it.
func New[t cmp.Ordered](vals ...t) *List[t] {
	l := new(List[t])
	return l.Add(vals...)
}

func NewIndexed[t cmp.Ordered](vals ...t) *List[t] {
	l := new(List[t])
	l.indexed = make(map[t]*Node[t])
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

	if n, ok := l.indexed[v]; ok {
		return n
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
	} else {
		l.start = n
	}

	if n != nil {
		n.prev = p
	} else {
		l.last = p
	}
	l.size--
	delete(l.indexed, v)
	return l
}

func insert_before_with_no_previous[t cmp.Ordered](c, n *Node[t]) {
	c.prev = nil
	c.next = n
	n.prev = c
}

func insert_after_with_no_next[t cmp.Ordered](c, p *Node[t]) {
	c.next = nil
	c.prev = p
	p.next = c
}

func insert_between[t cmp.Ordered](p, c, n *Node[t]) {
	c.next = n
	c.prev = n.prev
	n.prev = c
	if c.prev != nil {
		c.prev.next = c
	}
}

func (l *List[t]) add(s t) *Node[t] {
	if l == nil {
		l = new(List[t])
	}

	c := &Node[t]{Val: s}
	if l.start == nil {
		l.start = c
		l.last = l.start
		l.size++
		l.index(c)
		return c
	}

	if n, ok := l.indexed[s]; ok {
		return n
	}

	n := l.start

	p := new(Node[t])
	for p = n; n != nil; n = n.next {
		if n.Val == s {
			return n // don't add again
		}

		if n.Val < s {
			p = n
			continue
		}

		if n.Val > s {
			// this means that to_add should be added before tw
			switch {
			case p.Val < n.Val:
				insert_between(p, c, n)
			case p.Val == n.Val:
				insert_before_with_no_previous(c, n)
			}

			if c.prev == nil {
				l.start = c
			}

			if c.next == nil {
				l.last = c
			}

			l.size++
			l.index(c)
			return c
		}
	}

	// wasn't added yet.
	// just add to end
	p.next = c
	c.prev = p
	l.last = c
	l.size++
	l.index(c)
	return c
}

func (l *List[t]) index(n *Node[t]) {
	if l.indexed != nil {
		l.indexed[n.Val] = n
	}
}
