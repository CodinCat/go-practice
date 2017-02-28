package practice

// Move rods from s1 to s3
func Hanoi(s1, s2, s3 *Stack, rodsToMove int) {
	if rodsToMove == 1 {
		move(s1, s3)
		return
	}
	Hanoi(s1, s3, s2, rodsToMove-1)
	move(s1, s3)
	Hanoi(s2, s1, s3, rodsToMove-1)
}

func move(s1, s2 *Stack) {
	(*s2).Push((*s1).Pop())
}
