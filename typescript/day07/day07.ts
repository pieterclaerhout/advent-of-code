const parseInput = (rawInput: string): Map<string, number> => {
  const inputLines = rawInput
    .split("\n")
    .map((v) => v.split(" "));

  const dirSizeMap: Map<string, number> = new Map();
  const currentDir: string[] = [];
  dirSizeMap.set("/", 0);
  for (const words of inputLines) {
    if (words[0] === "$") {
      if (words[1] === "cd") {
        const changeTo = words[2];
        if (changeTo === "..") {
          currentDir.pop();
        } else if (changeTo === "/") {
          currentDir.length = 0;
        } else {
          currentDir.push(changeTo);
        }
      }
    } else if (words[0] === "dir") {
      dirSizeMap.set([...currentDir, words[1]].join("/"), 0);
    } else {
      const path = [...currentDir];
      while (path.length > 0) {
        dirSizeMap.set(
          path.join("/"),
          dirSizeMap.get(path.join("/"))! + +words[0],
        );
        path.pop();
      }
      dirSizeMap.set("/", dirSizeMap.get("/")! + +words[0]);
    }
  }
  return dirSizeMap;
};

const part1 = (parsedInput: Map<string, number>): number => {
  let totalSizeOfDirsUnder100000 = 0;
  for (const [_path, dirSize] of parsedInput) {
    if (dirSize <= 100000) {
      totalSizeOfDirsUnder100000 += dirSize;
    }
  }
  return totalSizeOfDirsUnder100000;
};

const part2 = (parsedInput: Map<string, number>): number => {
  const needToFree = 30000000 - (70000000 - parsedInput.get("/")!);

  let closestSize = Infinity;
  for (const [_path, dirSize] of parsedInput) {
    if (dirSize >= needToFree && dirSize < closestSize) {
      closestSize = dirSize;
    }
  }

  return closestSize;
};

export default function (rawInput: string): [number, number] {
  const parsedInput = parseInput(rawInput);

  return [part1(parsedInput), part2(parsedInput)];
}
