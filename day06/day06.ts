import { readFileSync } from "../utils/readfile.ts";

const findUniqueCharsPosition = (chars: string[], setSize: number) => {
  const lastChars: string[] = [];
  for (let i = 0; i < chars.length; i++) {
    lastChars.push(chars[i]);
    if (lastChars.length > setSize) {
      lastChars.shift();
    }
    if (new Set(lastChars).size === setSize) {
      return i + 1;
    }
  }
  throw "failed";
};

export default function (inputPath: string) {
  const parsedInput = readFileSync(inputPath).split("");

  console.log("Part 1:", findUniqueCharsPosition(parsedInput, 4));
  console.log("Part 2:", findUniqueCharsPosition(parsedInput, 14));
}
