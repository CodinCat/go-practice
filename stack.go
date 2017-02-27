package practice

type Stack []int

// Create a new stack with initial items
func NewStack(items ...int) Stack {
	var stack Stack
	for _, item := range items {
		stack.Push(item)
	}
	return stack
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	itemsToPop := (*s)[len(*s)-1]
	s.removeLastItem()
	return itemsToPop
}

func (s *Stack) removeLastItem() {
	*s = append((*s)[:len(*s)-1])
}
