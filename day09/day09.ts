import { readFileSync } from "../utils/readfile.ts";

type StepsWithDirection = ["L" | "R" | "U" | "D", number];

const parseInput = (path: string): StepsWithDirection[] => {
  return readFileSync(path)
    .split("\n")
    .map((v) => {
      const [dir, steps] = v.split(" ");
      return [dir, +steps] as StepsWithDirection;
    });
};

const clamp = (value: number, min: number, max: number) =>
  Math.min(Math.max(value, min), max);

const simulateRope = (
  input: StepsWithDirection[],
  tailKnots: number,
): Set<string> => {
  const head: [number, number] = [0, 0];
  const tails: [number, number][] = [];

  for (let t = 0; t < tailKnots; t++) {
    tails.push([0, 0]);
  }

  const tailGrids = new Set(["0:0"]);

  for (const [dir, steps] of input) {
    for (let s = 0; s < steps; s++) {
      const index = dir === "L" || dir === "R" ? 0 : 1;
      const delta = dir === "L" || dir === "U" ? -1 : 1;
      head[index] += delta;

      for (let t = 0; t < tails.length; t++) {
        const follow = t === 0 ? head : tails[t - 1];
        const hDist = follow[0] - tails[t][0];
        const vDist = follow[1] - tails[t][1];

        if (Math.abs(hDist) > 1 || Math.abs(vDist) > 1) {
          tails[t][0] += clamp(hDist, -1, 1);
          tails[t][1] += clamp(vDist, -1, 1);
          if (t === tails.length - 1) {
            tailGrids.add(tails[t].join(":"));
          }
        }
      }
    }
  }

  return tailGrids;
};

const part1 = (parsedInput: StepsWithDirection[]) => {
  const tailGrids = simulateRope(parsedInput, 1);
  console.log("Part 1:", tailGrids.size);
};

const part2 = (parsedInput: StepsWithDirection[]) => {
  const tailGrids = simulateRope(parsedInput, 9);
  console.log("Part 2:", tailGrids.size);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
};

export default run;
