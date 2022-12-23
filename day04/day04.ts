import { readFileSync } from "../utils/readfile.ts";

type Range = [number, number];
type Pair = [Range, Range];

const parseInput = (path: string): Pair[] => {
  return readFileSync(path)
    .split("\n")
    .map((v) => v.split(",").map((r) => r.split("-").map((s) => +s)) as Pair);
};

const isContained = (pair: Pair): boolean => {
  const [left, right] = pair;
  if (left[0] <= right[0] && right[1] <= left[1]) {
    return true;
  }
  if (right[0] <= left[0] && left[1] <= right[1]) {
    return true;
  }
  return false;
};

const isOverlapping = (pair: Pair): boolean => {
  const [left, right] = pair;
  let [min, max] = [right, left];
  if (left[0] < right[0]) {
    [min, max] = [left, right];
  }
  return min[1] >= max[0];
};

const part1 = (parsedInput: Pair[]) => {
  const result = parsedInput.filter(isContained).length;
  console.log("Part 1:", result);
};

const part2 = (parsedInput: Pair[]) => {
  const result = parsedInput.filter(isOverlapping).length;
  console.log("Part 2:", result);
};

export default function (inputPath: string) {
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
}
