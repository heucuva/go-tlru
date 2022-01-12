package tlru

import (
	"time"

	"github.com/heucuva/go-tlru/hash"
)

type node struct {
	value interface{}
	key   interface{}
	hash  hash.Hash
	hit   time.Time
	left  *node
	right *node
}

func (n *node) touch() {
	n.hit = time.Now()
}

func (n *node) age() time.Duration {
	return time.Since(n.hit)
}
