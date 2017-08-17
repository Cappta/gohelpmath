package gohelpmath

import "crypto/sha1"

// Hash uses sha1 to hash and generate a specific output length
func Hash(value []byte, length int) []byte {
	hash := sha1.New()
	hash.Write(value)
	hash.Write(Int64ToBytes(int64(length)))

	checkSum := hash.Sum(nil)
	output := make([]byte, 0, length)
	for len(output) < length {
		hash.Write(checkSum)

		checkSum = hash.Sum(nil)
		output = append(output, checkSum...)
	}
	return output[:length]
}
