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
	// same test data from fuzz tests output files
	// fuzz tests keep adding files there, use any new file it creates
	a := []byte("\x8bl")
	b := []byte("\x8bl\xcef\x1ed\x05\x99\xbfu\xac\x1c")

	testCollisionTwoByteSlices(t, a, b)
}

func FuzzCollisionTwoByteSlices(f *testing.F) { f.Fuzz(testCollisionTwoByteSlices) }
