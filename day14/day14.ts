const parseInput = (rawInput: string) =>
  rawInput
    .split("\n")
    .map((line) =>
      line
        .split(" -> ")
        .map((coord) => coord.split(",").map((v) => +v) as [number, number])
    );

const buildMap = (lines: [number, number][][]) => {
  const map: ("X" | "S")[][] = [];
  for (const line of lines) {
    let prevX: number, prevY: number;
    let firstPoint = true;
    for (const [toX, toY] of line) {
      if (firstPoint) {
        firstPoint = false;
      } else {
        const xA = prevX! < toX ? prevX! : toX;
        const xB = prevX! < toX ? toX : prevX!;
        const yA = prevY! < toY ? prevY! : toY;
        const yB = prevY! < toY ? toY : prevY!;
        for (let x = xA; x <= xB; x++) {
          for (let y = yA; y <= yB; y++) {
            if (!map[y]) map[y] = [];
            map[y][x] = "X";
          }
        }
      }
      prevX = toX;
      prevY = toY;
    }
  }
  return map;
};

const simulateSand = (map: ("X" | "S")[][], ground?: number): number => {
  let restingSand = 0;
  while (true) {
    const sand = [500, 0];
    while (true) {
      if (ground && sand[1] === ground - 1) {
        restingSand++;
        if (!map[sand[1]]) {
          map[sand[1]] = [];
        }
        map[sand[1]][sand[0]] = "S";
        break;
      }
      const nextY = map[sand[1] + 1];
      const down = nextY && nextY[sand[0]];
      if (!down) {
        sand[1]++;
        if (sand[1] > map.length) {
          return restingSand;
        }
        continue;
      }
      const downLeft = nextY[sand[0] - 1];
      if (!downLeft) {
        sand[1]++;
        sand[0]--;
        continue;
      }
      const downRight = nextY[sand[0] + 1];
      if (!downRight) {
        sand[1]++;
        sand[0]++;
        continue;
      }
      restingSand++;
      if (sand[0] === 500 && sand[1] === 0) {
        return restingSand;
      }
      if (!map[sand[1]]) {
        map[sand[1]] = [];
      }
      map[sand[1]][sand[0]] = "S";
      break;
    }
  }
};

const part1 = (rawInput: string): number => {
  const input = parseInput(rawInput);
  const map = buildMap(input);
  return simulateSand(map);
};

const part2 = (rawInput: string): number => {
  const input = parseInput(rawInput);
  const map = buildMap(input);
  const groundY = map.length + 1;
  return simulateSand(map, groundY);
};

export default function (rawInput: string): [number, number] {
  return [part1(rawInput), part2(rawInput)];
}
