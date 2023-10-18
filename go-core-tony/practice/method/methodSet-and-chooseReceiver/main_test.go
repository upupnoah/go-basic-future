package main

import "testing"

// go test
func TestDumpMethodSet(t *testing.T) {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t1 T
	dumpMethodSet(t1)
	dumpMethodSet(&t1)

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}
