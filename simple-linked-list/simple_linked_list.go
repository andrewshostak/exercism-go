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
	var next *Element
	for i := len(elements) - 1; i >= 0; i-- {
		e := &Element{value: elements[i], next: next}
		next = e
	}

	return &List{length: len(elements), first: next}
}

func (l *List) Size() int {
	return l.length
}

func (l *List) Push(element int) {
	defer func() { l.length++ }()
	if l.first == nil {
		l.first = &Element{value: element, next: nil}
		return
	}

	next := l.first
	for next.next != nil {
		next = next.next
	}

	next.next = &Element{value: element, next: nil}
}

func (l *List) Pop() (int, error) {
	if l.length == 0 {
		return 0, errors.New("list is empty")
	}

	previous := l.first
	last := l.first

	for {
		if last.next == nil {
			l.length--
			previous.next = nil
			return last.value, nil
		}
		previous = last
		last = last.next
	}
}

func (l *List) Peek() (int, error) {
	if l.length == 0 {
		return 0, errors.New("list is empty")
	}

	next := l.first
	for next.next != nil {
		next = next.next
	}

	v := next.value
	return v, nil
}

func (l *List) Array() []int {
	arr := make([]int, 0, l.length)
	if l.first == nil {
		return arr
	}

	next := l.first
	for i := 1; i <= l.length; i++ {
		arr = append(arr, next.value)
		next = next.next
	}

	return arr
}

func (l *List) Reverse() *List {
	arr := l.Array()
	reversed := make([]int, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}

	return New(reversed)
}
