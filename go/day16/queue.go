package day16

type Queue []string

func (queue *Queue) PopLeft() string {
	elem := (*queue)[0]
	*queue = (*queue)[1:]
	return elem
}

func (queue *Queue) Contains(substr string) bool {
	for _, x := range *queue {
		if x == substr {
			return true
		}
	}
	return false
}
