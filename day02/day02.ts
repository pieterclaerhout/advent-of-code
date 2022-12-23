import { readAndSplitFileSync } from "../utils/readfile.ts";

type ScoreMap = Record<string, number>;

const scoresPart1: ScoreMap = {
  "B X": 1,
  "C Y": 2,
  "A Z": 3,
  "A X": 4,
  "B Y": 5,
  "C Z": 6,
  "C X": 7,
  "A Y": 8,
  "B Z": 9,
};

const scoresPart2: ScoreMap = {
  "B X": 1,
  "C X": 2,
  "A X": 3,
  "A Y": 4,
  "B Y": 5,
  "C Y": 6,
  "C Z": 7,
  "A Z": 8,
  "B Z": 9,
};

const calculateScore = (inputPath: string, scores: ScoreMap): number => {
  return readAndSplitFileSync(inputPath, "\n")
    .map((line) => scores[line])
    .reduce((a, c) => a + c);
};

export default function (inputPath: string) {
  console.log("Part 1:", calculateScore(inputPath, scoresPart1));
  console.log("Part 2:", calculateScore(inputPath, scoresPart2));
}
