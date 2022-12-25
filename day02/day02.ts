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

const calculateScore = (rawInput: string, scores: ScoreMap): number => {
  return rawInput
    .split("\n")
    .map((line) => scores[line])
    .reduce((a, c) => a + c);
};

export default function (rawInput: string) {
  console.log("Part 1:", calculateScore(rawInput, scoresPart1));
  console.log("Part 2:", calculateScore(rawInput, scoresPart2));
}
