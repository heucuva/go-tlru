package tlru

import (
	"time"

	"github.com/heucuva/go-tlru/hash"
)

// TLRU is a genericized Time-aware Least Recently Used cache
type TLRU struct {
	// MaxSize if > 0 will limit the number of entries the map TLRU will contain
	// if == 0, then no limit
	MaxSize int

	m map[hash.Hash]*node
	h *node
	t *node
}

// Expire will remove any entries older than maxage
func (t *TLRU) Expire(maxage time.Duration) {
	t.listiterate(func(n *node) bool {
		if n.age() < maxage {
			return false
		}

		t.mapremove(n.hash)
		return true
	})
}

// Iterate will call the function fn with each entry in age order from oldest to newest.
// Age is a metric of least-recent use
func (t *TLRU) Iterate(fn func(key, value interface{}, age time.Duration) bool) {
	t.listiterate(func(n *node) bool {
		return fn(n.key, n.value, n.age())
	})
}

// Len returns the number of entries in the cache
func (t *TLRU) Len() int {
	return len(t.m)
}

//// helper functions ////

func (t *TLRU) ensure() {
	if t.m == nil {
		t.m = make(map[hash.Hash]*node)
	}
}

func (t *TLRU) mapset(h hash.Hash, key, value interface{}) {
	t.ensure()

	n, found := t.m[h]
	if n == nil {
		n = &node{
			hash: h,
		}
	}
	if !found {
		t.m[h] = n
	}

	n.key = key
	n.value = value
	n.touch()
	t.listpush(n)
}

func (t *TLRU) mapget(h hash.Hash) (interface{}, bool) {
	if t.m == nil {
		return nil, false
	}

	n, found := t.m[h]
	if !found {
		return nil, false
	}

	n.touch()
	t.listpush(n)

	return n.value, true
}

func (t *TLRU) mapremove(h hash.Hash) {
	if t.m == nil {
		return
	}

	n, found := t.m[h]
	if !found {
		return
	}

	delete(t.m, h)
	t.listremove(n)
}

func (t *TLRU) listpush(n *node) {
	t.listremove(n)

	tail := t.t
	t.t = n
	if tail == nil {
		t.h = n
		return
	}

	n.left = tail
	n.right = nil
	tail.right = n

	for t.h != nil && t.MaxSize > 0 && len(t.m) > t.MaxSize {
		delete(t.m, t.h.hash)
		t.listremove(t.h)
	}
}

func (t *TLRU) listremove(n *node) {
	prev := n.left
	next := n.right
	if prev != nil {
		prev.right = n.right
	}
	if next != nil {
		next.left = n.left
	}
	n.right = nil
	n.left = nil

	if t.t == n {
		t.t = prev
	}

	if t.h == n {
		t.h = next
	}
}

func (t *TLRU) listiterate(fn func(n *node) bool) {
	n := t.h
	for n != nil {
		next := n.right
		if !fn(n) {
			return
		}
		n = next
	}
}
