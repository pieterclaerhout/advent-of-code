import { readFileSync } from "../utils/readfile.ts";

type NumberMonkey = {
  name: string;
  type: "number";
  number: number;
};

type OperationMonkey = {
  name: string;
  type: "operation";
  operation: "+" | "*" | "/" | "-" | "=";
  a: string;
  b: string;
  number: false | number;
};

type AnyMonkey = NumberMonkey | OperationMonkey;

const sign = (value: number) => (value >= 0 ? 1 : -1);

const parseInput = (path: string): AnyMonkey[] =>
  readFileSync(path)
    .split("\n")
    .map((line) => {
      const name = line.match(/^([a-z]+):/)![1];
      const numberMatch = line.match(/: (\d+)$/);
      if (numberMatch) {
        return { name, type: "number", number: +numberMatch[1] };
      }
      const [a, operation, b] = line.split(":")[1].trim().split(" ");
      return {
        name,
        type: "operation",
        operation,
        a,
        b,
        number: false,
      } as OperationMonkey;
    });

const solveRoot = (monkeyMap: Map<string, AnyMonkey>): AnyMonkey => {
  const rootMonkey = monkeyMap.get("root") as OperationMonkey;
  while (true) {
    for (const [, monkey] of monkeyMap) {
      if (monkey.type === "number") {
        continue;
      }
      if (monkey.number !== false) {
        continue;
      }
      const a = monkeyMap.get(monkey.a)!;
      if (a.number === false) {
        continue;
      }
      const b = monkeyMap.get(monkey.b)!;
      if (b.number === false) {
        continue;
      }
      switch (monkey.operation) {
        case "+":
          monkey.number = a.number + b.number;
          break;
        case "*":
          monkey.number = a.number * b.number;
          break;
        case "-":
          monkey.number = a.number - b.number;
          break;
        case "/":
          if (a.number % b.number !== 0) {
            return rootMonkey;
          }
          monkey.number = a.number / b.number;
          break;
        case "=":
          monkey.number = a.number === b.number ? 1 : 0;
          break;
      }
    }
    if (rootMonkey.number !== false) {
      return rootMonkey;
    }
  }
};

const part1 = (path: string) => {
  const input = parseInput(path);
  const monkeyMap: Map<string, AnyMonkey> = new Map(
    input.map((m) => [m.name, m]),
  );
  const result = solveRoot(monkeyMap).number;

  console.log("Part 1:", result);
};

const part2 = (path: string) => {
  const input = parseInput(path);
  let equalDiff = 0;
  let hStep = 1;
  let lastGoodH = 0;
  for (let h = 0; h < Number.MAX_SAFE_INTEGER; h += hStep) {
    const monkeyMap: Map<string, AnyMonkey> = new Map(
      input.map((m) => [m.name, { ...m }]),
    );
    monkeyMap.get("humn")!.number = h;
    const rootMonkey = monkeyMap.get("root") as OperationMonkey;
    rootMonkey.operation = "=";
    const rootA = monkeyMap.get(rootMonkey.a)!;
    const rootB = monkeyMap.get(rootMonkey.b)!;
    solveRoot(monkeyMap);
    if (rootMonkey.number === 1) {
      console.log("Part 2:", h);
      return;
    } else if (rootMonkey.number === 0) {
      const prevEqualDiff = equalDiff;
      equalDiff = (rootA.number as number) - (rootB.number as number);
      if (prevEqualDiff !== 0) {
        if (sign(equalDiff) !== sign(prevEqualDiff)) {
          h = lastGoodH;
          hStep = Math.max(1, hStep / 100);
          equalDiff = 0;
        } else {
          lastGoodH = h;
          hStep = Math.min(1e12, hStep * 10);
        }
      }
    }
  }
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;

  part1(inputPath);
  part2(inputPath);
};

export default run;
