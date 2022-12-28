package day05

type Stack struct {
	elements []rune
}

func (s *Stack) Push(r ...rune) {
	s.elements = append(s.elements, r...)
}

func (s *Stack) Pop() rune {
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return r
}

func (s *Stack) PopN(n int) []rune {
	r := s.elements[len(s.elements)-n : len(s.elements)]
	s.elements = s.elements[:len(s.elements)-n]
	return r
}

func (s *Stack) AddToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s Stack) String() string {
	var str string
	for _, r := range s.elements {
		str += string(r) + " "
	}
	return str
}
