package mmh3

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	mt "math/rand"
	"testing"
	"time"
)

func TestHash128(t *testing.T) {
	h := New128()
	s := []byte("我能吞下玻璃而不伤身体")
	h.Write(s)
	if string(h.Sum(nil)) != string(Hash128(s)) {
		t.Fatal()
	}

	s, _ = hex.DecodeString(testString)
	h.Reset()
	h.Write(s)
	if string(h.Sum(nil)) != string(Hash128(s)) {
		t.Fatal()
	}

	h.Reset()
	for i := 0; i < 37; i++ {
		io.WriteString(h, "o")
	}
	if string(h.Sum(nil)) != string(Hash128(bytes.Repeat([]byte{'o'}, 37))) {
		t.Fatal()
	}

	mt.Seed(time.Now().UnixNano())
	for i := 0; i < 1024; i++ {
		s := make([]byte, mt.Intn(2048))
		io.ReadFull(rand.Reader, s)
		h.Reset()
		h.Write(s)
		if string(h.Sum(nil)) != string(Hash128(s)) {
			t.Fatal()
		}
	}

	// for coverage
	if h.BlockSize() != 16 {
		t.Fatal()
	}
	if h.Size() != 16 {
		t.Fatal()
	}
}

func BenchmarkH128(b *testing.B) {
	h := New128()
	s := bytes.Repeat([]byte("o"), 1024)
	b.SetBytes(int64(len(s)))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Reset()
		h.Write(s)
		h.Sum(nil)
	}
}
