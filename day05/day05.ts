const part1 = (parsedInput: string[][]): string => {
  const [input, moves] = parsedInput;

  const stacks: string[][] = [];

  for (let i = 0; i < 9; i++) {
    const stack: string[] = [];

    for (let j = input.length - 2; j >= 0; j--) {
      const char = input[j][i * 4 + 1]?.trim();
      if (char) {
        stack.push(char);
      }
    }

    stacks.push(stack);
  }

  moves.forEach((move) => {
    const [no, fr, to] = move
      .split(/(move | from | to )/g)
      .filter((l) => !/[\sa-z]/.test(l) && l.length > 0)
      .map(Number);

    for (let j = 0; j < no; j++) {
      const move = stacks[fr - 1].pop();
      stacks[to - 1].push(move!);
    }
  });

  return stacks.map((s) => s.pop()).join("");
};

const part2 = (parsedInput: string[][]): string => {
  const [input, moves] = parsedInput;

  const stacks: string[][] = [];

  for (let i = 0; i < new Number(input.at(-1)?.trim().at(-1)); i++) {
    const stack: string[] = [];

    for (let j = input.length - 2; j >= 0; j--) {
      const char = input[j][i * 4 + 1]?.trim();
      if (char) {
        stack.push(char);
      }
    }

    stacks.push(stack);
  }

  moves.forEach((move) => {
    const [no, fr, to] = move
      .split(/(move | from | to )/g)
      .filter((l) => !/[\sa-z]/.test(l) && l.length > 0)
      .map(Number);

    const moving: string[] = [];
    for (let j = 0; j < no; j++) {
      moving.push(stacks[fr - 1].pop()!);
    }

    moving.reverse().forEach((m) => stacks[to - 1].push(m));
  });

  return stacks.map((s) => s.pop()).join("");
};

export default function (rawInput: string): [string, string] {
  const parsedInput = rawInput
    .split("\n\n")
    .map((p) => p.split("\n"));

  return [part1(parsedInput), part2(parsedInput)];
}
