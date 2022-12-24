const parseInput = (rawInput: string): number[] => {
  return rawInput
    .split("\n\n")
    .map((chunk) => chunk.split("\n").map((v) => +v).reduce((p, c) => p + c))
    .sort((a, b) => b - a);
};

export default function (_inputPath: string, rawInput: string) {
  const caloriesPerElf = parseInput(rawInput);

  const highest = caloriesPerElf[0];
  console.log("Part 1:", highest);

  const top3 = caloriesPerElf.slice(0, 3).reduce((p, c) => p + c);
  console.log("Part 2:", top3);
}
