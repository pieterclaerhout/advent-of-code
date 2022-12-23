import { readFileSync } from "../utils/readfile.ts";

type ParsedInput = [[number, number], [number, number]][];

const parseInput = (path: string): ParsedInput => {
  return readFileSync(path)
    .split("\n")
    .map((v) =>
      v.split(",").map((r) => r.split("-").map((s) => +s))
    ) as ParsedInput;
};

const part1 = (parsedInput: ParsedInput) => {
  const result = parsedInput
    .filter((input) => {
      const [[aStart, aEnd], [bStart, bEnd]] = input;
      return (
        (aStart <= bStart && aEnd >= bEnd) || (bStart <= aStart && bEnd >= aEnd)
      );
    })
    .length;

  console.log("Part 1:", result);
};

const part2 = (parsedInput: ParsedInput) => {
  const result = parsedInput
    .filter((input) => {
      const [[aStart, aEnd], [bStart, bEnd]] = input;
      for (let ai = aStart; ai <= aEnd; ai++) {
        if (ai >= bStart && ai <= bEnd) {
          return true;
        }
      }
      return false;
    })
    .length;

  console.log("Part 2:", result);
};

export default function (inputPath: string) {
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
}
