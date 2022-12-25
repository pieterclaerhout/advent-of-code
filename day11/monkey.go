package day11

type MonkeyOld struct {
	Items        []int
	Operation    func(int) int
	TestAndThrow func(int) int
	TestingValue int
}
