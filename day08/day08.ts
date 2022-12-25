const parseInput = (rawInput: string): number[][] => {
  return rawInput
    .split("\n")
    .map((v) => v.split("").map((h) => +h));
};

const part1 = (parsedInput: number[][]) => {
  const visible: Set<string> = new Set();

  const testVisibility = (
    startX: number,
    startY: number,
    directionX: -1 | 0 | 1,
    directionY: -1 | 0 | 1,
  ) => {
    let x = startX;
    let y = startY;
    let lastHeight = -1;

    while (
      y >= 0 && y < parsedInput.length && x >= 0 && x < parsedInput[y].length
    ) {
      const height = parsedInput[y][x];
      if (height > lastHeight) {
        visible.add(`${x}:${y}`);
        lastHeight = height;
      }
      x += directionX;
      y += directionY;
    }
  };

  for (let y = 0; y < parsedInput.length; y++) {
    testVisibility(0, y, 1, 0);
    testVisibility(parsedInput[y].length - 1, y, -1, 0);
  }

  for (let x = 0; x < parsedInput[0].length; x++) {
    testVisibility(x, 0, 0, 1);
    testVisibility(x, parsedInput.length - 1, 0, -1);
  }

  console.log("Part 1:", visible.size);
};

const part2 = (parsedInput: number[][]) => {
  const getVisibility = (
    startX: number,
    startY: number,
    directionX: -1 | 0 | 1,
    directionY: -1 | 0 | 1,
  ) => {
    const startHeight = parsedInput[startY][startX];

    let x = startX + directionX;
    let y = startY + directionY;
    let score = 0;

    while (
      y >= 0 && y < parsedInput.length && x >= 0 && x < parsedInput[y].length
    ) {
      score++;
      if (parsedInput[y][x] >= startHeight) {
        break;
      }
      x += directionX;
      y += directionY;
    }

    return score;
  };

  let bestVisibilityScore = 0;
  for (let y = 0; y < parsedInput.length; y++) {
    const row = parsedInput[y];
    for (let x = 0; x < row.length; x++) {
      const leftScore = getVisibility(x, y, -1, 0);
      const rightScore = getVisibility(x, y, 1, 0);
      const upScore = getVisibility(x, y, 0, -1);
      const downScore = getVisibility(x, y, 0, 1);
      const totalScore = leftScore * rightScore * upScore * downScore;
      if (totalScore > bestVisibilityScore) {
        bestVisibilityScore = totalScore;
      }
    }
  }

  console.log("Part 2:", bestVisibilityScore);
};

export default function (rawInput: string) {
  const parsedInput = parseInput(rawInput);

  part1(parsedInput);
  part2(parsedInput);
}
