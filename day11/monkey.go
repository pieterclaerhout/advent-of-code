package day11

type Monkey struct {
	Items        []int
	Operation    func(int) int
	TestAndThrow func(int) int
	TestingValue int
}
