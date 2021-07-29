package main

import "fmt"

/*
stack looks like stack of plates. so last in plate will be taken first.

so each time we add to our stack we move our top pointer one level down.

Let's say our stack looks like this below
push1	->	[	"google"	]		<- top		next -> nil
push2 	->	[	"facebook"	]		<- New top	next -> ["google"]
push3	->	[	"linkedin"	]		<- New top	next -> ["facebook"]
push3 	->	[	"twitter"	]		<- bottom	next -> ["linkedin"]

Pop1    ->	"twitter"
Pop2	->  "facebook"
Pop3	->	"google"
*/
type node struct {
	value interface{}
	prev *node
}

func NewNode(val interface{}) *node {
	return &node{val, nil}
}

type Stack struct {
	top *node
	length	int
}

func (s *Stack) Push(val interface{}) {
	newNode := NewNode(val)
	if s.top == nil {
		s.top = newNode
	} else {
		topPointer := s.top
		newTop := newNode
		newTop.prev = topPointer
		s.top = newTop
	}
	s.length++
}

func (s *Stack) Pop() (interface{}, int) {
	if s.top == nil {
		return nil, s.length
	}
	top := s.top
	newTop := top.prev
	s.top = newTop
	top.prev = nil		// for garbage collection.
	length := s.length
	s.length--
	return top.value, length
}


func NewStack() *Stack{
	return &Stack{}
}


/*
Queue behaves in first in first out way.
ox01	["first",	ox02]	->	top 
ox02	["second",	ox03]	-> 	
ox03	["third",	nil]	->
*/

type Queue []interface{}

func (q *Queue) push(val interface{}){
	*q = append(*q, val)
}

func (q *Queue) pop() interface{}{
	if len(*q) < 1 {
		return nil
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func main() {
	stack := NewStack()
	stack.Push("Discord")
	stack.Push("google")
	stack.Push("Youtube")
	fmt.Println(stack.Pop())
	stack.Push("twitter")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	queue := new(Queue)
	queue.push("google")
	blah_json := make(map[string]string)
	blah_json["id"] = "1"
	blah_json["name"] = "akhil"
	queue.push(blah_json)
	queue.push("25")
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
}