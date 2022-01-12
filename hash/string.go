package hash

// HashString offers a way to hash a string of arbitrary data,
// but does so 8 bytes at a time.
func HashString(s string) Hash {
	h := fnvOffset

	for len(s) >= 8 {
		var l [8]byte
		copy(l[:], s[0:8])
		h = doHash(h, l)
		s = s[8:]
	}

	if len(s) >= 4 {
		h = (h ^ Hash(s[0])) * fnvPrime
		h = (h ^ Hash(s[1])) * fnvPrime
		h = (h ^ Hash(s[2])) * fnvPrime
		h = (h ^ Hash(s[3])) * fnvPrime
		s = s[4:]
	}

	if len(s) >= 2 {
		h = (h ^ Hash(s[0])) * fnvPrime
		h = (h ^ Hash(s[1])) * fnvPrime
		s = s[2:]
	}

	if len(s) > 0 {
		h = (h ^ Hash(s[0])) * fnvPrime
	}

	return h
}
