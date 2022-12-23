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
};

const part1 = (parsedInput: number[]) => {
  const highest = parsedInput[0];
  console.log("Part 1:", highest);
};

const part2 = (parsedInput: number[]) => {
  const top3 = parsedInput.slice(0, 3).reduce((p, c) => p + c);
  console.log("Part 2:", top3);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;
  const caloriesPerElf = parseInput(inputPath);

  part1(caloriesPerElf);
  part2(caloriesPerElf);
};

export default run;
