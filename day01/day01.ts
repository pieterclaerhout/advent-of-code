import { readAndSplitFileSync } from "../utils/readfile.ts";

const parseInput = (path: string): number[] => {
  return readAndSplitFileSync(path, "\n\n")
    .map((chunk) => chunk.split("\n").map((v) => +v).reduce((p, c) => p + c))
    .sort((a, b) => b - a);
};

export default function (inputPath: string) {
  const caloriesPerElf = parseInput(inputPath);

  const highest = caloriesPerElf[0];
  console.log("Part 1:", highest);

  const top3 = caloriesPerElf.slice(0, 3).reduce((p, c) => p + c);
  console.log("Part 2:", top3);
}
