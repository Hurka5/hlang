package runestack

const EOF = -1 

type Node struct {
	r    rune
	next *Node
}

type RuneStack struct {
	start *Node
}

func New() RuneStack {
	return RuneStack{}
}

func (s *RuneStack) Push(r rune) {
	node := &Node{r: r}
	if s.start == nil {
		s.start = node
	} else {
		node.next = s.start
		s.start = node
	}
}

func (s *RuneStack) Pop() rune {
	if s.start == nil {
		return EOF
	} else {
		n := s.start
		s.start = n.next
		return n.r
	}
}

func (s *RuneStack) Clear() {
	s.start = nil
}

