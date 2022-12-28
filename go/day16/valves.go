package day16

type Valves struct {
	FlowRates   map[string]int
	Connections map[string][]string
}

func (valves *Valves) NonEmpty() map[string]int {
	nonEmpty := map[string]int{}
	for label, flow := range valves.FlowRates {
		if flow > 0 {
			nonEmpty[label] = flow
		}
	}
	return nonEmpty
}

func (valves *Valves) ValvesDistances(nonEmpty map[string]int) map[string]map[string]int {
	distances := map[string]map[string]int{}

	for origin := range nonEmpty {
		distances[origin] = valves.computeDistanceToAll(nonEmpty, origin)
	}

	distances["AA"] = valves.computeDistanceToAll(nonEmpty, "AA")

	return distances
}

func (valves *Valves) computeDistanceToAll(nonEmpty map[string]int, origin string) map[string]int {
	distances := map[string]int{}

	dist := 0
	q := append(Queue{}, valves.Connections[origin]...)

	for dist <= 26 {
		dist++
		tmpQueue := Queue{}
		for len(q) > 0 {
			next := q.PopLeft()
			if _, ok := distances[next]; ok || next == origin {
				continue
			}
			if _, ok := nonEmpty[next]; ok {
				distances[next] = dist
			}
			for _, conn := range valves.Connections[next] {
				if !tmpQueue.Contains(conn) {
					tmpQueue = append(tmpQueue, conn)
				}
			}
		}
		q = append(q, tmpQueue...)
	}

	return distances
}
