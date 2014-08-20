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

func bench128(b *testing.B, bytes int) {
	bs := make([]byte, bytes)
	io.ReadFull(rand.Reader, bs)
	b.SetBytes(int64(bytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum128(bs)
	}
}

func BenchmarkHash128_1(b *testing.B)    { bench128(b, 1) }
func BenchmarkHash128_2(b *testing.B)    { bench128(b, 2) }
func BenchmarkHash128_4(b *testing.B)    { bench128(b, 4) }
func BenchmarkHash128_8(b *testing.B)    { bench128(b, 8) }
func BenchmarkHash128_16(b *testing.B)   { bench128(b, 16) }
func BenchmarkHash128_32(b *testing.B)   { bench128(b, 32) }
func BenchmarkHash128_64(b *testing.B)   { bench128(b, 64) }
func BenchmarkHash128_128(b *testing.B)  { bench128(b, 128) }
func BenchmarkHash128_256(b *testing.B)  { bench128(b, 256) }
func BenchmarkHash128_512(b *testing.B)  { bench128(b, 512) }
func BenchmarkHash128_1024(b *testing.B) { bench128(b, 1024) }
func BenchmarkHash128_2048(b *testing.B) { bench128(b, 2048) }
func BenchmarkHash128_4096(b *testing.B) { bench128(b, 4096) }
func BenchmarkHash128_8192(b *testing.B) { bench128(b, 8192) }
