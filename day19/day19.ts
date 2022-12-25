const regex =
  /Each ore robot costs (?<ore_ore_cost>[0-9]+) ore\. Each clay robot costs (?<clay_ore_cost>[0-9]+) ore\. Each obsidian robot costs (?<obsidian_ore_cost>[0-9]+) ore and (?<obsidian_clay_cost>[0-9]+) clay\. Each geode robot costs (?<geode_ore_cost>[0-9]+) ore and (?<geode_obsidian_cost>[0-9]+) obsidian\.$/m;

type Resource = "ore" | "clay" | "obsidian" | "geode";
type Blueprint = Record<Resource, Map<Resource, number>>;

type ParsedInput = {
  blueprint: Blueprint;
  max_costs: Map<Resource, number>;
};

const parseInput = (rawInput: string): ParsedInput[] =>
  rawInput
    .replaceAll("\r", "").split("\n")
    .map((line): Blueprint => {
      const {
        ore_ore_cost,
        clay_ore_cost,
        obsidian_ore_cost,
        obsidian_clay_cost,
        geode_ore_cost,
        geode_obsidian_cost,
      } = line.match(regex)!.groups!;

      return {
        ore: new Map(
          Object.entries({ ore: parseInt(ore_ore_cost) }) as [
            Resource,
            number,
          ][],
        ),
        clay: new Map(
          Object.entries({ ore: parseInt(clay_ore_cost) }) as [
            Resource,
            number,
          ][],
        ),
        obsidian: new Map(
          Object.entries({
            ore: parseInt(obsidian_ore_cost),
            clay: parseInt(obsidian_clay_cost),
          }) as [Resource, number][],
        ),
        geode: new Map(
          Object.entries({
            ore: parseInt(geode_ore_cost),
            obsidian: parseInt(geode_obsidian_cost),
          }) as [Resource, number][],
        ),
      };
    }).map((blueprint) => {
      const max_costs = new Map<Resource, number>();
      for (const costs of Object.values(blueprint)) {
        for (const [resource_type, amount] of costs) {
          if (
            !max_costs.has(resource_type) ||
            max_costs.get(resource_type)! < amount
          ) {
            max_costs.set(resource_type, amount);
          }
        }
      }

      return {
        blueprint: blueprint,
        max_costs: max_costs,
      };
    });

const cache = new Map<string, number>();

const dfs = (
  input: ParsedInput[],
  time_remaining: number,
  blueprintID: number,
  robots: Record<Resource, number> = { ore: 1, clay: 0, obsidian: 0, geode: 0 },
  resources: Record<Resource, number> = {
    ore: 0,
    clay: 0,
    obsidian: 0,
    geode: 0,
  },
): number => {
  if (time_remaining < 0) {
    throw "?";
  }
  if (time_remaining === 0) {
    return resources.geode;
  }
  const cache_key = JSON.stringify([
    time_remaining,
    blueprintID,
    robots,
    resources,
  ]);
  const cached_result = cache.get(cache_key);
  if (typeof cached_result === "number") {
    return cached_result;
  }

  const { blueprint, max_costs } = input[blueprintID];
  let max_geodes = resources.geode + robots.geode * time_remaining;
  robots:
  for (const robot_type of ["clay", "ore", "obsidian", "geode"] as Resource[]) {
    if (robots[robot_type] >= max_costs.get(robot_type)!) {
      continue robots;
    }

    const new_resources: Record<Resource, number> = { ...resources };

    let time_to_make_robot = 0;
    for (const [resource, total_needed] of blueprint[robot_type]) {
      new_resources[resource] -= total_needed;
      const more_needed = total_needed - resources[resource];
      const time_to_get_more = Math.ceil(more_needed / robots[resource]);
      if (time_to_get_more >= time_remaining) {
        continue robots;
      }

      time_to_make_robot = Math.max(time_to_make_robot, time_to_get_more);
    }

    time_to_make_robot += 1;

    const new_time_remaining = time_remaining - time_to_make_robot;

    for (const resource of Object.keys(new_resources) as Resource[]) {
      new_resources[resource] += robots[resource] * time_to_make_robot;

      if (resource !== "geode") {
        const most_spend = new_time_remaining * max_costs.get(resource)!;
        if (new_resources[resource] > most_spend) {
          new_resources[resource] = most_spend;
        }
      }
    }

    const new_robots = { ...robots };
    new_robots[robot_type] += 1;

    const result = dfs(
      input,
      time_remaining - time_to_make_robot,
      blueprintID,
      new_robots,
      new_resources,
    );

    if (result > max_geodes) {
      max_geodes = result;
    }
  }

  cache.set(cache_key, max_geodes);
  return max_geodes;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const result = input.reduce(
    (acc, _, i) => acc + dfs(input, 24, i) * (i + 1),
    0,
  );

  console.log("Part 1:", result);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const result = input.slice(0, 3).reduce(
    (acc, _, i) => acc * dfs(input, 32, i),
    1,
  );

  console.log("Part 2:", result);
};
export default function (_inputPath: string, rawInput: string) {
  part1(rawInput);
  part2(rawInput);
}
