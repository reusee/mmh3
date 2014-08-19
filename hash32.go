package mmh3

import (
	"hash"
	"reflect"
	"unsafe"
)

const (
	c1_32 uint32 = 0xcc9e2d51
	c2_32 uint32 = 0x1b873593
)

type hash32 struct {
	hash uint32
	tail []byte
	size uint32
}

func New32() hash.Hash32 {
	return new(hash32)
}

func (h *hash32) BlockSize() int {
	return 4
}

func (h *hash32) Reset() {
	h.hash = 0
	h.tail = nil
	h.size = 0
}

func (h *hash32) Size() int {
	return 4
}

func (h *hash32) Sum(in []byte) []byte {
	var k uint32
	if h.tail != nil {
		switch len(h.tail) {
		case 3:
			k ^= uint32(h.tail[2]) << 16
			fallthrough
		case 2:
			k ^= uint32(h.tail[1]) << 8
			fallthrough
		case 1:
			k ^= uint32(h.tail[0])
			k *= c1_32
			k = (k << 15) | (k >> (32 - 15))
			k *= c2_32
			h.hash ^= k
		}
	}
	h.hash ^= h.size
	h.hash ^= h.hash >> 16
	h.hash *= 0x85ebca6b
	h.hash ^= h.hash >> 13
	h.hash *= 0xc2b2ae35
	h.hash ^= h.hash >> 16
	return append(in, byte(h.hash), byte(h.hash>>8), byte(h.hash>>16), byte(h.hash>>24))
}

func (h *hash32) Sum32() uint32 {
	res := h.Sum(nil)
	return uint32(res[0]) + uint32(res[1])<<8 + uint32(res[2])<<16 + uint32(res[3])<<24
}

func (h *hash32) Write(key []byte) (n int, err error) {
	n = len(key)
	h.size += uint32(n)

	if h.tail != nil {
		for len(key) > 0 && len(h.tail) < 4 {
			h.tail = append(h.tail, key[0])
			key = key[1:]
		}
		if len(h.tail) == 4 { // a full block
			h.sumBlock(uint32(h.tail[0]) +
				uint32(h.tail[1])<<8 +
				uint32(h.tail[2])<<16 +
				uint32(h.tail[3])<<24)
			h.tail = nil
		}
	}

	length := len(key)
	nblocks := length / 4
	if nblocks > 0 {
		var blocks []uint32
		keyHeader := (*reflect.SliceHeader)(unsafe.Pointer(&key))
		blocksHeader := (*reflect.SliceHeader)(unsafe.Pointer(&blocks))
		blocksHeader.Data = keyHeader.Data
		blocksHeader.Len = nblocks
		blocksHeader.Cap = nblocks
		for _, k := range blocks {
			h.sumBlock(k)
		}
	}

	if length%4 != 0 {
		h.tail = key[nblocks*4 : length]
	}

	return
}

func (h *hash32) sumBlock(k uint32) {
	k *= c1_32
	k = (k << 15) | (k >> (32 - 15))
	k *= c2_32
	h.hash ^= k
	h.hash = (h.hash << 13) | (h.hash >> (32 - 13))
	h.hash = (h.hash * 5) + 0xe6546b64
}
