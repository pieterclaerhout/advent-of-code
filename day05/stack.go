package day05

type stack struct {
	elements []rune
}

func (s *stack) push(r ...rune) {
	s.elements = append(s.elements, r...)
}

func (s *stack) pop() rune {
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return r
}

func (s *stack) popN(n int) []rune {
	r := s.elements[len(s.elements)-n : len(s.elements)]
	s.elements = s.elements[:len(s.elements)-n]
	return r
}

func (s *stack) addToBottom(r rune) {
	s.elements = append([]rune{r}, s.elements...)
}

func (s stack) String() string {
	var str string
	for _, r := range s.elements {
		str += string(r) + " "
	}
	return str
}
