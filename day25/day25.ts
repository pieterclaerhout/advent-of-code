const snafuToDecimal = (input: string): number => {
  const mapping: Record<string, number> = {
    "=": -2,
    "-": -1,
    "0": 0,
    "1": 1,
    "2": 2,
  };

  let sum = 0;
  for (const line of input.split("\n")) {
    let n = 0;
    for (const c of line.split("")) {
      n = 5 * n + mapping[c];
    }
    sum += n;
  }
  return sum;
};

const decimalToSnafu = (sum: number): string => {
  const keys: string[] = "=-012".split("");

  let snafu = "";
  while (sum > 0) {
    snafu = keys[(sum + 2) % 5] + snafu;
    sum = Math.floor((sum + 2) / 5);
  }

  return snafu;
};

export default function (rawInput: string) {
  console.log("Part 1:", snafuToDecimal(rawInput));
  console.log("Part 2:", decimalToSnafu(snafuToDecimal(rawInput)));
}
