import { readFileSync } from "../utils/readfile.ts";

const parseInput = (path: string): string[] => {
  return readFileSync(path).split('');
}

const findUniqueCharsPosition = (chars: string[], setSize: number) => {
	const lastChars: string[] = [];
	for (let i = 0; i < chars.length; i++) {
		lastChars.push(chars[i]);
		if (lastChars.length > setSize) {
      lastChars.shift();
    }
		if (new Set(lastChars).size === setSize) {
			return i + 1
		}
	}
	throw 'failed'
}

const part1 = (parsedInput: string[]) => {
  const result = findUniqueCharsPosition(parsedInput, 4);
  console.log("Part 1:",  result)
}

const part2 = (parsedInput: string[]) => {
  const result = findUniqueCharsPosition(parsedInput, 14);
  console.log("Part 2:",  result);
}

const run = () => {
  const inputPath = new URL('input.txt', import.meta.url).pathname;
  const parsedInput = parseInput(inputPath);

  part1(parsedInput);
  part2(parsedInput);
}

export default run;
