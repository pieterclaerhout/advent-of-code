const parseInput = (rawInput: string): number[] => {
  return rawInput
    .split("\n\n")
    .map((chunk) =>
      chunk.split("\n")
        .map((v) => +v)
        .reduce((p, c) => p + c)
    )
    .sort((a, b) => b - a);
};

export default function (rawInput: string): [number, number] {
  const caloriesPerElf = parseInput(rawInput);

  const highest = caloriesPerElf[0];
  const top3 = caloriesPerElf.slice(0, 3).reduce((p, c) => p + c);

  return [highest, top3];
}
