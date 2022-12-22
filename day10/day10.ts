import { readFileSync } from "../utils/readfile.ts";

type Operation = (["addx", number] | ["noop"]);

const parseInput = (path: string): Operation[] => {
  return readFileSync(path)
    .split('\n')
    .map((v) => {
      const [instr, val] = v.split(' ')
      return [instr, +val] as ['addx', number] | ['noop']
    });
}

// const clamp = (value: number, min: number, max: number) =>
// 	Math.min(Math.max(value, min), max);

// const simulateRope = (
//   input: StepsWithDirection[],
//   tailKnots: number
// ) : Set<string> => {
//   const head: [number, number] = [0, 0];
//   const tails: [number, number][] = [];

//   for (let t = 0; t < tailKnots; t++) {
//     tails.push([0, 0]);
//   }

//   const tailGrids = new Set(['0:0']);

//   for (const [dir, steps] of input) {
//     for (let s = 0; s < steps; s++) {
//       const index = dir === 'L' || dir === 'R' ? 0 : 1;
//       const delta = dir === 'L' || dir === 'U' ? -1 : 1;
//       head[index] += delta;

//       for (let t = 0; t < tails.length; t++) {
//         const follow = t === 0 ? head : tails[t - 1];
//         const hDist = follow[0] - tails[t][0];
//         const vDist = follow[1] - tails[t][1];

//         if (Math.abs(hDist) > 1 || Math.abs(vDist) > 1) {
//           tails[t][0] += clamp(hDist, -1, 1);
//           tails[t][1] += clamp(vDist, -1, 1);
//           if (t === tails.length - 1) {
//             tailGrids.add(tails[t].join(':'));
//           }
//         }
//       }
//     }
//   }

//   return tailGrids
// }

const part1 = (parsedInput: Operation[]) => {
  let cycle = 1;
  let x = 1;
  let signalSum = 0;

  const checkStrength = () => {
    if ((cycle - 20) % 40 === 0) {
      signalSum += cycle * x;
    }
  };

  for (const [instr, val] of parsedInput) {
    cycle++;
    if (instr === 'noop') {
      continue;
    }
    checkStrength();
    cycle++;
    x += val;

    checkStrength();
  }

  console.log("Part 1:", signalSum);
}

const part2 = (parsedInput: Operation[]) => {
  let cycle = 1;
  let x = 1;

  const screen: string[][] = [[], [], [], [], [], []];
  const draw = () => {
    const pos = cycle - 1;
    const row = Math.floor(pos / 40);
    const drawn = Math.abs((pos % 40) - x) <= 1;
    screen[row].push(drawn ? '#' : '.');
  }

  for (const [instr, val] of parsedInput) {
    draw();
    cycle += 1
    if (instr === 'noop') {
      continue;
    }
    draw();
    cycle += 1;
    x += val;
  }

  const result = screen.map((row) => row.join('')).join('\n');
  console.log("Part 2:\n" + result);
}

const run = () => {
  const inputPath = new URL('input.txt', import.meta.url).pathname;
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
}

export default run;
