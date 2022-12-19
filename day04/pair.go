package day04

import "fmt"

type Pair struct {
	StartFirst  int
	EndFirst    int
	StartSecond int
	EndSecond   int
}

func (p Pair) OverlapsCompletely() bool {
	if p.StartSecond >= p.StartFirst && p.EndSecond <= p.EndFirst {
		return true
	}
	if p.StartFirst >= p.StartSecond && p.EndFirst <= p.EndSecond {
		return true
	}
	return false
}

func (p Pair) Overlaps() bool {
	if p.StartSecond <= p.EndFirst && p.EndSecond >= p.StartFirst {
		return true
	}
	if p.StartFirst <= p.EndSecond && p.EndFirst >= p.StartSecond {
		return true
	}
	return false
}

func NewPair(input string) Pair {
	p := Pair{}
	fmt.Sscanf(
		input,
		"%d-%d,%d-%d",
		&p.StartFirst,
		&p.EndFirst,
		&p.StartSecond,
		&p.EndSecond,
	)
	return p
}
