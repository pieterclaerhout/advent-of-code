const findUniqueCharsPosition = (chars: string[], setSize: number): number => {
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

export default function (rawInput: string): [number, number] {
  const parsedInput = rawInput.split("");

  return [
    findUniqueCharsPosition(parsedInput, 4),
    findUniqueCharsPosition(parsedInput, 14),
  ];
}
