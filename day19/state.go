package day19

const (
	Ore      int = 0
	Clay     int = 1
	Obsidian int = 2
	Geode    int = 3
)

type State struct {
	MaxMinutes int
	Blueprint  Blueprint
	Cache      map[string]int
	MaxRate    [4]int
}

func max(a ...int) int {
	max := a[0]
	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

func NewState(b Blueprint) State {
	return State{
		MaxMinutes: 24,
		Blueprint:  b,
		Cache:      make(map[string]int),
		MaxRate: [4]int{
			max(b[Ore][Ore], b[Clay][Ore], b[Obsidian][Ore], b[Geode][Ore]),
			max(b[Ore][Clay], b[Clay][Clay], b[Obsidian][Clay], b[Geode][Clay]),
			max(b[Ore][Obsidian], b[Clay][Obsidian], b[Obsidian][Obsidian], b[Geode][Obsidian]),
			max(b[Ore][Geode], b[Clay][Geode], b[Obsidian][Geode], b[Geode][Geode]),
		},
	}
}

func (st *State) Dfs(minute int, resources, robots, doNotBuild Resources) int {
	if minute >= st.MaxMinutes {
		return resources.Geodes
	}

	nextDoNotBuild := Resources{}
	if resources.Ores >= st.Blueprint[Geode][Ore] &&
		resources.Obsidian >= st.Blueprint[Geode][Obsidian] {
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.Geodes++
		newResources.Ores -= st.Blueprint[Geode][Ore]
		newResources.Obsidian -= st.Blueprint[Geode][Obsidian]
		return st.Dfs(minute+1, newResources, newRobots, Resources{})
	}

	maxGeodes := 0

	if doNotBuild.Ores == 0 &&
		resources.Ores >= st.Blueprint[Ore][Ore] && robots.Ores < st.MaxRate[Ore] {
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.Ores++
		newResources.Ores -= st.Blueprint[Ore][Ore]
		newMaxGeodes := st.Dfs(minute+1, newResources, newRobots, Resources{})
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
		nextDoNotBuild.Ores++
	}

	if doNotBuild.Clay == 0 &&
		resources.Ores >= st.Blueprint[Clay][Ore] && robots.Clay < st.MaxRate[Clay] {
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.Clay++
		newResources.Ores -= st.Blueprint[Clay][Ore]
		newMaxGeodes := st.Dfs(minute+1, newResources, newRobots, Resources{})
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
		nextDoNotBuild.Clay++
	}

	if doNotBuild.Obsidian == 0 &&
		resources.Ores >= st.Blueprint[Obsidian][Ore] &&
		resources.Clay >= st.Blueprint[Obsidian][Clay] &&
		robots.Obsidian < st.MaxRate[Obsidian] {
		newResources := mineResources(resources, robots)
		newRobots := robots.Clone()
		newRobots.Obsidian++
		newResources.Ores -= st.Blueprint[Obsidian][Ore]
		newResources.Clay -= st.Blueprint[Obsidian][Clay]
		newMaxGeodes := st.Dfs(minute+1, newResources, newRobots, Resources{})
		if newMaxGeodes > maxGeodes {
			maxGeodes = newMaxGeodes
		}
		nextDoNotBuild.Obsidian++
	}

	resources = mineResources(resources, robots)
	newMaxGeodes := st.Dfs(minute+1, resources, robots, nextDoNotBuild)
	if newMaxGeodes > maxGeodes {
		maxGeodes = newMaxGeodes
	}

	return maxGeodes
}
