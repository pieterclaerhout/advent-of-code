const items = [..." abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"];

const part1 = (parsedInput: string[]): number => {
  const input = parsedInput
    .map((v) => [
      v.slice(0, v.length / 2).split(""),
      v.slice(v.length / 2).split(""),
    ]);

  let prioritySum = 0;
  for (const sack of input) {
    const commonItem = sack[0].find((c) => sack[1].includes(c));
    prioritySum += items.indexOf(commonItem!);
  }

  return prioritySum;
};

const part2 = (parsedInput: string[]): number => {
  let prioritySum = 0;
  for (let i = 0; i < parsedInput.length; i += 3) {
    const badgeItem = [...parsedInput[i]].find(
      (item) =>
        parsedInput[i + 1].includes(item) && parsedInput[i + 2].includes(item),
    );
    prioritySum += items.indexOf(badgeItem!);
  }

  return prioritySum;
};

export default function (rawInput: string): [number, number] {
  const parsedInput = rawInput.split("\n");

  return [part1(parsedInput), part2(parsedInput)];
}
