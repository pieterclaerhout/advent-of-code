export function readFileSync(path: string): string {
  return Deno.readTextFileSync(path)
    .trimEnd()
    .replaceAll(/\r?\n/g, "\n");
}
