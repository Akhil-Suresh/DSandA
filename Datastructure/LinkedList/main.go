package main

import linkedlist "LinkedList"
import "fmt"


func main() {
	ll := linkedlist.NewLinkedList(10, 20)
	ll.append(30)
	ll.prepend(5)
	ll.insert(2, 15)

	fmt.Println(ll.print())
	fmt.Println("Linked List Length: ", ll.Length)
	ll.delete(2)
	fmt.Println(ll.print())
	fmt.Println(ll.search(10))
}