package splaytree

// Item represents a single object in the tree.
type Item interface {
	// Less tests whether the current item is less than the given argument.
	//
	// This must provide a strict weak ordering.
	// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
	// hold one of either a or b in the tree).
	Less(than Item) bool
}

// ItemIterator allows callers of Ascend* to iterate in-order over portions of
// the tree.  When this function returns false, iteration will stop and the
// associated Ascend* function will immediately return.
type ItemIterator func(i Item) bool

type Tree interface {
	// Get looks for the key item in the tree, returning it.  It returns nil if
	// unable to find that item.
	Get(key Item) Item

	// Has returns true if the given key is in the tree.
	Has(key Item) bool

	// ReplaceOrInsert adds the given item to the tree.  If an item in the tree
	// already equals the given one, it is removed from the tree and returned.
	// Otherwise, nil is returned.
	//
	// nil cannot be added to the tree (will panic).
	ReplaceOrInsert(item Item) Item

	// Delete removes an item equal to the passed in item from the tree, returning
	// it.  If no such item exists, returns nil.
	Delete(item Item) Item

	// DeleteMin removes the smallest item in the tree and returns it.
	// If no such item exists, returns nil.
	DeleteMin() Item

	// DeleteMax removes the largest item in the tree and returns it.
	// If no such item exists, returns nil.
	DeleteMax() Item

	// Len returns the number of items currently in the tree.
	Len() int

	// AscendRange calls the iterator for every value in the tree within the range
	// [greaterOrEqual, lessThan), until iterator returns false.
	AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator)

	// AscendLessThan calls the iterator for every value in the tree within the range
	// [first, pivot), until iterator returns false.
	AscendLessThan(pivot Item, iterator ItemIterator)

	// AscendGreaterOrEqual calls the iterator for every value in the tree within
	// the range [pivot, last], until iterator returns false.
	AscendGreaterOrEqual(pivot Item, iterator ItemIterator)
}
