const part1 = (rawInput: string): number => {
  type Player = {
    pos: number[];
    direction: number;
  };

  type Instruction = string | number;

  const getNewDirection = (oldDir: number, instruction: string): number => {
    if (instruction === "R") {
      if (oldDir === 3) {
        return 0;
      }
      return oldDir + 1;
    } else {
      if (oldDir === 0) {
        return 3;
      }
      return oldDir - 1;
    }
  };

  const getPosAfterSteps = (
    player: Player,
    steps: number,
    map: string[][],
  ): number[] => {
    let currPos = player.pos;
    let stepsLeft = steps;
    while (stepsLeft > 0) {
      if (player.direction === 0) { // Go right
        let nextPos = [currPos[0], currPos[1] + 1];
        if (
          nextPos[1] >= map[0].length || map[nextPos[0]][nextPos[1]] === " "
        ) {
          nextPos = [nextPos[0], map[nextPos[0]].findIndex((l) => l !== " ")];
        }
        if (map[nextPos[0]][nextPos[1]] === "#") {
          return currPos;
        }
        currPos = nextPos;
        stepsLeft--;
      } else if (player.direction === 2) { // Go left
        let nextPos = [currPos[0], currPos[1] - 1];
        if (nextPos[1] < 0 || map[nextPos[0]][nextPos[1]] === " ") {
          nextPos = [
            nextPos[0],
            map[nextPos[0]].reduce(
              (prev, curr, idx) => curr != " " ? idx : prev,
              0,
            ),
          ];
        }
        if (map[nextPos[0]][nextPos[1]] === "#") {
          return currPos;
        }
        currPos = nextPos;
        stepsLeft--;
      } else if (player.direction === 1) { // Go down
        let nextPos = [currPos[0] + 1, currPos[1]];
        if (nextPos[0] >= map.length || map[nextPos[0]][nextPos[1]] === " ") {
          nextPos = [map.findIndex((l) => l[nextPos[1]] !== " "), nextPos[1]];
        }
        if (map[nextPos[0]][nextPos[1]] === "#") {
          return currPos;
        }
        currPos = nextPos;
        stepsLeft--;
      } else { // Go up
        let nextPos = [currPos[0] - 1, currPos[1]];
        if (nextPos[0] < 0 || map[nextPos[0]][nextPos[1]] === " ") {
          nextPos = [
            map.reduce(
              (prev, curr, idx) => curr[nextPos[1]] != " " ? idx : prev,
              0,
            ),
            nextPos[1],
          ];
        }
        if (map[nextPos[0]][nextPos[1]] === "#") {
          return currPos;
        }
        currPos = nextPos;
        stepsLeft--;
      }
    }
    return currPos;
  };

  const getNextPlayer = (
    player: Player,
    instruction: Instruction,
    map: string[][],
  ): Player => {
    if (typeof instruction === "string") {
      return {
        ...player,
        direction: getNewDirection(player.direction, instruction),
      };
    }

    return {
      ...player,
      pos: getPosAfterSteps(player, instruction, map),
    };
  };

  const input = rawInput
    .split("\n")
    .filter((l) => l.length !== 0);

  const movement = input[input.length - 1]
    .match(/\d+|[A-Z]/g)
    ?.map((instr) => {
      if (instr.match(/\d+/)) {
        return Number(instr);
      }
      return instr;
    }) as Instruction[];

  const height = input.length - 1;
  const width = Math.max(...input.slice(0, height).map((l) => l.length));

  const map = input
    .slice(0, height)
    .map((l) => {
      return l.padEnd(width, " ").split("");
    });

  const startingPlayer: Player = {
    pos: [0, map[0].findIndex((l) => l === ".")],
    direction: 0,
  };

  const finalPlayer = movement.reduce(
    (player: Player, instruction: Instruction) =>
      getNextPlayer(player, instruction, map),
    startingPlayer,
  );

  const answer = 1000 * (finalPlayer.pos[0] + 1) +
    4 * (finalPlayer.pos[1] + 1) +
    finalPlayer.direction;

  return answer;
};

