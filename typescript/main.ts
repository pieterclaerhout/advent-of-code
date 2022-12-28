#!/usr/bin/env -S deno run --allow-read

const runDay = async (day: string): Promise<void> => {
  const dayAsString = day.padStart(2, "0");

  const command = (await import(
    `./day${dayAsString}/day${dayAsString}.ts`
  )).default;
  if (!command) {
    throw new Error(`Unknown day: ${day}`);
  }

  const inputPath = new URL(
    `../inputs/day${dayAsString}/input.txt`,
    import.meta.url,
  ).pathname;

  const rawIput = Deno.readTextFileSync(inputPath)
    .trimEnd()
    .replaceAll(/\r?\n/g, "\n");

  const [result1, result2] = command(rawIput);
  console.log("Part 1: " + result1);
  console.log("Part 2: " + result2);
};

if (import.meta.main) {
  try {
    const day = Deno.args[0];
    if (!day) {
      for (let i = 1; i <= 25; i++) {
        console.log(`>>> Day ${i} <<<`);
        await runDay(String(i));
        console.log();
      }
    } else {
      await runDay(day);
    }
  } catch (e) {
    console.log(e);
    Deno.exit(1);
  }
}
