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

func bench32(b *testing.B, bytes int) {
	bs := make([]byte, bytes)
	io.ReadFull(rand.Reader, bs)
	b.SetBytes(int64(bytes))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sum32(bs)
	}
}

func BenchmarkHash32_1(b *testing.B)    { bench32(b, 1) }
func BenchmarkHash32_2(b *testing.B)    { bench32(b, 2) }
func BenchmarkHash32_4(b *testing.B)    { bench32(b, 4) }
func BenchmarkHash32_8(b *testing.B)    { bench32(b, 8) }
func BenchmarkHash32_16(b *testing.B)   { bench32(b, 16) }
func BenchmarkHash32_32(b *testing.B)   { bench32(b, 32) }
func BenchmarkHash32_64(b *testing.B)   { bench32(b, 64) }
func BenchmarkHash32_128(b *testing.B)  { bench32(b, 128) }
func BenchmarkHash32_256(b *testing.B)  { bench32(b, 256) }
func BenchmarkHash32_512(b *testing.B)  { bench32(b, 512) }
func BenchmarkHash32_1024(b *testing.B) { bench32(b, 1024) }
func BenchmarkHash32_2048(b *testing.B) { bench32(b, 2048) }
func BenchmarkHash32_4096(b *testing.B) { bench32(b, 4096) }
func BenchmarkHash32_8192(b *testing.B) { bench32(b, 8192) }
