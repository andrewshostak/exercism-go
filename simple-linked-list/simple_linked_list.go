package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(numbers []int) *List {
	return &List{
		head: createLinkedList(numbers),
		size: len(numbers),
	}
}

func createLinkedList(numbers []int) *Element {
	if len(numbers) == 0 {
		return nil
	}

	element := &Element{data: numbers[0]}

	tmp := element
	for _, v := range numbers[1:] {
		t := &Element{data: v}
		tmp.next = t
		tmp = tmp.next
	}

	return element
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(number int) {
	if l.Size() == 0 {
		l.head = &Element{data: number}
		l.size++
		return
	}

	last := l.head
	for {
		if last.next == nil {
			last.next = &Element{data: number}
			break
		}
		last = last.next
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.Size() == 0 {
		return 0, errors.New("pop from empty list is not allowed")
	}

	last := l.head
	previous := l.head
	for {
		if last.next == nil {
			l.size--
			previous.next = nil
			return last.data, nil
		}
		previous = last
		last = last.next
	}
}

func (l *List) Array() []int {
	var array []int

	element := l.head
	for ; element != nil; element = element.next {
		array = append(array, element.data)
	}

	return array
}

func (l *List) Reverse() *List {
	if l.Size() == 0 {
		return l
	}

	array := l.Array()
	reversed := make([]int, 0, len(array))

	for i := len(array) - 1; i >= 0; i-- {
		reversed = append(reversed, array[i])
	}

	l.head = createLinkedList(reversed)

	return l
}
