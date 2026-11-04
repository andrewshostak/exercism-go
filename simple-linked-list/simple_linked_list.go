package simplelinkedlist

import "errors"

type List struct {
	length int
	first  *Element
}

type Element struct {
	value int
	next  *Element
}

func New(elements []int) *List {
	l := List{}
	for _, v := range elements {
		l.Push(v)
	}

	return &l
}

func (l *List) Size() int {
	return l.length
}

func (l *List) Push(element int) {
	l.first = &Element{
		value: element,
		next:  l.first,
	}
	l.length++
}

func (l *List) Pop() (int, error) {
	if l.length == 0 {
		return 0, errors.New("list is empty")
	}

	value := l.first.value
	l.first = l.first.next
	l.length--

	return value, nil
}

func (l *List) Peek() (int, error) {
	if l.length == 0 {
		return 0, errors.New("list is empty")
	}

	return l.first.value, nil
}

func (l *List) Array() []int {
	arr := make([]int, l.length)

	current := l.first
	for i := l.length - 1; current != nil; i-- {
		arr[i] = current.value
		current = current.next
	}

	return arr
}

func (l *List) Reverse() *List {
	arr := make([]int, 0, l.length)
	current := l.first
	for current != nil {
		arr = append(arr, current.value)
		current = current.next
	}

	return New(arr)
}
