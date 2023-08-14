package sll

import (
	"fmt"
	"strings"
	"testing"
)

type node[T comparable] struct {
	data T
	next *node[T]
}

func NewNode[T comparable](data T) *node[T] {
	return &node[T]{data: data}
}

func (n *node[T]) String() string {
	return fmt.Sprintf("%v", n.data)
}

func TestNode(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(11)
	n1.next = n2
	fmt.Printf("n1: %v\n", n1)
	fmt.Printf("n1.String(): %v\n", n1.String())
}

/*
Homogenous singly linked list
*/
type linkedList[T comparable] struct {
	head *node[T]
}

func NewLinkedList[T comparable]() *linkedList[T] {
	return &linkedList[T]{}
}

func (list linkedList[T]) IsEmpty() bool {
	return list.head == nil
}

/*
Return the number  of node in the list, it takes O(n) time.
*/
func (list linkedList[T]) Size() int {
	current := list.head
	count := 0

	for current != nil {
		current = current.next
		count++
	}

	return count
}

/*
Add new node containing data at head of the list, it takes O(1) time.
*/
func (list *linkedList[T]) Add(data T) {
	newNode := NewNode(data)
	newNode.next = list.head
	list.head = newNode
}

/*
Produce string values of the list, it takes O(n) time.
*/
func (list linkedList[T]) String() string {
	nodes := []string{}
	current := list.head

	for current != nil {
		if current == list.head {
			nodes = append(nodes, fmt.Sprintf("[Head: %v]", current.data))
		} else if current.next == nil {
			nodes = append(nodes, fmt.Sprintf("[Tail: %v]", current.data))
		} else {
			nodes = append(nodes, fmt.Sprintf("[%v]", current.data))
		}

		current = current.next
	}

	return strings.Join(nodes, "-> ")
}

/*
Check whether given key value exist on the list, it takes O(n) time.
*/
func (list linkedList[T]) Exists(key T) bool {
	current := list.head

	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}
	return false
}

/*
Check the first node containing data that matches the given key value, it takes O(n) time.
It will return the node if founded, or nil if not found.
*/
func (list linkedList[T]) Search(key T) *node[T] {
	current := list.head

	for current != nil {
		if current.data == key {
			return current
		}
		current = current.next
	}
	return nil
}

/*
Insert a new note containing data in idex possition.
Inserting takes O(1) time, but finding the intesrtion point takes O(n)

It takes overall O(n) time.
*/
func (list *linkedList[T]) Insert(data T, index int) {
	if index == 0 {
		list.Add(data)
		return
	}

	newNode := NewNode(data)
	currentNode := list.head
	position := index
	for position > 1 {
		currentNode = currentNode.next
		position--
	}

	previousNode := currentNode
	nextNode := currentNode.next
	previousNode.next = newNode
	newNode.next = nextNode
}

/*
Remove element from the list and return the value, it takes O(n) time.
*/
func (list *linkedList[T]) Remove(key T) *node[T] {
	current := list.head
	var previous *node[T]
	isFound := false

	for current != nil && !isFound {
		if current.data == key && current == list.head {
			list.head = current.next
			isFound = true
		} else if current.data == key {
			previous.next = current.next
			isFound = true
		} else {
			previous = current
			current = current.next
		}
		previous = current
		current = current.next
	}

	return current
}

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(10)
	list.Add(11)
	list.Add(12)
	list.Add(13)
	list.Add(14)

	fmt.Printf("list.IsEmpty(): %v\n", list.IsEmpty())
	fmt.Printf("list.Size(): %v\n", list.Size())
	fmt.Printf("list: %v\n", list)
	fmt.Printf("list.String(): %v\n", list.String())
	fmt.Printf("list.Exists(10): %v\n", list.Exists(10))
	fmt.Printf("list.Search(1): %v\n", list.Search(1))

	list.Insert(5, 1)
	fmt.Printf("list inserted 5 on 0th index: %v\n", list)

	list.Insert(6, 5)
	fmt.Printf("list inserted 6 on 5th index: %v\n", list)

	list.Remove(10)
	fmt.Printf("list removed 10: %v\n", list)
}
