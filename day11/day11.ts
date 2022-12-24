import { readFileSync } from "../utils/readfile.ts";

const FIRST_PART_ROUNDS = 20;
const SECOND_PART_ROUNDS = 10000;

interface Monkey {
  items: number[];
  operation: string;
  test: number;
  testTrue: number;
  testFalse: number;
  activity: number;
}

const getItems = (monkey: string): number[] =>
  monkey.split("\n")[1]
    .replace("Starting items: ", "")
    .replaceAll(", ", ",")
    .split(",")
    .map((item) => parseInt(item));

const getOperation = (monkey: string): string =>
  monkey.split("\n")[2]
    .replace("Operation: new = ", "").trim();

const getTest = (monkey: string): number =>
  parseInt(
    monkey.split("\n")[3]
      .replace("Test: divisible by ", ""),
  );

const getTestTrue = (monkey: string): number =>
  parseInt(
    monkey.split("\n")[4]
      .replace("If true: throw to monkey", ""),
  );

const getTestFalse = (monkey: string): number =>
  parseInt(
    monkey.split("\n")[5]
      .replace("If false: throw to monkey", ""),
  );

const getMonkeys = (rawInput: string): Monkey[] => {
  return rawInput
    .split("\n\n")
    .map((monkey: string) => {
      return {
        items: getItems(monkey),
        operation: getOperation(monkey),
        test: getTest(monkey),
        testTrue: getTestTrue(monkey),
        testFalse: getTestFalse(monkey),
        activity: 0,
      };
    });
};

const runMonkeyBusiness = (rawInput: string, rounds = FIRST_PART_ROUNDS) => {
  const monkeys = getMonkeys(rawInput);

  const multiModule = monkeys.map((m) => m.test)
    .reduce((a, b) => a * b, 1);

  for (let i = 0; i < rounds; i++) {
    for (let j = 0; j < monkeys.length; j++) {
      monkeys[j].items.forEach((item) => {
        const operation = monkeys[j].operation.replaceAll(
          "old",
          item.toString(),
        );
        const newItemValue = (rounds === FIRST_PART_ROUNDS)
          ? Math.floor(Number(eval(operation)) / 3)
          : Number(eval(operation)) % multiModule;

        if (newItemValue % monkeys[j].test === 0) {
          monkeys[monkeys[j].testTrue].items.push(newItemValue);
        } else {
          monkeys[monkeys[j].testFalse].items.push(newItemValue);
        }
        monkeys[j].activity++;
      });
      monkeys[j].items = [];
    }
  }

  return monkeys.map((m) => m.activity)
    .sort((a, b) => a - b)
    .slice(-2)
    .reduce((a, b) => a * b, 1);
};

const part1 = (rawInput: string) => {
  const result = runMonkeyBusiness(rawInput, FIRST_PART_ROUNDS);
  console.log("Part 1:", result);
};

const part2 = (rawInput: string) => {
  const result = runMonkeyBusiness(rawInput, SECOND_PART_ROUNDS);
  console.log("Part 2:", result);
};

export default function (_inputPath: string, rawInput: string) {
  part1(rawInput);
  part2(rawInput);
}
