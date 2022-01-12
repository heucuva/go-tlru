package hash

// Hash is a simple identifier with unlikely collision
type Hash uint64

// Hashable defines an interface for something hashing itself
type Hashable interface {
	Hash() Hash
}

// Built-in hashing based on FNV
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function

// FNV hash function values
const (
	fnvOffset = Hash(0xcbf29ce484222325)
	fnvPrime  = Hash(0x00000100000001B3)
)

// HashHashable provides an easy way to generate a hash from a Hashable interface,
// with 0 being the result of a nil hashing.
func HashHashable(v Hashable) Hash {
	if v == nil {
		return 0
	}

	return v.Hash()
}
