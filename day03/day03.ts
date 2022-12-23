import { readFileSync } from "../utils/readfile.ts";

const items = [..." abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"];

const parseInput = (path: string): string[] => {
  return readFileSync(path).split("\n");
};

const part1 = (parsedInput: string[]) => {
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

  console.log("Part 1:", prioritySum);
};

const part2 = (parsedInput: string[]) => {
  let prioritySum = 0;
  for (let i = 0; i < parsedInput.length; i += 3) {
    const badgeItem = [...parsedInput[i]].find(
      (item) =>
        parsedInput[i + 1].includes(item) && parsedInput[i + 2].includes(item),
    );
    prioritySum += items.indexOf(badgeItem!);
  }

  console.log("Part 2:", prioritySum);
};

export default function (inputPath: string) {
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
}
