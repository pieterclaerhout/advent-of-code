type NumberContainer = { n: number };

const parseInput = (rawInput: string, decryptKey = 1): NumberContainer[] =>
  rawInput
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

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);
  const mixed = mix(input);
  return getCoordinates(mixed);
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput, 811589153);
  const mixed = mix(input, 10);
  return getCoordinates(mixed);
};

export default function (rawInput: string): [number, number] {
  return [part1(rawInput), part2(rawInput)];
}
