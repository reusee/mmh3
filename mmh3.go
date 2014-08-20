package mmh3

var (
	Hash32x86  = Sum32
	Hash128x64 = Sum128

	//for backward compatible
	Hash32  = Sum32
	Hash128 = Sum128
)

func Sum32(key []byte) uint32 {
	h := New32()
	h.Write(key)
	return h.Sum32()
}

func Sum128(key []byte) []byte {
	h := New128()
	h.Write(key)
	return h.Sum(nil)
}
