package splaytree

type binaryNode struct {
	left, right *binaryNode
	item        Item
}

type splayTree struct {
	root *binaryNode
}

func NewSplayTree() Tree {
	return &splayTree{}
}

func (st *splayTree) Get(key Item) {
	panic("not implemented")
}

func (st *splayTree) Has(key Item) bool {
	panic("not implemented")
}

func (st *splayTree) ReplaceOrInsert(item Item) Item {
	panic("not implemented")
}

func (st *splayTree) Delete(item Item) Item {
	panic("not implemented")
}

func (st *splayTree) DeleteMin() Item {
	panic("not implemented")
}

func (st *splayTree) DeleteMax() Item {
	panic("not implemented")
}

func (st *splayTree) Len() int {
	panic("not implemented")
}

func (st *splayTree) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	panic("not implemented")
}

func (st *splayTree) AscendLessThan(pivot Item, iterator ItemIterator) {
	panic("not implemented")
}

func (st *splayTree) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	panic("not implemented")
}
