package main

import (
	"testing"

	"go-pkgs/crypto/met"
)

var src = &met.Source{
	Str: "Hello, World!",
}

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Md5Hash()
	}
}

func BenchmarkSha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Sha1Hash()
	}
}

func BenchmarkMurmurhassh32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Murmur32()
	}
}

func BenchmarkMurmurhassh64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.Murmur64()
	}
}
