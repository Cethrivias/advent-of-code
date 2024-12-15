export function partOne(filename: string): number {
    const input = getInput(filename);

    for (const line of input) {
        console.log(line);
    }

    return 0;
}

export function partTwo(_filename: string): number {
    return 0;
}

function getInput(filename: string) {
    const input = getInput(filename);

    for (const line of input) {
        console.log(line);
    }

    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
