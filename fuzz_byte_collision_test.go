package main_test

import (
	"bytes"
	"testing"
)

func testCollisionTwoByteSlices(t *testing.T, a, b []byte) {
	aBefore := make([]byte, len(a))
	copy(aBefore, a)

	bBefore := make([]byte, len(b))
	copy(bBefore, b)

	aAfter := append(a, b...)

	if !bytes.Equal(aBefore, aAfter[:len(aBefore)]) {
		t.Error(aBefore, aAfter, a)
	}
	if !bytes.Equal(bBefore, b) {
		// WARNING: b nor bBefore was ever modified!
		t.Error(aBefore, bBefore, a, b, aAfter)
	}
}

func TestCollisionTwoByteSlices(t *testing.T) {
	a := []byte{1, 128, 163}
	b := []byte{1, 128, 163, 1, 0, 255}

	testCollisionTwoByteSlices(t, a, b)
}

func FuzzCollisionTwoByteSlices(f *testing.F) { f.Fuzz(testCollisionTwoByteSlices) }
