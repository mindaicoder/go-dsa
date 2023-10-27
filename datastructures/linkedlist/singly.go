package linkedlist

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
}

func NewSinglyNode[T any](data T) *node[T] {
	return &node[T]{data: data, next: nil}
}

type singly[T any] struct {
	head *node[T]
}

func NewSingly[T any]() *singly[T] {
	return &singly[T]{head: nil}
}

func (l *singly[T]) InsertHead(data T) {
	newNode := NewSinglyNode[T](data)
	newNode.next = l.head
	l.head = newNode
}

func (l *singly[T]) InsertTail(data T) {
	if l.IsEmpty() {
		l.InsertHead(data)
		return
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = NewSinglyNode[T](data)
}

func (l *singly[T]) Insert(index int, data T) error {
	if index < 0 {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		l.InsertHead(data)
		return nil
	}

	if index == l.Len() {
		l.InsertTail(data)
		return nil
	}

	count := 0
	cur := l.head
	prev := cur

	for cur != nil {
		if index == count {
			newNode := NewSinglyNode[T](data)
			prev.next = newNode
			newNode.next = cur
			return nil
		}

		count++
		prev = cur
		cur = cur.next
	}

	return ErrIndexOutOfRange
}

func (l *singly[T]) DeleteHead() error {
	if l.IsEmpty() {
		return ErrEmpty
	}

	l.head = l.head.next
	return nil
}

func (l *singly[T]) DeleteTail() error {
	if l.IsEmpty() {
		return ErrEmpty
	}

	cur := l.head
	for cur.next.next != nil {
		cur = cur.next
	}
	cur.next = nil

	return nil
}

func (l *singly[T]) Delete(index int) error {
	if index < 0 {
		return ErrIndexOutOfRange
	}

	if index == 0 {
		return l.DeleteHead()
	}

	count := 0
	cur := l.head
	prev := cur

	for cur != nil {
		if index == count {
			prev.next = cur.next
			return nil
		}

		count++
		prev = cur
		cur = cur.next
	}

	return ErrIndexOutOfRange
}

func (l *singly[T]) Head() (T, error) {
	var data T

	if l.IsEmpty() {
		return data, ErrEmpty
	}

	return l.head.data, nil
}

func (l *singly[T]) Tail() (T, error) {
	var data T

	if l.IsEmpty() {
		return data, ErrEmpty
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	return cur.data, nil
}

func (l *singly[T]) Len() int {
	count := 1
	cur := l.head

	for cur.next != nil {
		count++
		cur = cur.next
	}

	return count
}

func (l *singly[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *singly[T]) Transverse() {
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}
