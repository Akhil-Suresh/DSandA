package main

type Queue struct {
}

func(s *Stack) push(val interface{}) {
	newNode := NewNode(val)
	if s.top == nil {
		s.top = newNode
	} else {
		currentBottom := traverse(s)
		s.bottom = newNode
		currentBottom.next = newNode
		
	}
	s.length += 1
}

func(s *Stack) peek() {

}
func(s *Stack) pop() interface{} {
	bottomNode := s.bottom
	s.bottom = bottomNode.prev
	topNode.next = nil

	return topNode.value
}