package mmh3

import (
	"reflect"
	"unsafe"
)

var (
	Hash32x86  = Hash32
	Hash128x64 = Hash128
)

func Hash32(key []byte) uint32 {
	length := len(key)
	if length == 0 {
		return 0
	}
	var c1, c2 uint32 = 0xcc9e2d51, 0x1b873593
	nblocks := length / 4
	var h, k uint32
	var blocks []uint32
	keyHeader := (*reflect.SliceHeader)(unsafe.Pointer(&key))
	blocksHeader := (*reflect.SliceHeader)(unsafe.Pointer(&blocks))
	blocksHeader.Data = keyHeader.Data
	blocksHeader.Len = length
	blocksHeader.Cap = length
	for i := 0; i < nblocks; i++ {
		k = blocks[i] // assuming no endian swapping
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
		h = (h << 13) | (h >> (32 - 13))
		h = (h * 5) + 0xe6546b64
	}
	k = 0
	tailIndex := nblocks * 4
	switch length & 3 {
	case 3:
		k ^= uint32(key[tailIndex+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[tailIndex+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[tailIndex])
		k *= c1
		k = (k << 15) | (k >> (32 - 15))
		k *= c2
		h ^= k
	}
	h ^= uint32(length)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}

func Hash128(key []byte) []byte {
	length := len(key)
	ret := make([]byte, 16)
	if length == 0 {
		return ret
	}
	var c1, c2 uint64 = 0x87c37b91114253d5, 0x4cf5ad432745937f
	nblocks := length / 16
	var h1, h2, k1, k2 uint64
	var uint64s []uint64
	keyHeader := (*reflect.SliceHeader)(unsafe.Pointer(&key))
	uint64sHeader := (*reflect.SliceHeader)(unsafe.Pointer(&uint64s))
	uint64sHeader.Data = keyHeader.Data
	uint64sHeader.Len = nblocks * 2
	uint64sHeader.Cap = uint64sHeader.Len
	for i := 0; i < nblocks*2; i += 2 {
		k1 = uint64s[i] // assuming no endian swapping
		k2 = uint64s[i+1]
		k1 *= c1
		k1 = (k1 << 31) | (k1 >> (64 - 31))
		k1 *= c2
		h1 ^= k1
		h1 = (h1 << 27) | (h1 >> (64 - 27))
		h1 += h2
		h1 = h1*5 + 0x52dce729
		k2 *= c2
		k2 = (k2 << 33) | (k2 >> (64 - 33))
		k2 *= c1
		h2 ^= k2
		h2 = (h2 << 31) | (h2 >> (64 - 31))
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}
	k1, k2 = 0, 0
	tailIndex := nblocks * 16
	switch length & 15 {
	case 15:
		k2 ^= uint64(key[tailIndex+14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(key[tailIndex+13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(key[tailIndex+12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(key[tailIndex+11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(key[tailIndex+10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(key[tailIndex+9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(key[tailIndex+8])
		k2 *= c2
		k2 = (k2 << 33) | (k2 >> (64 - 33))
		k2 *= c1
		h2 ^= k2
		fallthrough
	case 8:
		k1 ^= uint64(key[tailIndex+7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(key[tailIndex+6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(key[tailIndex+5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(key[tailIndex+4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(key[tailIndex+3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(key[tailIndex+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(key[tailIndex+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(key[tailIndex])
		k1 *= c1
		k1 = (k1 << 31) | (k1 >> (64 - 31))
		k1 *= c2
		h1 ^= k1
	}
	h1 ^= uint64(length)
	h2 ^= uint64(length)
	h1 += h2
	h2 += h1
	h1 ^= h1 >> 33
	h1 *= 0xff51afd7ed558ccd
	h1 ^= h1 >> 33
	h1 *= 0xc4ceb9fe1a85ec53
	h1 ^= h1 >> 33
	h2 ^= h2 >> 33
	h2 *= 0xff51afd7ed558ccd
	h2 ^= h2 >> 33
	h2 *= 0xc4ceb9fe1a85ec53
	h2 ^= h2 >> 33
	h1 += h2
	h2 += h1
	retHeader := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
	var tuple []uint64
	tupleHeader := (*reflect.SliceHeader)(unsafe.Pointer(&tuple))
	tupleHeader.Data = retHeader.Data
	tupleHeader.Len = 2
	tupleHeader.Cap = 2
	tuple[0] = h1
	tuple[1] = h2
	return ret
}
