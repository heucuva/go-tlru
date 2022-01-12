package tlru

import (
	"github.com/heucuva/go-tlru/hash"
)

// SetString sets a value into the cache based on a string key
func (t *TLRU) SetString(key string, value interface{}) {
	h := hash.HashString(key)
	t.mapset(h, key, value)
}

// GetString gets a value from the cache based on a string key
func (t *TLRU) GetString(key string) (interface{}, bool) {
	h := hash.HashString(key)
	return t.mapget(h)
}

// DeleteString removes an entry from the cache based on a string key
func (t *TLRU) DeleteString(key string) {
	h := hash.HashString(key)
	t.mapremove(h)
}
