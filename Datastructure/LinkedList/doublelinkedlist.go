package main

import (
	"errors"
	"fmt"
)

type node struct {
	Data interface{}
	Prev *node
	Next *node	
}

type DLinkedList struct {
	Length int
	Head	*node
	Tail	*node
}

func NewNode(val interface{}) *node {
	return &node{val, nil, nil}
}

func NewDLinkedList(headValue, tailValue interface{})*DLinkedList {
	head := NewNode(headValue)
	tail := NewNode(tailValue)
	head.Next = tail
	tail.Prev = head
	return &DLinkedList{2, head, tail}
}
// BigO O(1)
func(l *DLinkedList) append(val interface{}){
	newTail := NewNode(val)
	newTail.Prev = l.Tail
	
	currentTail := l.Tail
	currentTail.Next = newTail
	
	l.Tail = newTail
	
	l.Length += 1
}
// BigO O(1)
func(l *DLinkedList) prepend(val interface{}){
	newHead := NewNode(val)
	currentHead := l.Head
	newHead.Next = currentHead
	currentHead.Prev = newHead
	l.Head = newHead
	l.Length += 1
}

// BigO O(n)
// Index starts at Head which is zero.
func(l *DLinkedList) insert(index int, val interface{}) {
	if index > l.Length {
		panic(InvalidIndex)
	} 

	if index == 0 {
		l.prepend(val)
		return
	}

	newNode := NewNode(val)
	//	0	1	2	3	 
	//	11	12	13	14
	// we need to insert at 2 20
	// 
	indexCount := 0
	currentNode := l.Head	
	// we need to find index node
	for indexCount < index {
		currentNode = currentNode.Next
		indexCount += 1
	}
	prevNode := currentNode.Prev
	prevNode.Next = newNode
	newNode.Next = currentNode
	newNode.Prev = prevNode
	currentNode.Prev = newNode
	l.Length  += 1
	
}

// for this to work I need prervious node and currentnode at index position
func(l *DLinkedList) delete(index int) {
	indexNode := l.Head
	indexCount := 0
	// index - 1 will get previous node
	for indexCount < index {
		indexCount += 1
		indexNode = indexNode.Next
	}
	prevIndexNode := indexNode.Prev
	afterIndexNode := indexNode.Next
	prevIndexNode.Next = afterIndexNode
	afterIndexNode.Prev = prevIndexNode
}

// BigO O(n)
func(l *DLinkedList) search(val interface{}) int {
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
func(l *DLinkedList) print() string {
	var output string
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode)
		output += fmt.Sprintf("%v, ", currentNode.Data)
		currentNode = currentNode.Next
	}
	return output
}