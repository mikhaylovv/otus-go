package doublelinkedlist

//Item for double linked list, can contains any element
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

//List - Double linked list structure
//With fast access to first and last elements.
type List struct {
	head *Item
	tail *Item
	size int // ?? uint
}

//Len - returns the length of list. O(1)
func (l List) Len() int {
	return l.size
}

//First - returns the first element of the list if it exist.
//Otherwise return empty Item. O(1)
func (l List) First() Item {
	if l.head != nil {
		return *l.head
	}

	return Item{}
}

//Last - returns the last element of the list if it exist.
//Otherwise return empty Item. O(1)
func (l List) Last() Item {
	if l.tail != nil {
		return *l.tail
	}

	return Item{}
}

//PushFront - pushes element to the front of the double linked list, if list exist. O(1)
func (l *List) PushFront(v interface{}) {
	if l == nil {
		return
	}

	el := Item {
		value: v,
		prev: nil,
		next: l.head,
	}

	if l.head != nil {
		l.head.prev = &el
	}

	l.head = &el

	if l.tail == nil {
		l.tail = &el
	}

	l.size++
}

//PushBack - pushes element to the end of the double linked list, if list exist. O(1)
func (l *List) PushBack(v interface{}) {
	if l == nil {
		return
	}

	el := Item {
		value: v,
		prev: l.tail,
		next: nil,
	}

	if l.tail != nil {
		l.tail.next = &el
	}

	l.tail = &el

	if l.head == nil {
		l.head = &el
	}

	l.size++
}

//Remove - removes element from list.
//This method implemented on trust, if you put as argument element from different list, it may corrupt your objects.
func (l *List) Remove(i Item) {
	if l == nil {
		return
	}

	if l.Len() == 0 {
		return
	}

	defer func() { l.size-- }()

	// head
	if i.Prev() == nil {
		l.head = i.Next()
		return
	}

	// tail
	if i.Next() == nil {
		l.tail = i.Prev()
		return
	}

	i.Prev().next = i.Next()
	i.Next().prev = i.Prev()
}

//Value - returns the copy of Item object.
func (i Item) Value() interface{} {
	return i.value
}

//Next - returns the copy of next obj pointer
func (i Item) Next() *Item {
	return i.next
}

//Prev - returns the copy of prev obj pointer
func (i Item) Prev() *Item {
	return i.prev
}
