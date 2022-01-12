package tlru

import (
	"github.com/heucuva/go-tlru/hash"
)

// SetInt sets a value into the cache based on an integer key
func (t *TLRU) SetInt(key int, value interface{}) {
	h := hash.HashUint64(uint64(key))
	t.mapset(h, key, value)
}

// GetInt gets a value from the cache based on an integer key
func (t *TLRU) GetInt(key int) (interface{}, bool) {
	h := hash.HashUint64(uint64(key))
	return t.mapget(h)
}

// DeleteInt removes an entry from the cache based on an integer key
func (t *TLRU) DeleteInt(key int) {
	h := hash.HashUint64(uint64(key))
	t.mapremove(h)
}
