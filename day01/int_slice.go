package day01

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

type IntSlice []int

func NewIntSlice(input string) IntSlice {
	result := IntSlice{}

	var currentSum int

	reader := strings.NewReader(input)
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		if sc.Text() == "" {
			result = append(result, currentSum)
			currentSum = 0
		}

		currentItem, err := strconv.Atoi(sc.Text())
		if err != nil {
			continue
		}

		currentSum += currentItem
	}

	return result
}

func (list IntSlice) Max() int {
	max := list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}

func (list IntSlice) SumTop(n int) int {
	var sum int
	for _, i := range list.Top(n) {
		sum += i
	}
	return sum
}

func (list IntSlice) Top(n int) []int {
	sortedList := list[:]
	sort.Sort(sort.Reverse(sort.IntSlice(sortedList)))

	if len(sortedList) < n {
		return sortedList
	}

	return sortedList[:n]
}
