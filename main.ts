if (import.meta.main) {
  try {
    const day = Deno.args[0];
    if (!day) {
      throw new Error("No day specified");
    }

    const dayAsString = String(day).padStart(2, "0");
    const command =
      (await import(`./day${dayAsString}/day${dayAsString}.ts`)).default;
    if (!command) {
      throw new Error(`Unknown day: ${day}`);
    }

    command();
  } catch (e) {
    console.log(e);
    Deno.exit(1);
  }
}
