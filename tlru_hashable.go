package tlru

import (
	"github.com/heucuva/go-tlru/hash"
)

// SetHashable sets a value into the cache based on a user-defined hashable key
func (t *TLRU) SetHashable(key hash.Hashable, value interface{}) {
	h := hash.HashHashable(key)
	t.mapset(h, key, value)
}

// GetHashable gets a value from the cache based on a user-defined hashable key
func (t *TLRU) GetHashable(key hash.Hashable) (interface{}, bool) {
	h := hash.HashHashable(key)

	return t.mapget(h)
}

// DeleteHashable removes an entry from the cache based on a user-defined hashable key
func (t *TLRU) DeleteHashable(key hash.Hashable) {
	h := hash.HashHashable(key)
	t.mapremove(h)
}
