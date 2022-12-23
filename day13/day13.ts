import { readFileSync } from "../utils/readfile.ts";

type NumberOrArray = number | NumberOrArray[];

const parseInput = (path: string) =>
  readFileSync(path)
    .split("\n\n")
    .map((pair) => {
      const [packet1, packet2] = pair.split("\n").map((p) => JSON.parse(p));
      return [packet1, packet2] as [NumberOrArray, NumberOrArray];
    });

const compareValues = (a: NumberOrArray, b: NumberOrArray): -1 | 0 | 1 => {
  if (!Array.isArray(a) && !Array.isArray(b)) {
    if (a < b) {
      return -1;
    }
    if (b < a) {
      return 1;
    }
    return 0;
  }
  if (!Array.isArray(b)) {
    b = [b];
  }
  if (!Array.isArray(a)) {
    a = [a];
  }
  for (let i = 0; i < Math.max(a.length, b.length); i++) {
    if (i + 1 > a.length) {
      return -1;
    } else if (i + 1 > b.length) {
      return 1;
    }
    const result = compareValues(a[i], b[i]);
    if (result !== 0) {
      return result;
    }
  }
  return 0;
};

const part1 = (path: string) => {
  const input = parseInput(path);

  let indexSum = 0;
  for (let p = 0; p < input.length; p++) {
    if (compareValues(...input[p]) === -1) {
      indexSum += p + 1;
    }
  }

  console.log("Part 1:", indexSum);
};

const part2 = (path: string) => {
  const input = parseInput(path);

  const flatPackets = [];
  for (const pair of input) {
    flatPackets.push(...pair);
  }

  const divider1 = [[2]];
  const divider2 = [[6]];
  flatPackets.push(divider1, divider2);
  flatPackets.sort(compareValues);

  const result = (
    (flatPackets.indexOf(divider1) + 1) *
    (flatPackets.indexOf(divider2) + 1)
  ).toString();

  console.log("Part 2:", result);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;

  part1(inputPath);
  part2(inputPath);
};

export default run;
