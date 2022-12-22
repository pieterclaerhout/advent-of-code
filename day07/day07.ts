import { readFileSync } from "../utils/readfile.ts";

const parseInput = (path: string): Map<string, number> => {
  const inputLines = readFileSync(path)
    .split('\n')
    .map((v) => v.split(' '));

  const dirSizeMap: Map<string, number> = new Map();
  const currentDir: string[] = [];
  dirSizeMap.set('/', 0);
  for (const words of inputLines) {
    if (words[0] === '$') {
      if (words[1] === 'cd') {
        const changeTo = words[2];
        if (changeTo === '..') {
          currentDir.pop();
        } else if (changeTo === '/') {
          currentDir.length = 0;
        } else {
          currentDir.push(changeTo);
        }
      }
    } else if (words[0] === 'dir') {
      dirSizeMap.set([...currentDir, words[1]].join('/'), 0);
    } else {
      const path = [...currentDir];
      while (path.length > 0) {
        dirSizeMap.set(
          path.join('/'),
          dirSizeMap.get(path.join('/'))! + +words[0]
        );
        path.pop();
      }
      dirSizeMap.set('/', dirSizeMap.get('/')! + +words[0]);
    }
  }
  return dirSizeMap;
}

const part1 = (parsedInput: Map<string, number>) => {
  let totalSizeOfDirsUnder100000 = 0;
  for (const [_path, dirSize] of parsedInput) {
    if (dirSize <= 100000) {
      totalSizeOfDirsUnder100000 += dirSize;
    }
  }

  console.log("Part 1:",  totalSizeOfDirsUnder100000);
}

const part2 = (parsedInput: Map<string, number>) => {
  const needToFree = 30000000 - (70000000 - parsedInput.get('/')!);

  let closestSize = Infinity;
  for (const [_path, dirSize] of parsedInput) {
    if (dirSize >= needToFree && dirSize < closestSize) {
      closestSize = dirSize;
    }
  }

  console.log("Part 2:",  closestSize);
}

const run = () => {
  const inputPath = new URL('input.txt', import.meta.url).pathname;
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
}

export default run;
