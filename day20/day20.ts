import { readFileSync } from "../utils/readfile.ts";

type NumberContainer = { n: number };

const parseInput = (path: string, decryptKey = 1): NumberContainer[] =>
  readFileSync(path)
    .split("\n")
    .map((line) => ({ n: +line * decryptKey }));

const mix = (input: NumberContainer[], times = 1): NumberContainer[] => {
  const mixed = [...input];
  for (let i = 0; i < times; i++) {
    for (const number of input) {
      const mixedIndex = mixed.indexOf(number);
      mixed.splice(mixedIndex, 1);
      const newIndex = (mixedIndex + number.n) % mixed.length;
      mixed.splice(newIndex, 0, number);
    }
  }
  return mixed;
};

const getCoordinates = (mixed: NumberContainer[]): number => {
  const theZeroIndex = mixed.findIndex(({ n }) => n === 0);
  return (
    mixed[(theZeroIndex + 1000) % mixed.length].n +
    mixed[(theZeroIndex + 2000) % mixed.length].n +
    mixed[(theZeroIndex + 3000) % mixed.length].n
  );
};

const part1 = (path: string) => {
  const input = parseInput(path);
  const mixed = mix(input);
  const result = getCoordinates(mixed);

  console.log("Part 1:", result);
};

const part2 = (path: string) => {
  const input = parseInput(path, 811589153);
  const mixed = mix(input, 10);
  const result = getCoordinates(mixed);

  console.log("Part 2:", result);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;

  part1(inputPath);
  part2(inputPath);
};

export default run;
