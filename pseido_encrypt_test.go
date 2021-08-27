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

func TestBoundedPseudoEncrypt(t *testing.T) {
	value := int64(1000000)
	min := int64(100000000000)
	max := int64(200000000000)
	res := BoundedPseudoEncrypt(value, max, min)

	if res > max || res < min {
		t.Fatal("generate value out of range")
	}

}
