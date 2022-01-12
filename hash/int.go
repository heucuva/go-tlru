package hash

import (
	"encoding/binary"
)

func doHash(h Hash, v [8]byte) Hash {
	h = (h ^ Hash(v[0])) * fnvPrime
	h = (h ^ Hash(v[1])) * fnvPrime
	h = (h ^ Hash(v[2])) * fnvPrime
	h = (h ^ Hash(v[3])) * fnvPrime
	h = (h ^ Hash(v[4])) * fnvPrime
	h = (h ^ Hash(v[5])) * fnvPrime
	h = (h ^ Hash(v[6])) * fnvPrime
	h = (h ^ Hash(v[7])) * fnvPrime
	return h
}

// HashUint64 offers a hashing of 8 bytes of data in uint64 form
func HashUint64(v uint64) Hash {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], v)
	return doHash(fnvOffset, b)
}
