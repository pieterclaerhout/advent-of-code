import { readFileSync } from "../utils/readfile.ts";

type XY = [number, number];

type Elf = {
  at: XY;
  propose: string | null;
};

const norths: XY[] = [
  [-1, -1],
  [0, -1],
  [1, -1],
];

const souths: XY[] = [
  [-1, 1],
  [0, 1],
  [1, 1],
];

const wests: XY[] = [
  [-1, -1],
  [-1, 0],
  [-1, 1],
];

const easts: XY[] = [
  [1, -1],
  [1, 0],
  [1, 1],
];

const allAround: XY[] = [
  ...norths,
  ...souths,
  ...wests,
  ...easts,
];

const grid = (xy: XY) => xy[0] + ":" + xy[1];
const xy = (grid: string): XY => grid.split(":").map((v) => +v) as XY;

const parseInput = (inputPath: string): Map<string, Elf> => {
  const elves: Map<string, Elf> = new Map();
  readFileSync(inputPath)
    .split("\n")
    .forEach((line, y) => {
      [...line].forEach((char, x) => {
        if (char === "#") {
          elves.set(grid([x, y]), { at: [x, y], propose: null });
        }
      });
    });
  return elves;
};

const addGrid = (a: XY, b: XY): XY => [a[0] + b[0], a[1] + b[1]];

const checkEmpty = (elves: Map<string, Elf>, from: XY, dirs: XY[]) =>
  dirs.every((dir) => !elves.has(grid(addGrid(from, dir))));

const doRound = (elves: Map<string, Elf>, dirChecks: XY[][]): boolean => {
  let settled = true;
  const proposeMap: Map<string, number> = new Map();
  for (const [, elf] of elves) {
    const alone = checkEmpty(elves, elf.at, allAround);
    if (alone) continue;
    settled = false;
    for (const dirCheck of dirChecks) {
      if (checkEmpty(elves, elf.at, dirCheck)) {
        elf.propose = grid(addGrid(elf.at, dirCheck[1]));
        proposeMap.set(elf.propose, (proposeMap.get(elf.propose) ?? 0) + 1);
        break;
      }
    }
  }
  if (settled) return true;
  for (const [elfPrevGrid, elf] of elves) {
    if (elf.propose === null) continue;
    if (proposeMap.get(elf.propose)! === 1) {
      elf.at = xy(elf.propose);
      elves.delete(elfPrevGrid);
      elves.set(grid(elf.at), elf);
    }
    elf.propose = null;
  }
  dirChecks.push(dirChecks.shift()!);
  return false;
};

const part1 = (inputPath: string): void => {
  const elves = parseInput(inputPath);
  const dirChecks = [norths, souths, wests, easts];
  for (let i = 0; i < 10; i++) {
    doRound(elves, dirChecks);
  }

  let minX = Infinity;
  let minY = Infinity;
  let maxX = -Infinity;
  let maxY = -Infinity;
  for (const [, { at: [x, y] }] of elves) {
    if (x < minX) minX = x;
    if (x > maxX) maxX = x;
    if (y < minY) minY = y;
    if (y > maxY) maxY = y;
  }
  const rectArea = (maxX - minX + 1) * (maxY - minY + 1);
  const result = (rectArea - elves.size);

  console.log("Part 1:", result);
};

const part2 = (inputPath: string): void => {
  const elves = parseInput(inputPath);
  const dirChecks = [norths, souths, wests, easts];

  let round = 1;
  while (true) {
    const settled = doRound(elves, dirChecks);
    if (settled) {
      break;
    }
    round++;
  }

  console.log("Part 2:", round);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;

  part1(inputPath);
  part2(inputPath);
};

export default run;
