package mmh3

import (
	"crypto/rand"
	"io"
	mt "math/rand"
	"testing"
	"time"
)

func TestHash32(t *testing.T) {
	h := New32()
	io.WriteString(h, "h")
	io.WriteString(h, "e")
	io.WriteString(h, "l")
	io.WriteString(h, "l")
	io.WriteString(h, "o")
	if h.Sum32() != Hash32([]byte("hello")) {
		t.Fatal()
	}

	h.Reset()
	io.WriteString(h, "hello")
	if h.Sum32() != Hash32([]byte("hello")) {
		t.Fatal()
	}

	mt.Seed(time.Now().UnixNano())
	for i := 0; i < 1024; i++ {
		s := make([]byte, mt.Intn(2048))
		io.ReadFull(rand.Reader, s)
		h.Reset()
		h.Write(s)
		if h.Sum32() != Hash32(s) {
			t.Fatal()
		}
	}

	// for coverage
	if h.BlockSize() != 4 {
		t.Fatal()
	}
	if h.Size() != 4 {
		t.Fatal()
	}
	h.Reset()
	h.Write([]byte{'f', 'o', 'o'})
	h.Sum(nil)
	h.Reset()
	h.Write([]byte{'f', 'o'})
	h.Sum(nil)
}
