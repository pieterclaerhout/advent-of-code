export function readFileSync(path: string) : string {
  return Deno.readTextFileSync(path)
    .replaceAll(/\r?\n/g, "\n");
}
