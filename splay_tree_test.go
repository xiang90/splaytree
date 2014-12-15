package splaytree

import "testing"

func TestSplayTree(t *testing.T) {
	tree := &splayTree{}
	n := 40000
	gap := 307
	for i := gap; i != 0; i = (i + gap) % n {
		tree.insert(Int(i))
	}

	for i := gap; i != 0; i = (i + gap) % n {
		g := tree.Get(Int(i))
		if g != Int(i) {
			t.Errorf("g = %v, want %v", g, i)
		}
	}
}

// Int implements the Item interface for integers.
type Int int

// Less returns true if int(a) < int(b).
func (a Int) Less(b Item) bool {
	return a < b.(Int)
}
