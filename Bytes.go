package gohelpmath

// Int32ToBytes converts a 32-bit integer to a byte slice
func Int32ToBytes(value int) []byte {
	return Uint32ToBytes(uint(value))
}

// BytesToInt32 converts a byte slice to a 32-bit integer
func BytesToInt32(value []byte) int {
	return int(int32(BytesToUint32(value)))
}

// Int64ToBytes converts a 64-bit integer to a byte slice
func Int64ToBytes(value int64) []byte {
	return Uint64ToBytes(uint64(value))
}

// BytesToInt64 converts a byte slice to a 64-bit integer
func BytesToInt64(value []byte) int64 {
	return int64(BytesToUint64(value))
}

// Uint32ToBytes converts a 32-bit unsigned integer to a byte slice
func Uint32ToBytes(value uint) []byte {
	return []byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

// BytesToUint32 converts a byte slice to a 32-bit unsigned integer
func BytesToUint32(value []byte) uint {
	return uint(value[0])<<24 | uint(value[1])<<16 | uint(value[2])<<8 | uint(value[3])
}

// Uint64ToBytes converts a 64-bit unsigned integer to a byte slice
func Uint64ToBytes(value uint64) []byte {
	return []byte{byte(value >> 56), byte(value >> 48), byte(value >> 40), byte(value >> 32),
		byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)}
}

// BytesToUint64 converts a byte slice to a 64-bit unsigned integer
func BytesToUint64(value []byte) uint64 {
	return uint64(value[0])<<56 | uint64(value[1])<<48 | uint64(value[2])<<40 | uint64(value[3])<<32 | uint64(value[4])<<24 | uint64(value[5])<<16 | uint64(value[6])<<8 | uint64(value[7])
}

// BytesTrimLeadingZero will return the provided slice's reference after all leading zeroes
func BytesTrimLeadingZero(value []byte) []byte {
	for i := 0; i < len(value); i++ {
		if value[i] != 0 {
			return value[i:]
		}
	}
	return value[len(value):]
}
