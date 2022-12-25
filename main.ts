#!/usr/bin/env -S deno run --allow-read

if (import.meta.main) {
  try {
    const day = Deno.args[0];
    if (!day) {
      throw new Error("No day specified");
    }

    const dayAsString = String(day).padStart(2, "0");

    const command = (await import(
      `./day${dayAsString}/day${dayAsString}.ts`
    )).default;
    if (!command) {
      throw new Error(`Unknown day: ${day}`);
    }

    const inputPath = new URL(
      `./day${dayAsString}/input.txt`,
      import.meta.url,
    ).pathname;

    const rawIput = Deno.readTextFileSync(inputPath)
      .trimEnd()
      .replaceAll(/\r?\n/g, "\n");

    command(rawIput);
  } catch (e) {
    console.log(e);
    Deno.exit(1);
  }
}
