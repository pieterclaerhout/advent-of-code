export function readFileSync(path: string): string {
  return Deno.readTextFileSync(path)
    .trimEnd()
    .replaceAll(/\r?\n/g, "\n");
}

export function readAndSplitFileSync(path: string, separator = "\n"): string[] {
  return readFileSync(path).split(separator);
}
