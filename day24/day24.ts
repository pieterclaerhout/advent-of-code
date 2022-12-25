enum Blizzard {
  Nothing = 0,
  Up = 1,
  Left = 2,
  Down = 4,
  Right = 8,
}

interface Position {
  row: number;
  column: number;
}

class Grid {
  grid: Blizzard[][] = [];
  minutes = 0;

  constructor(lines: string[]) {
    for (let row = 1; row < lines.length - 1; row++) {
      const rowArray: Blizzard[] = [];
      this.grid.push(rowArray);
      for (let column = 1; column < lines[row].length - 1; column++) {
        const letter = lines[row][column];
        let blizzardNode: Blizzard = Blizzard.Nothing;
        switch (letter) {
          case ".":
            break;
          case ">":
            blizzardNode = Blizzard.Right;
            break;
          case "<":
            blizzardNode = Blizzard.Left;
            break;
          case "^":
            blizzardNode = Blizzard.Up;
            break;
          case "v":
            blizzardNode = Blizzard.Down;
            break;
          default:
            throw new Error("we shouldn't get here");
        }

        rowArray.push(blizzardNode);
      }
    }
  }

  nextMinute = () => {
    const newGrid: Blizzard[][] = [];

    for (let row = 0; row < this.grid.length; row++) {
      newGrid.push(new Array(this.grid[row].length).fill(Blizzard.Nothing));
    }

    for (let row = 0; row < this.grid.length; row++) {
      for (let column = 0; column < this.grid[0].length; column++) {
        const blizzardFlag = this.grid[row][column];
        if (blizzardFlag !== Blizzard.Nothing) {
          if (blizzardFlag & Blizzard.Left) {
            const newColumn = (column - 1 + this.grid[0].length) %
              this.grid[0].length;
            newGrid[row][newColumn] |= Blizzard.Left;
          }
          if (blizzardFlag & Blizzard.Up) {
            const newRow = (row - 1 + this.grid.length) % this.grid.length;
            newGrid[newRow][column] |= Blizzard.Up;
          }
          if (blizzardFlag & Blizzard.Down) {
            const newRow = (row + 1) % this.grid.length;
            newGrid[newRow][column] |= Blizzard.Down;
          }
          if (blizzardFlag & Blizzard.Right) {
            const newColumn = (column + 1) % this.grid[0].length;
            newGrid[row][newColumn] |= Blizzard.Right;
          }
        }
      }
    }

    this.grid = newGrid;
  };

  toKey = (position: Position) =>
    String(position.row) + "," + String(position.column);

  minutesGo = () => {
    let queue: Position[] = [];
    while (true) {
      this.minutes++;
      this.nextMinute();
      if (this.grid[0][0] === Blizzard.Nothing) {
        queue.push({ row: 0, column: 0 });
      }

      const newQueue: Position[] = [];
      const visited = new Set<string>();
      for (const entry of queue) {
        if (visited.has(this.toKey(entry))) {
          continue;
        }
        if (
          entry.row === (this.grid.length - 1) &&
          entry.column === (this.grid[0].length - 1)
        ) {
          return this.minutes;
        }
        newQueue.push(...this.iterateMinute(entry));
        visited.add(this.toKey(entry));
      }
      queue = newQueue;
    }
  };

  minutesReturn = () => {
    let queue: Position[] = [];
    const endRow = this.grid.length - 1;
    const endColumn = this.grid[0].length - 1;

    while (true) {
      this.minutes++;
      this.nextMinute();
      if (this.grid[endRow][endColumn] === Blizzard.Nothing) {
        queue.push({ row: endRow, column: endColumn });
      }

      const newQueue: Position[] = [];
      const visited = new Set<string>();
      for (const entry of queue) {
        if (visited.has(this.toKey(entry))) {
          continue;
        }
        if (entry.row === 0 && entry.column === 0) {
          return this.minutes;
        }
        newQueue.push(...this.iterateMinute(entry));
        visited.add(this.toKey(entry));
      }
      queue = newQueue;
    }
  };

  iterateMinute = (position: Position) => {
    const up = position.row - 1;
    const down = position.row + 1;
    const left = position.column - 1;
    const right = position.column + 1;
    const result: Position[] = [];

    if (this.grid[position.row][position.column] === Blizzard.Nothing) {
      result.push(position);
    }
    if (up >= 0 && this.grid[up][position.column] === Blizzard.Nothing) {
      result.push({ row: up, column: position.column });
    }
    if (
      down < this.grid.length &&
      this.grid[down][position.column] === Blizzard.Nothing
    ) {
      result.push({ row: down, column: position.column });
    }
    if (left >= 0 && this.grid[position.row][left] === Blizzard.Nothing) {
      result.push({ row: position.row, column: left });
    }
    if (
      right < this.grid[0].length &&
      this.grid[position.row][right] === Blizzard.Nothing
    ) {
      result.push({ row: position.row, column: right });
    }
    return result;
  };
}

export default function (_inputPath: string, rawInput: string) {
  const grid = new Grid(rawInput.split("\n"));

  const part1 = grid.minutesGo();
  grid.minutesReturn();
  const part2 = grid.minutesGo();

  console.log("Part 1:", part1);
  console.log("Part 2:", part2);
}