const part2 = (rawInput: string): number => {
  type Location = {
    face: number;
    pos: number[];
  };

  type Player = {
    loc: Location;
    direction: number;
  };

  type Face = string[][];

  type Instruction = string | number;

  const offsetForFace = (face: number) => {
    switch (face) {
      case 0:
        return [0, 50];
      case 1:
        return [0, 100];
      case 2:
        return [50, 50];
      case 3:
        return [100, 50];
      case 4:
        return [100, 0];
      case 5:
        return [150, 0];
    }
    return [0, 0];
  };

  const getNewDirection = (oldDir: number, instruction: string): number => {
    if (instruction === "R") {
      if (oldDir === 3) {
        return 0;
      }
      return oldDir + 1;
    } else {
      if (oldDir === 0) {
        return 3;
      }
      return oldDir - 1;
    }
  };

  const getPlayerAfterSteps = (
    player: Player,
    steps: number,
    map: Face[],
  ): Player => {
    let currPos = player.loc.pos;
    let currDir = player.direction;
    let currFace = player.loc.face;
    let stepsLeft = steps;
    while (stepsLeft > 0) {
      if (currDir === 0) { // Go right
        let nextPos = [currPos[0], currPos[1] + 1];
        let nextDir = currDir;
        let nextFace = currFace;
        if (nextPos[1] >= 50) {
          if (currFace === 0) {
            nextPos = [nextPos[0], 0];
            nextFace = 1;
            nextDir = 0;
          } else if (currFace === 1) {
            nextPos = [49 - nextPos[0], 49];
            nextFace = 3;
            nextDir = 2;
          } else if (currFace === 2) {
            nextPos = [49, nextPos[0]];
            nextFace = 1;
            nextDir = 3;
          } else if (currFace === 3) {
            nextPos = [49 - nextPos[0], 49];
            nextFace = 1;
            nextDir = 2;
          } else if (currFace === 4) {
            nextPos = [nextPos[0], 0];
            nextFace = 3;
            nextDir = 0;
          } else if (currFace === 5) {
            nextPos = [49, nextPos[0]];
            nextFace = 3;
            nextDir = 3;
          }
        }
        if (map[nextFace][nextPos[0]][nextPos[1]] === "#") {
          return {
            loc: {
              pos: currPos,
              face: currFace,
            },
            direction: currDir,
          } as Player;
        }
        currPos = nextPos;
        currDir = nextDir;
        currFace = nextFace;
        stepsLeft--;
      } else if (currDir === 2) { // Go left
        let nextPos = [currPos[0], currPos[1] - 1];
        let nextDir = currDir;
        let nextFace = currFace;
        if (nextPos[1] < 0) {
          if (currFace === 0) {
            nextPos = [49 - nextPos[0], 0];
            nextFace = 4;
            nextDir = 0;
          } else if (currFace === 1) {
            nextPos = [nextPos[0], 49];
            nextFace = 0;
            nextDir = 2;
          } else if (currFace === 2) {
            nextPos = [0, nextPos[0]];
            nextFace = 4;
            nextDir = 1;
          } else if (currFace === 3) {
            nextPos = [nextPos[0], 49];
            nextFace = 4;
            nextDir = 2;
          } else if (currFace === 4) {
            nextPos = [49 - nextPos[0], 0];
            nextFace = 0;
            nextDir = 0;
          } else if (currFace === 5) {
            nextPos = [0, nextPos[0]];
            nextFace = 0;
            nextDir = 1;
          }
        }
        if (map[nextFace][nextPos[0]][nextPos[1]] === "#") {
          return {
            loc: {
              pos: currPos,
              face: currFace,
            },
            direction: currDir,
          } as Player;
        }
        currPos = nextPos;
        currDir = nextDir;
        currFace = nextFace;
        stepsLeft--;
      } else if (currDir === 1) { // Go down
        let nextPos = [currPos[0] + 1, currPos[1]];
        let nextDir = currDir;
        let nextFace = currFace;
        if (nextPos[0] >= 50) {
          if (currFace === 0) {
            nextPos = [0, nextPos[1]];
            nextFace = 2;
            nextDir = 1;
          } else if (currFace === 1) {
            nextPos = [nextPos[1], 49];
            nextFace = 2;
            nextDir = 2;
          } else if (currFace === 2) {
            nextPos = [0, nextPos[1]];
            nextFace = 3;
            nextDir = 1;
          } else if (currFace === 3) {
            nextPos = [nextPos[1], 49];
            nextFace = 5;
            nextDir = 2;
          } else if (currFace === 4) {
            nextPos = [0, nextPos[1]];
            nextFace = 5;
            nextDir = 1;
          } else if (currFace === 5) {
            nextPos = [0, nextPos[1]];
            nextFace = 1;
            nextDir = 1;
          }
        }
        if (map[nextFace][nextPos[0]][nextPos[1]] === "#") {
          return {
            loc: {
              pos: currPos,
              face: currFace,
            },
            direction: currDir,
          } as Player;
        }
        currPos = nextPos;
        currDir = nextDir;
        currFace = nextFace;
        stepsLeft--;
      } else { // Go up
        let nextPos = [currPos[0] - 1, currPos[1]];
        let nextDir = currDir;
        let nextFace = currFace;
        if (nextPos[0] < 0) {
          if (currFace === 0) {
            nextPos = [nextPos[1], 0];
            nextFace = 5;
            nextDir = 0;
          } else if (currFace === 1) {
            nextPos = [49, nextPos[1]];
            nextFace = 5;
            nextDir = 3;
          } else if (currFace === 2) {
            nextPos = [49, nextPos[1]];
            nextFace = 0;
            nextDir = 3;
          } else if (currFace === 3) {
            nextPos = [49, nextPos[1]];
            nextFace = 2;
            nextDir = 3;
          } else if (currFace === 4) {
            nextPos = [nextPos[1], 0];
            nextFace = 2;
            nextDir = 0;
          } else if (currFace === 5) {
            nextPos = [49, nextPos[1]];
            nextFace = 4;
            nextDir = 3;
          }
        }
        if (map[nextFace][nextPos[0]][nextPos[1]] === "#") {
          return {
            loc: {
              pos: currPos,
              face: currFace,
            },
            direction: currDir,
          } as Player;
        }
        currPos = nextPos;
        currDir = nextDir;
        currFace = nextFace;
        stepsLeft--;
      }
    }
    return {
      loc: {
        pos: currPos,
        face: currFace,
      },
      direction: currDir,
    } as Player;
  };

  const getNextPlayer = (
    player: Player,
    instruction: Instruction,
    map: Face[],
  ): Player => {
    if (typeof instruction === "string") {
      return {
        ...player,
        direction: getNewDirection(player.direction, instruction),
      };
    }

    return getPlayerAfterSteps(player, instruction, map);
  };

  const input = rawInput
    .split("\n")
    .filter((l) => l.length !== 0);

  const movement = input[input.length - 1]
    .match(/\d+|[A-Z]/g)
    ?.map((instr) => {
      if (instr.match(/\d+/)) {
        return Number(instr);
      }
      return instr;
    }) as Instruction[];

  const height = input.length - 1;

  const inputMap = input
    .slice(0, height)
    .map((l) => {
      return l.replace(/\s/g, "");
    });

  const map: Face[] = [
    inputMap.slice(0, 50).map((l) => l.slice(0, 50)),
    inputMap.slice(0, 50).map((l) => l.slice(50, 100)),
    inputMap.slice(50, 100).map((l) => l.slice(0, 50)),
    inputMap.slice(100, 150).map((l) => l.slice(50, 100)),
    inputMap.slice(100, 150).map((l) => l.slice(0, 50)),
    inputMap.slice(150, 200).map((l) => l.slice(0, 50)),
  ].map((m) => m.map((l) => l.split("")));

  const startingPlayer: Player = {
    loc: {
      pos: [0, map[0][0].findIndex((l) => l === ".")],
      face: 0,
    },
    direction: 0,
  };

  const finalPlayer = movement.reduce(
    (player: Player, instruction: Instruction) =>
      getNextPlayer(player, instruction, map),
    startingPlayer,
  );

  const currOffset = offsetForFace(finalPlayer.loc.face);
  const answer = 1000 * (finalPlayer.loc.pos[0] + 1 + currOffset[0]) +
    4 * (finalPlayer.loc.pos[1] + 1 + currOffset[1]) + finalPlayer.direction;

  return answer;
};
export default function (rawInput: string): [number, number] {
  return [part1(rawInput), part2(rawInput)];
}
