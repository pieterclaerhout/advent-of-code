export function readFileSync(path: string) : string {
  return Deno.readTextFileSync(path)
    .trim()
    .replaceAll(/\r?\n/g, "\n");
}
