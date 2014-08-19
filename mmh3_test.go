package mmh3

import (
	"encoding/hex"
	"fmt"
	"testing"
)

/* 1.4
func generateTestCases() {
	mr.Seed(time.Now().UnixNano())
	buf := make([]byte, 256)
	var bs []byte
	var coverage float64
	for {
		length := mr.Intn(256)
		bs = buf[:length]
		_, err := rand.Read(bs)
		if err != nil {
			panic(err)
		}
		Hash32(bs)
		Hash128(bs)
		if c := testing.Coverage(); c > coverage {
			fmt.Printf("%x %f\n", bs, c)
			coverage = c
		}
		if coverage == 1 {
			break
		}
	}
}
*/

const testString = `1bb2cc0687b19ff4863639208620f85ed175d856f461b18566c06900e628afd1079ed2bd859fac5d0e9c6e6c3761e50352452695f515c7b22b695484fdd168b8fee6176f68b3e8aca2460cc45eca3bd45fdc90e8a810118c97af69b0b65cb3a7e1aaf4402e837d114470d636d5aea7dda4c88576f560face5c181466546da7084ac93700a392ead1404e13ffff0d71`

func TestAll(t *testing.T) {
	s := []byte("hello")
	if Hash32(s) != 0x248bfa47 {
		t.Fatal()
	}
	if fmt.Sprintf("%x", Hash128(s)) != "029bbd41b3a7d8cb191dae486a901e5b" {
		t.Fatal()
	}
	s = []byte("Winter is coming")
	if Hash32(s) != 0x43617e8f {
		t.Fatal()
	}
	if fmt.Sprintf("%x", Hash128(s)) != "95eddc615d3b376c13fb0b0cead849c5" {
		t.Fatal()
	}

	if Hash32([]byte{}) != 0 {
		t.Fatal()
	}
	if fmt.Sprintf("%x", Hash128([]byte{})) != "00000000000000000000000000000000" {
		t.Fatal()
	}

	// for coverage
	s, err := hex.DecodeString(testString)
	if err != nil {
		t.Fatal()
	}
	if Hash32(s) != 160380997 {
		t.Fatal()
	}
	if fmt.Sprintf("%x", Hash128(s)) != "cf046caad6f7ee280019d651dd2a9635" {
		t.Fatal()
	}
}

func Benchmark32(b *testing.B) {
	bs := []byte("hello")
	for i := 0; i < b.N; i++ {
		Hash32(bs)
	}
}

func Benchmark128(b *testing.B) {
	s, err := hex.DecodeString(testString)
	if err != nil {
		b.Fatal()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hash128(s)
	}
}
