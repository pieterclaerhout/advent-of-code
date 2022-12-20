package day16

import (
	"sort"
	"strings"
)

type Cache struct {
	TotalMinutes     int
	StateMaxPressure map[string]int
	Distances        map[string]map[string]int
	ValvesFlow       map[string]int
}

func (c *Cache) DFS(path []string, valveState string) int {
	currValve := path[len(path)-1]

	remaining, totalPressure := c.calculateTotalPressure(path)
	if n, ok := c.StateMaxPressure[valveState]; ok && n > totalPressure {
		return n
	}
	c.StateMaxPressure[valveState] = totalPressure

	maxPressure := totalPressure
	for nextValve := range c.Distances[currValve] {
		if remaining < 0 { // cannot reach it in time
			break
		}
		if strings.Contains(valveState, nextValve) { // already open
			continue
		}

		newPath := append(path, nextValve)
		newState := c.pathToState(newPath)
		pressure := c.DFS(newPath, newState)
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}

	return maxPressure
}

func (c *Cache) MaxExclusivePair() int {
	max := 0

	for state1, pressure1 := range c.StateMaxPressure {
		for state2, pressure2 := range c.StateMaxPressure {
			if c.overlap(state1, state2) {
				continue
			}
			if pressure1+pressure2 > max {
				max = pressure1 + pressure2
			}
		}
	}

	return max
}

func (c *Cache) calculateTotalPressure(path []string) (voidMinutes int, totalFlow int) {
	totalPressure := 0
	dist := 0

	for i, v := range path[1:] {
		neededMinutes := c.Distances[path[i]][v] + 1
		dist += neededMinutes
		totalPressure += (c.TotalMinutes - dist) * c.ValvesFlow[v]
	}

	return c.TotalMinutes - dist, totalPressure
}

func (c *Cache) pathToState(s []string) string {
	new := append([]string{}, s[1:]...)
	sort.Strings(new)
	return strings.Join(new, ",")
}

func (c *Cache) overlap(s1 string, s2 string) bool {
	for i, j := 0, 0; i <= len(s1)-2 && j <= len(s2)-2; {
		if s1[i:i+2] == s2[j:j+2] {
			return true
		}
		if s1[i:i+2] < s2[j:j+2] {
			i += 3
		} else {
			j += 3
		}
	}
	return false
}
