import { readFileSync } from "../utils/readfile.ts";

type Round = ["A" | "B" | "C", "X" | "Y" | "Z"];

const parseInput = (path: string): Round[] => {
  return readFileSync(path)
    .split("\n")
    .map((v) => v.split(" ") as Round);
};

const part1 = (parsedInput: Round[]) => {
  const myMoves = {
    X: { value: 1, A: 3, B: 0, C: 6 },
    Y: { value: 2, A: 6, B: 3, C: 0 },
    Z: { value: 3, A: 0, B: 6, C: 3 },
  };

  let score = 0;
  for (const round of parsedInput) {
    score += myMoves[round[1]].value + myMoves[round[1]][round[0]];
  }

  console.log("Part 1:", score);
};

const part2 = (parsedInput: Round[]) => {
  const theirMoves = {
    A: { X: 3, Y: 1, Z: 2 },
    B: { X: 1, Y: 2, Z: 3 },
    C: { X: 2, Y: 3, Z: 1 },
  };

  const myMoveScores = {
    X: 0,
    Y: 3,
    Z: 6,
  };

  let score = 0;
  for (const round of parsedInput) {
    score += theirMoves[round[0]][round[1]] + myMoveScores[round[1]];
  }

  console.log("Part 2:", score);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
};

export default run;
