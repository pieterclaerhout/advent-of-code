type NumberOrArray = number | NumberOrArray[];

const parseInput = (rawInput: string) =>
  rawInput
    .split("\n\n")
    .map((pair) => {
      const [packet1, packet2] = pair.split("\n").map((p) => JSON.parse(p));
      return [packet1, packet2] as [NumberOrArray, NumberOrArray];
    });

const compareValues = (a: NumberOrArray, b: NumberOrArray): -1 | 0 | 1 => {
  if (!Array.isArray(a) && !Array.isArray(b)) {
    if (a < b) {
      return -1;
    }
    if (b < a) {
      return 1;
    }
    return 0;
  }
  if (!Array.isArray(b)) {
    b = [b];
  }
  if (!Array.isArray(a)) {
    a = [a];
  }
  for (let i = 0; i < Math.max(a.length, b.length); i++) {
    if (i + 1 > a.length) {
      return -1;
    } else if (i + 1 > b.length) {
      return 1;
    }
    const result = compareValues(a[i], b[i]);
    if (result !== 0) {
      return result;
    }
  }
  return 0;
};

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput);

  let indexSum = 0;
  for (let p = 0; p < input.length; p++) {
    if (compareValues(...input[p]) === -1) {
      indexSum += p + 1;
    }
  }

  return indexSum;
};

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput);

  const flatPackets = [];
  for (const pair of input) {
    flatPackets.push(...pair);
  }

  const divider1 = [[2]];
  const divider2 = [[6]];
  flatPackets.push(divider1, divider2);
  flatPackets.sort(compareValues);

  return (
    (flatPackets.indexOf(divider1) + 1) *
    (flatPackets.indexOf(divider2) + 1)
  );
};

export default function (rawInput: string): [number, number] {
  return [part1(rawInput), part2(rawInput)];
}
