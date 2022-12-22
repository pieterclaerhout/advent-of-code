import { readFileSync } from "../utils/readfile.ts";

const parseInput = (path: string): number[] => {
  return readFileSync(path)
    .split("\n\n")
    .map((line: string): number => {
      return line.split("\n")
        .map((v) => +v)
        .reduce((p, c) => p + c);
    })
    .sort((a, b) => b - a);
}

const run = () => {
  const inputPath = new URL('input.txt', import.meta.url).pathname;
  const caloriesPerElf = parseInput(inputPath);

  const highest = caloriesPerElf[0];
  console.log("Part 1:", highest);

  const top3 = caloriesPerElf.slice(0, 3).reduce((p, c) => p + c);
  console.log("Part 2:", top3);
}

export default run;
