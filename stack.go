package practice

type Stack []int

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
