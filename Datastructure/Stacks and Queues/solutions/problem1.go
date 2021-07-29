package main

import "fmt"


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
func (q *Queue) bottom() interface{} {
	return (*q)[len(*q) -1]
}

func (q *Queue) length() int {
	return len((*q))
}

type MyStack struct {
    length int
    queue  *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{0, new(Queue)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x interface{})  {
    this.queue.push(x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() interface{} {
    queue_length := this.queue.length()
    for count:= 1; count < queue_length; count++ {
        this.queue.push(this.queue.pop())
    }
    return this.queue.pop()
}

func (this *MyStack) Top() interface{} {
	return this.queue.bottom()
}

func (this *MyStack) Empty() bool {
	return this.queue.length() == 0
}

func main() {
	myStack := Constructor()
	myStack.Push("apple")
	myStack.Push("grape")
	myStack.Push("orange")
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
	fmt.Println(myStack.Top())
}
