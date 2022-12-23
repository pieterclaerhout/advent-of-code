import { readFileSync } from "../utils/readfile.ts";

type Operation = ["addx", number] | ["noop"];

const parseInput = (path: string): Operation[] => {
  return readFileSync(path)
    .split("\n")
    .map((v) => {
      const [instr, val] = v.split(" ");
      return [instr, +val] as ["addx", number] | ["noop"];
    });
};

const part1 = (parsedInput: Operation[]) => {
  let cycle = 1;
  let x = 1;
  let signalSum = 0;

  const checkStrength = () => {
    if ((cycle - 20) % 40 === 0) {
      signalSum += cycle * x;
    }
  };

  for (const [instr, val] of parsedInput) {
    cycle++;
    if (instr === "noop") {
      continue;
    }
    checkStrength();
    cycle++;
    x += val;

    checkStrength();
  }

  console.log("Part 1:", signalSum);
};

const part2 = (parsedInput: Operation[]) => {
  let cycle = 1;
  let x = 1;

  const screen: string[][] = [[], [], [], [], [], []];
  const draw = () => {
    const pos = cycle - 1;
    const row = Math.floor(pos / 40);
    const drawn = Math.abs((pos % 40) - x) <= 1;
    screen[row].push(drawn ? "#" : ".");
  };

  for (const [instr, val] of parsedInput) {
    draw();
    cycle += 1;
    if (instr === "noop") {
      continue;
    }
    draw();
    cycle += 1;
    x += val;
  }

  const result = screen.map((row) => row.join("")).join("\n");
  console.log("Part 2:\n" + result);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
};

export default run;
