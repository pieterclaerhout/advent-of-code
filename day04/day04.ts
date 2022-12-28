type Range = [number, number];
type Pair = [Range, Range];

const isContained = (pair: Pair): boolean => {
  const [left, right] = pair;
  if (left[0] <= right[0] && right[1] <= left[1]) {
    return true;
  }
  if (right[0] <= left[0] && left[1] <= right[1]) {
    return true;
  }
  return false;
};

const isOverlapping = (pair: Pair): boolean => {
  const [left, right] = pair;
  let [min, max] = [right, left];
  if (left[0] < right[0]) {
    [min, max] = [left, right];
  }
  return min[1] >= max[0];
};

const part1 = (parsedInput: Pair[]): number => {
  return parsedInput.filter(isContained).length;
};

const part2 = (parsedInput: Pair[]): number => {
  return parsedInput.filter(isOverlapping).length;
};

export default function (rawInput: string): [number, number] {
  const parsedInput = rawInput
    .split("\n")
    .map((v) => v.split(",").map((r) => r.split("-").map((s) => +s)) as Pair);

  return [part1(parsedInput), part2(parsedInput)];
}
