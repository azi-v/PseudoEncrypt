package pseudoencrypt

import (
	"testing"
	"time"
)

func TestPseudoEncrypt(t *testing.T) {
	v1 := time.Now().UnixNano()
	v2 := time.Now().Add(30).UnixNano()
	if PseudoEncrypt(v1) == PseudoEncrypt(v2) {
		t.Fatal("generate same value")
	}
}

func BenchmarkPseudoEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PseudoEncrypt(int64(i))
	}
}
func BenchmarkPseudoEncryptParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			PseudoEncrypt(time.Now().Unix())
		}
	})
}

func TestBoundedPseudoEncrypt(t *testing.T) {
	value := int64(1000000)
	min := int64(100000000000)
	max := int64(200000000000)
	res := BoundedPseudoEncrypt(value, max, min)

	if res > max || res < min {
		t.Fatal("generate value out of range")
	}

}

func BenchmarkBoundedPseudoEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BoundedPseudoEncrypt(int64(i), 1000000, 10003)
	}
}

func BenchmarkBoundedPseudoEncryptParallel(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			BoundedPseudoEncrypt(time.Now().Unix(), 1000000, 10003)
		}
	})
}
