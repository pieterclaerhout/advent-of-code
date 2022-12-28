type ParsedInput = {
  valveFlowMap: Map<string, number>;
  tunnelMap: Map<string, string[]>;
};

const parseInput = (rawInput: string): ParsedInput => {
  const lines = rawInput.split("\n");

  const valves = new Map<string, number>();
  const tunnels = new Map<string, string[]>();

  lines.forEach((line) => {
    const valve = line.split(" has ")[0].slice(-2);
    const flow = Number(line.match(/\d+/)![0]);
    const tunnel = line
      .replace(/valve /, "valves ")
      .split("valves ")[1]
      .split(", ");
    valves.set(valve, flow);
    tunnels.set(valve, tunnel);
  });

  return {
    valveFlowMap: valves,
    tunnelMap: tunnels,
  };
};

const solvePart1 = (rawInput: string): number => {
  const { valveFlowMap, tunnelMap } = parseInput(rawInput);

  const distanceMap: Record<string, Record<string, number>> = {};
  const nonEmptyValves: string[] = [];

  for (const [currentValve, currentFlow] of valveFlowMap) {
    if (currentValve !== "AA") {
      if (currentFlow === 0) {
        continue;
      } else {
        nonEmptyValves.push(currentValve);
      }
    }

    distanceMap[currentValve] = { [currentValve]: 0, AA: 0 };

    const visitedNodes = new Set<string>(currentValve);
    const bfsQueue: [string, number][] = [[currentValve, 0]];

    while (bfsQueue.length > 0) {
      const [currentNode, currentDistance] = bfsQueue.shift()!;
      const neighboringNodes = tunnelMap.get(currentNode)!;

      for (const currentNeighbor of neighboringNodes) {
        if (visitedNodes.has(currentNeighbor)) {
          continue;
        }

        visitedNodes.add(currentNeighbor);

        if (valveFlowMap.get(currentNeighbor)) {
          const valveDistances = distanceMap[currentValve];
          valveDistances[currentNeighbor] = currentDistance + 1;
        }

        bfsQueue.push([currentNeighbor, currentDistance + 1]);
      }
    }

    delete distanceMap[currentValve][currentValve];
    if (currentValve !== "AA") {
      delete distanceMap[currentValve].AA;
    }
  }

  const valveIndices: Record<string, number> = {};

  nonEmptyValves.forEach((element, index) => {
    valveIndices[element] = index;
  });

  const dfsCache: Record<string, number> = {};

  const dfs = (
    remainingTime: number,
    valve: string,
    visitedBitmask: number,
  ) => {
    const cacheKey = [remainingTime, valve, visitedBitmask].join(",");

    if (cacheKey in dfsCache) {
      return dfsCache[cacheKey];
    }

    let maxFlow = 0;
    const valveDist = distanceMap[valve];
    for (const neighbor in valveDist) {
      const neighborBit = 1 << valveIndices[neighbor];

      if (visitedBitmask & neighborBit) {
        continue;
      }

      const neighborRemainingTime = remainingTime - (valveDist[neighbor] + 1);

      if (neighborRemainingTime <= 0) {
        continue;
      }

      const neighborFlow = dfs(
        neighborRemainingTime,
        neighbor,
        visitedBitmask | neighborBit,
      );
      const neighborTotalFlow = neighborFlow +
        valveFlowMap.get(neighbor)! * neighborRemainingTime;
      maxFlow = Math.max(maxFlow, neighborTotalFlow);
    }

    dfsCache[cacheKey] = maxFlow;
    return maxFlow;
  };

  return dfs(30, "AA", 0);
};

const solvePart2 = (rawInput: string): number => {
  const { valveFlowMap, tunnelMap } = parseInput(rawInput);

  const distanceMap: Record<string, Record<string, number>> = {};
  const nonEmptyValves: string[] = [];

  for (const [currentValve, currentFlow] of valveFlowMap) {
    if (currentValve !== "AA") {
      if (currentFlow === 0) {
        continue;
      } else {
        nonEmptyValves.push(currentValve);
      }
    }

    distanceMap[currentValve] = { [currentValve]: 0, AA: 0 };

    const visitedNodes = new Set<string>(currentValve);
    const bfsQueue: [string, number][] = [[currentValve, 0]];

    while (bfsQueue.length > 0) {
      const [currentNode, currentDistance] = bfsQueue.shift()!;
      const neighboringNodes = tunnelMap.get(currentNode)!;

      for (const currentNeighbor of neighboringNodes) {
        if (visitedNodes.has(currentNeighbor)) {
          continue;
        }

        visitedNodes.add(currentNeighbor);

        if (valveFlowMap.get(currentNeighbor)) {
          const valveDistances = distanceMap[currentValve];
          valveDistances[currentNeighbor] = currentDistance + 1;
        }

        bfsQueue.push([currentNeighbor, currentDistance + 1]);
      }
    }

    delete distanceMap[currentValve][currentValve];
    if (currentValve !== "AA") {
      delete distanceMap[currentValve].AA;
    }
  }

  const valveIndices: Record<string, number> = {};

  nonEmptyValves.forEach((element, index) => {
    valveIndices[element] = index;
  });

  const dfsCache: Record<string, number> = {};

  const dfs = (
    remainingTime: number,
    valve: string,
    visitedBitmask: number,
  ) => {
    const cacheKey = [remainingTime, valve, visitedBitmask].join(",");

    if (cacheKey in dfsCache) {
      return dfsCache[cacheKey];
    }

    let maxFlow = 0;
    const valveDist = distanceMap[valve];
    for (const neighbor in valveDist) {
      const neighborBit = 1 << valveIndices[neighbor];

      if (visitedBitmask & neighborBit) {
        continue;
      }

      const neighborRemainingTime = remainingTime - (valveDist[neighbor] + 1);

      if (neighborRemainingTime <= 0) {
        continue;
      }

      const neighborFlow = dfs(
        neighborRemainingTime,
        neighbor,
        visitedBitmask | neighborBit,
      );
      const neighborTotalFlow = neighborFlow +
        valveFlowMap.get(neighbor)! * neighborRemainingTime;
      maxFlow = Math.max(maxFlow, neighborTotalFlow);
    }

    dfsCache[cacheKey] = maxFlow;
    return maxFlow;
  };

  const b = (1 << nonEmptyValves.length) - 1;

  let maxFlow = 0;

  for (let i = 0; i < Math.ceil((b + 1) / 2); i++) {
    maxFlow = Math.max(maxFlow, dfs(26, "AA", i) + dfs(26, "AA", b ^ i));
  }

  return maxFlow;
};

export default function (rawInput: string): [number, number] {
  return [solvePart1(rawInput), solvePart2(rawInput)];
}
