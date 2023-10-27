package linkedlist

import (
	"reflect"
	"testing"
)

func TestSingly_InsertHead(t *testing.T) {
	list := NewSingly[int]()
	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)

	want := []any{1, 2, 3}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_InsertTail(t *testing.T) {
	list := NewSingly[int]()
	list.InsertTail(3)
	list.InsertHead(2)
	list.InsertHead(1)
	list.InsertTail(4)
	list.InsertTail(5)
	list.InsertTail(6)

	want := []any{1, 2, 3, 4, 5, 6}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_Insert(t *testing.T) {
	list := NewSingly[int]()

	if err := list.Insert(-2, 4); err == nil {
		t.Errorf(err.Error())
	}

	if err := list.Insert(0, 1); err != nil {
		t.Errorf(err.Error())
	}

	if err := list.Insert(1, 2); err != nil {
		t.Errorf(err.Error())
	}

	if err := list.Insert(2, 3); err != nil {
		t.Errorf(err.Error())
	}

	if err := list.Insert(4, 6); err == nil {
		t.Errorf(err.Error())
	}

	if err := list.Insert(1, 4); err != nil {
		t.Errorf(err.Error())
	}

	want := []any{1, 4, 2, 3}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_DeleteHead(t *testing.T) {
	list := NewSingly[int]()

	if err := list.DeleteHead(); err == nil {
		t.Errorf(err.Error())
	}

	list.InsertTail(3)
	list.InsertHead(2)
	list.InsertHead(1)

	if err := list.DeleteHead(); err != nil {
		t.Errorf(err.Error())
	}

	want := []any{2, 3}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_DeleteTail(t *testing.T) {
	list := NewSingly[int]()

	if err := list.DeleteTail(); err == nil {
		t.Errorf(err.Error())
	}

	list.InsertTail(3)
	list.InsertHead(2)
	list.InsertHead(1)

	if err := list.DeleteTail(); err != nil {
		t.Errorf(err.Error())
	}

	want := []any{1, 2}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_Delete(t *testing.T) {
	list := NewSingly[int]()

	if err := list.Delete(-2); err == nil {
		t.Errorf(err.Error())
	}

	list.InsertTail(3)
	list.InsertTail(5)
	list.InsertTail(7)
	list.InsertHead(2)
	list.InsertHead(1)

	if err := list.Delete(1); err != nil {
		t.Errorf(err.Error())
	}

	if err := list.Delete(0); err != nil {
		t.Errorf(err.Error())
	}

	if err := list.Delete(7); err == nil {
		t.Errorf(err.Error())
	}

	want := []any{3, 5, 7}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_Head(t *testing.T) {
	list := NewSingly[int]()

	_, err := list.Head()

	if err == nil {
		t.Errorf(err.Error())
	}

	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)

	want := 1
	got, err := list.Head()

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_Tail(t *testing.T) {
	list := NewSingly[int]()

	_, err := list.Tail()

	if err == nil {
		t.Errorf(err.Error())
	}

	list.InsertTail(1)
	list.InsertTail(2)
	list.InsertTail(3)

	want := 3
	got, err := list.Tail()

	if err != nil {
		t.Errorf(err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_Len(t *testing.T) {
	list := NewSingly[int]()
	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)

	want := 3
	got := list.Len()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSingly_IsEmpty(t *testing.T) {
	list := NewSingly[int]()

	if !list.IsEmpty() {
		t.Errorf(ErrNotEmpty.Error())
	}

	list.InsertHead(3)
	list.InsertHead(2)
	list.InsertHead(1)

	if list.IsEmpty() {
		t.Errorf(ErrNotEmpty.Error())
	}
}

func TestSingly_Tranverse(t *testing.T) {
	list := NewSingly[int]()
	list.InsertHead(1)
	list.InsertHead(2)
	list.InsertHead(3)

	want := []any{3, 2, 1}
	got := []any{}
	cur := list.head
	got = append(got, cur.data)

	for cur.next != nil {
		cur = cur.next
		got = append(got, cur.data)
	}

	list.Transverse()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
