type Vector = [row: number, column: number];
type ElfMove = [from: Vector, to: Vector];
type PotentialMove = (row: number, column: number) => ElfMove | null;

class Board {
  private OFFSET = 50;
  private board: boolean[][];

  private elves = new Map<number, Vector>();

  public resultPart1 = 0;
  public resultPart2 = 0;

  constructor(size: number, input: string[]) {
    this.board = new Array(size)
      .fill(undefined)
      .map(() => new Array(size).fill(false));

    for (let row = 0; row < input.length; ++row) {
      for (let column = 0; column < input[row].length; ++column) {
        if (input[row][column] === "#") {
          this.addElf(row, column);
        }
      }
    }

    for (let round = 0; round < 1000; ++round) {
      const moves: Array<ElfMove> = [];
      const moveDupeChecker = new Map<number, number>();
      const potentialMoves = this.potentialMoves();

      for (const [row, column] of this.elves.values()) {
        if (this.hasNoPotentialMoves(row, column)) {
          continue;
        }

        for (let moveIndex = round; moveIndex < round + 4; ++moveIndex) {
          const tryMove = potentialMoves[moveIndex % potentialMoves.length](
            row,
            column,
          );
          if (tryMove === null) {
            continue;
          }
          moves.push(tryMove);
          const position = this.encodePosition(tryMove[1][0], tryMove[1][1]);
          moveDupeChecker.set(
            position,
            (moveDupeChecker.get(position) ?? 0) + 1,
          );
          break;
        }
      }

      for (const move of moves) {
        const position = this.encodePosition(move[1][0], move[1][1]);
        if (moveDupeChecker.get(position) === 1) {
          this.removeElf(move[0][0], move[0][1]);
          this.addElf(move[1][0], move[1][1]);
        }
      }

      if (round === 9) {
        const rows = [...this.elves.values()].map(([row]) => row);
        const columns = [...this.elves.values()].map(([_, column]) => column);
        const emptyTiles = (Math.max(...rows) - Math.min(...rows) + 1) *
            (Math.max(...columns) - Math.min(...columns) + 1) -
          this.elves.size;

        this.resultPart1 = emptyTiles;
      }

      if (moves.length === 0) {
        this.resultPart2 = round + 1;
        break;
      }
    }
  }

  private addElf(row: number, column: number): void {
    this.board[row + this.OFFSET][column + this.OFFSET] = true;
    this.elves.set(this.encodePosition(row, column), [row, column]);
  }

  private removeElf(row: number, column: number): void {
    this.board[row + this.OFFSET][column + this.OFFSET] = false;
    this.elves.delete(this.encodePosition(row, column));
  }

  private hasElf(row: number, column: number): boolean {
    return this.board[row + this.OFFSET][column + this.OFFSET];
  }

  private hasNoPotentialMoves(row: number, column: number): boolean {
    return (
      !this.hasElf(row - 1, column - 1) &&
      !this.hasElf(row - 1, column) &&
      !this.hasElf(row - 1, column + 1) &&
      !this.hasElf(row, column - 1) &&
      !this.hasElf(row, column + 1) &&
      !this.hasElf(row + 1, column - 1) &&
      !this.hasElf(row + 1, column) &&
      !this.hasElf(row + 1, column + 1)
    );
  }

  private hasNoElves(
    row1: number,
    col1: number,
    row2: number,
    col2: number,
    row3: number,
    col3: number,
  ): boolean {
    return !this.hasElf(row1, col1) && !this.hasElf(row2, col2) &&
      !this.hasElf(row3, col3);
  }

  private potentialMoves(): Array<PotentialMove> {
    return [
      (row, col) =>
        this.hasNoElves(row - 1, col - 1, row - 1, col, row - 1, col + 1)
          ? [[row, col], [row - 1, col]]
          : null,
      (row, col) =>
        this.hasNoElves(row + 1, col - 1, row + 1, col, row + 1, col + 1)
          ? [[row, col], [row + 1, col]]
          : null,
      (row, col) =>
        this.hasNoElves(row + 1, col - 1, row, col - 1, row - 1, col - 1)
          ? [[row, col], [row, col - 1]]
          : null,
      (row, col) =>
        this.hasNoElves(row + 1, col + 1, row, col + 1, row - 1, col + 1)
          ? [[row, col], [row, col + 1]]
          : null,
    ];
  }

  private encodePosition = (row: number, column: number): number => {
    const encodeOffset = (1 << 16) / 2;
    return ((row + encodeOffset) << 16) | (column + encodeOffset);
  };
}

export default function (rawInput: string) {
  const board = new Board(200, rawInput.split("\n"));

  console.log("Part 1:", board.resultPart1);
  console.log("Part 2:", board.resultPart2);
}
