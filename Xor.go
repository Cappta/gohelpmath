package gohelpmath

// Xor will apply XOR to each byte of a add b and return the result
func Xor(a, b []byte) []byte {
	if len(a) > len(b) {
		return Xor(b, a)
	}

	aLength := len(a)
	bLength := len(b)

	dst := make([]byte, bLength)
	for i := 0; i < aLength; i++ {
		dst[i] = a[i] ^ b[i]
	}
	for i := aLength; i < bLength; i++ {
		dst[i] = b[i]
	}
	return dst
}
