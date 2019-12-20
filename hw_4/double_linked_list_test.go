package doublelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemValueSimple(t *testing.T) {
	as := assert.New(t)

	{
		el := Item{}
		as.Empty(el.Value())
	}

	el := Item{
		value: 10,
		prev:  nil,
		next:  nil,
	}

	as.EqualValues(10, el.Value())
}

func TestItemNextSimple(t *testing.T) {
	as := assert.New(t)
	{
		el := Item{}
		as.Empty(el.Next())
	}

	el := Item{
		value: 10,
		prev:  nil,
		next:  nil,
	}

	sec := Item{
		value: 20,
		prev:  nil,
		next:  &el,
	}

	as.EqualValues(10, sec.Next().Value())
}

func TestItemPrevSimple(t *testing.T) {
	as := assert.New(t)
	{
		el := Item{}
		as.Empty(el.Prev())
	}

	el := Item{
		value: 10,
		prev:  nil,
		next:  nil,
	}

	sec := Item{
		value: 20,
		prev:  &el,
		next:  nil,
	}

	as.EqualValues(10, sec.Prev().Value())
}

func TestListEmptyLen(t *testing.T) {
	as := assert.New(t)

	list := List{}
	as.EqualValues(0, list.Len())
}

func setUpListTest(t *testing.T) (*assert.Assertions, List) {
	as := assert.New(t)

	list := List{}
	for i := 0; i < 10; i++ {
		list.PushBack(i)
	}

	return as, list
}

func TestListSimpleLen(t *testing.T) {
	as, list := setUpListTest(t)

	as.EqualValues(10, list.Len())
}

func TestListSimpleFirst(t *testing.T) {
	as, list := setUpListTest(t)

	as.EqualValues(0, list.First().value)
}

func TestListSimpleLast(t *testing.T) {
	as, list := setUpListTest(t)

	as.EqualValues(9, list.Last().value)
}

func TestListSimplePushFront(t *testing.T) {
	as, list := setUpListTest(t)

	list.PushFront(20)
	as.EqualValues(20, list.First().value)
	as.EqualValues(11, list.Len())
}

func TestListSimplePushBack(t *testing.T) {
	as, list := setUpListTest(t)

	list.PushBack(30)
	as.EqualValues(30, list.Last().value)
	as.EqualValues(11, list.Len())
}

func TestListRemoveSimle(t *testing.T) {
	as, list := setUpListTest(t)

	list.Remove(list.First())
	as.EqualValues(1, list.First().value)
	as.EqualValues(9, list.Len())

	list.Remove(list.Last())
	as.EqualValues(8, list.Last().value)
	as.EqualValues(8, list.Len())

	list.Remove(*list.First().Next())
	as.EqualValues(7, list.Len())
	as.EqualValues(1, list.First().value)
	as.EqualValues(8, list.Last().value)
}
