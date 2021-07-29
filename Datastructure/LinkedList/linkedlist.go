package main

import (
	"errors"
	"fmt"
)

var InvalidIndex = errors.New("Invalid Index")

type node struct {
	Data interface{}
	Next *node
}

type LinkedList struct {
	Length int
	Head	*node
	Tail	*node
}

func NewNode(val interface{}) *node {
	return &node{val, nil}
}

func NewLinkedList(headValue, tailValue interface{})*LinkedList {
	head := NewNode(headValue)
	tail := NewNode(tailValue)
	head.Next = tail
	return &LinkedList{2, head, tail}
}
// BigO O(1)
func(l *LinkedList) append(val interface{}){
	node := NewNode(val)
	l.Tail.Next = node
	l.Length += 1
}
// BigO O(1)
func(l *LinkedList) prepend(val interface{}){
	newHead := NewNode(val)
	currentHead := l.Head
	newHead.Next = currentHead
	l.Head = newHead
	l.Length += 1
}

// BigO O(n)
// Index starts at Head which is zero.
func(l *LinkedList) insert(index int, val interface{}) {
	if index > l.Length {
		panic(InvalidIndex)
	} 

	if index == 0 {
		l.prepend(val)
		return
	}

	newNode := NewNode(val)

	indexCount := 1
	currentNode := l.Head	
	// we need to find prev index node
	// then we update its next to our new node
	// then we update our new node next to index node
	// so currentNode is going to be beforeIndexNode 
	for indexCount < index {
		currentNode = currentNode.Next
		indexCount += 1
	}
	afterIndexNode := currentNode.Next
	currentNode.Next = newNode
	newNode.Next = afterIndexNode
	l.Length  += 1
	
}

// for this to work I need prervious node and currentnode at index position
func(l *LinkedList) delete(index int) {
	previousNode := l.Head
	indexCount := 0
	// index - 1 will get previous node
	for indexCount < index - 1 {
		indexCount += 1
		previousNode = previousNode.Next
	}
	afterIndexNode := previousNode.Next.Next
	previousNode.Next = afterIndexNode
}

// BigO O(n)
func(l *LinkedList) search(val interface{}) int {
	currentHead := l.Head
	index := 0
	for currentHead.Next != nil {
		if val == currentHead.Data {
			return index
		}
		currentHead = currentHead.Next
		index += 1
	}
	return -1
}

// BigO O(n)
func(l *LinkedList) print() string {
	var output string
	currentNode := l.Head
	for currentNode != nil {
		output += fmt.Sprintf("%v, ", currentNode.Data)
		currentNode = currentNode.Next
	}
	return output
}

func main() {
	ll := NewLinkedList(10, 20)
	fmt.Println("New Linked List")
	fmt.Println(ll.print())	
	
	fmt.Println("After append")
	ll.append(30)
	fmt.Println(ll.print())
	
	fmt.Println("After Prepend")
	ll.prepend(5)
	fmt.Println(ll.print())
	
	fmt.Println("After insert")
	ll.insert(2, 15)
	fmt.Println(ll.print())
	
	fmt.Println("Linked List Length: ", ll.Length)
	
	fmt.Println("After Delete")
	ll.delete(2)
	fmt.Println(ll.print())
	
	fmt.Println("search result : ", ll.search(10))
}