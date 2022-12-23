import { readFileSync } from "../utils/readfile.ts";

const regexp =
  /^Valve (?<id>[A-Z]{2}) has flow rate=(?<flow_rate>[0-9]+); tunnels? leads? to valves? (?<tunnels>(?:[A-Z]{2},? ?)+)$/;

const flow_rates: Record<string, number> = {};
const graph: Record<string, string[]> = {};
const bit: Record<string, bigint> = {};
const vertices: string[] = [];

// const part1 = (path: string) => {
//   const input = parseInput(path);
//   pathCache.clear();
//   evalCache.clear();

//   const bestPressure = evaluate(input);

//   console.log("Part 1:", bestPressure);
// };

// const part2 = (path: string) => {
//   const input = parseInput(path);
//   pathCache.clear();
//   evalCache.clear();

//   const MINUTES = 26;
//   const ratedValves = [...input]
//     .filter(([, valve]) => valve.rate > 0)
//     .map(([name]) => name);
//   ratedValves.sort();
//   const ratedValveCount = ratedValves.length;
//   const maxValvesPerAgent = Math.min(
//     ratedValveCount,
//     Math.floor((MINUTES - 1) / 3),
//   );
//   const minBanned = Math.max(1, ratedValveCount - maxValvesPerAgent);
//   const maxBanned = Math.floor(ratedValveCount / 2);
//   const banLists = new Array(1 << ratedValveCount)
//     .fill(0)
//     .map((e1, i) => ratedValves.filter((e2, j) => i & (1 << j)))
//     .filter((list) => list.length >= minBanned && list.length <= maxBanned);
//   let bestPressure = 0;
//   const attemptedBanInverts: Set<string> = new Set();
//   for (const banList of banLists) {
//     const banString = banList.join(":");
//     const banInverse = ratedValves.filter((v) => !banList.includes(v));
//     if (attemptedBanInverts.has(banString)) continue;
//     const banInverseString = banInverse.join(":");
//     const pressure = evaluate(input, MINUTES, banList);
//     const inversePressure = evaluate(input, MINUTES, banInverse);
//     const combinedPressure = pressure + inversePressure;
//     if (combinedPressure > bestPressure) bestPressure = combinedPressure;
//     attemptedBanInverts.add(banInverseString);
//   }

//   console.log("Part 2:", bestPressure);
// };

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;

  let bit_index = 0n;
  for (const line of readFileSync(inputPath).split("\n")) {
    const { id, flow_rate, tunnels } = line.match(regexp)!.groups!;
    flow_rates[id] = parseInt(flow_rate);
    graph[id] = tunnels.split(", ");
    if (flow_rates[id] !== 0) {
      bit[id] = 1n << bit_index;
      bit_index++;
    }

    vertices.push(id);
  }

  const objectFromKeys = <T>(keys: string[], valueFn: (key: string) => T) =>
    Object.fromEntries(keys.map((k) => [k, valueFn(k)]));

  let distances: Record<string, Record<string, number>> = objectFromKeys(
    vertices,
    () => objectFromKeys(vertices, () => Infinity),
  );

  for (const v of vertices) {
    distances[v][v] = 0;
    for (const u of graph[v]) {
      distances[u][v] = 1;
    }
  }

  for (const k of vertices) {
    for (const i of vertices) {
      for (const j of vertices) {
        if (distances[i][j] > distances[i][k] + distances[k][j]) {
          distances[i][j] = distances[i][k] + distances[k][j];
        }
      }
    }
  }

  for (const [from, mapping] of Object.entries(distances)) {
    for (const to of Object.keys(mapping)) {
      if (to === from || flow_rates[to] === 0) {
        delete mapping[to];
      }
    }
  }

  const cache = new Map<string, number>();
  const dfs = (time: number, valve: string, open: bigint): number => {
    const cacheKey = `${time},${valve},${open}`;
    const cachedValue = cache.get(cacheKey);
    if (cachedValue !== undefined) {
      return cachedValue;
    }

    let max_val = 0;
    for (const [neighbour, dist] of Object.entries(distances[valve])) {
      if (open & bit[neighbour]) {
        continue;
      }
      const remaining = time - 1 -
        dist;
      if (remaining <= 0) {
        continue;
      }
      const newBest = dfs(remaining, neighbour, open | bit[neighbour]);
      max_val = Math.max(max_val, newBest + flow_rates[neighbour] * remaining);
    }

    cache.set(cacheKey, max_val);
    return max_val;
  };

  console.log("Part 1:", dfs(30, "AA", 0n));

  let max = 0;
  const max_bitmask = vertices.map((v) => bit[v]).reduce(
    (a, b) => a | (b ?? 0n),
    0n,
  );
  for (let i = 0n; i < max_bitmask / 2n; i++) {
    const v = dfs(26, "AA", i) + dfs(26, "AA", max_bitmask ^ i);
    if (v > max) {
      max = v;
    }
  }

  console.log("Part 2:", max);
};

export default run;
