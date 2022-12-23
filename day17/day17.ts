import { readFileSync } from "../utils/readfile.ts";

const parseInput = (path: string) =>
  readFileSync(path)
    .split("")
    .map((dir) => (dir === ">" ? 1 : -1)) as (-1 | 1)[];

const ROCKS = [
  { shape: ["####"], width: 4, height: 1 },
  { shape: [".#.", "###", ".#."], width: 3, height: 3 },
  { shape: ["..#", "..#", "###"].reverse(), width: 3, height: 3 },
  { shape: ["#", "#", "#", "#"], width: 1, height: 4 },
  { shape: ["##", "##"], width: 2, height: 2 },
] as const;

const getRockShape = function* (rock: typeof ROCKS[number]) {
  for (let ry = 0; ry < rock.height; ry++) {
    for (let rx = 0; rx < rock.width; rx++) {
      if (rock.shape[ry][rx] !== "#") continue;
      yield [rx, ry];
    }
  }
};

const moveRock = (
  leftEdge: number,
  bottomEdge: number,
  down: 0 | 1,
  wind: -1 | 0 | 1,
  rock: typeof ROCKS[number],
  map: (1 | undefined)[][],
) => {
  if (wind === -1 && leftEdge === 0) return false;
  if (wind === 1 && leftEdge + rock.width >= 7) return false;
  const nextBottom = bottomEdge - down;
  const nextSide = leftEdge + wind;
  if (nextBottom === 0) return false;
  for (const [rx, ry] of getRockShape(rock)) {
    const mapX = nextSide + rx;
    const mapY = nextBottom + ry;
    if (map[mapY] && map[mapY][mapX]) {
      return false;
    }
  }
  return true;
};

const commitRock = (
  leftEdge: number,
  bottomEdge: number,
  rock: typeof ROCKS[number],
  map: (1 | undefined)[][],
) => {
  for (const [rx, ry] of getRockShape(rock)) {
    if (rock.shape[ry][rx] !== "#") continue;
    const mapX = leftEdge + rx;
    const mapY = bottomEdge + ry;
    if (!map[mapY]) map[mapY] = [];
    map[mapY][mapX] = 1;
  }
};

const part1 = (path: string) => {
  const input = parseInput(path);
  let w = 0;
  let highestPoint = 0;
  const map: (1 | undefined)[][] = [];
  for (let r = 0; r < 2022; r++) {
    const rock = ROCKS[r % ROCKS.length];
    let leftEdge = 2;
    let bottomEdge = highestPoint + 4;
    while (true) {
      const wind = input[w++ % input.length];
      if (moveRock(leftEdge, bottomEdge, 0, wind, rock, map)) {
        leftEdge += wind;
      }
      if (moveRock(leftEdge, bottomEdge, 1, 0, rock, map)) {
        bottomEdge--;
      } else {
        commitRock(leftEdge, bottomEdge, rock, map);
        const topEdge = bottomEdge + rock.height - 1;
        if (topEdge > highestPoint) {
          highestPoint = topEdge;
        }
        break;
      }
    }
  }

  console.log("Part 1:", highestPoint);
};

const part2 = (path: string) => {
  const input = parseInput(path);
  const windLoop = input.length * 5;

  let w = 0;
  let highestPoint = 0;

  const map: (1 | undefined)[][] = [];

  let heightAtLastLoop = 0;
  let rocksAtLastLoop = 0;
  let loopHeightDelta = 0;
  let skippedHeight = 0;
  let _lastW = 0;
  let windLooped = false;
  const ROCK_TOTAL = 1000000000000;

  for (let r = 0; r < ROCK_TOTAL; r++) {
    if (skippedHeight === 0 && windLooped && r % 5 === 0) {
      windLooped = false;
      const newHeight = highestPoint;
      if (r > 0 && loopHeightDelta === newHeight - heightAtLastLoop) {
        const rocksPerLoop = r - rocksAtLastLoop;
        const loopsLeft = Math.floor((ROCK_TOTAL - r) / rocksPerLoop);
        r += rocksPerLoop * loopsLeft;
        skippedHeight = loopHeightDelta * loopsLeft;
      } else {
        loopHeightDelta = newHeight - heightAtLastLoop;
      }
      heightAtLastLoop = newHeight;
      rocksAtLastLoop = r;
    }
    const rock = ROCKS[r % ROCKS.length];
    let leftEdge = 2;
    let bottomEdge = highestPoint + 4;
    while (true) {
      const wind = input[w % input.length];
      w = (w + 1) % windLoop;
      if (w === 0) windLooped = true;
      _lastW = w;
      if (moveRock(leftEdge, bottomEdge, 0, wind, rock, map)) {
        leftEdge += wind;
      }
      if (moveRock(leftEdge, bottomEdge, 1, 0, rock, map)) {
        bottomEdge--;
      } else {
        commitRock(leftEdge, bottomEdge, rock, map);
        const topEdge = bottomEdge + rock.height - 1;
        if (topEdge > highestPoint) {
          highestPoint = topEdge;
        }
        break;
      }
    }
  }

  const result = (highestPoint + skippedHeight);
  console.log("Part 2:", result);
};

const run = () => {
  const inputPath = new URL("input.txt", import.meta.url).pathname;

  part1(inputPath);
  part2(inputPath);
};

export default run;
