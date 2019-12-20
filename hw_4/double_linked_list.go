package doublelinkedlist

type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

type List struct {
	head *Item
	tail *Item
	size int // ?? uint
}

func (l List) Len() int {
	return l.size
}

func (l List) First() Item {
	if l.head != nil {
		return *l.head
	}

	return Item{}
}

func (l List) Last() Item {
	if l.tail != nil {
		return *l.tail
	}

	return Item{}
}

func (l *List) PushFront(v interface{}) {

}

func (l *List) PushBack(v interface{}) {

}

func (l *List) Remove(i Item) {

}

func (i Item) Value() interface{} {
	return i.value
}

func (i Item) Next() *Item {
	return i.next
}

func (i Item) Prev() *Item {
	return i.prev
}
