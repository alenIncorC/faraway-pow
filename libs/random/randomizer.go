package random

var charSet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomNumber(seed uint64) uint64 {
	seed ^= seed << 21
	seed ^= seed >> 35
	seed ^= seed << 4
	return seed
}

func RandomString(str []byte, offset int, seed uint64) uint64 {
	for i := offset; i < len(str); i++ {
		seed = RandomNumber(seed)
		str[i] = charSet[seed%62]
	}

	return seed
}
