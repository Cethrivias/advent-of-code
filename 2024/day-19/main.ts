export function partOne(filename: string): number {
    const input = getInput(filename);

    const towels = input[0].split(", ");

    let total = 0;

    for (let i = 2; i < input.length; i++) {
        if (canBeArranged(towels, input[i])) total++;
    }

    return total;
}

function canBeArranged(towels: string[], pattern: string): boolean {
    if (!pattern.length) return true;

    for (const towel of towels) {
        if (!pattern.startsWith(towel)) continue;
        if (!canBeArranged(towels, pattern.slice(towel.length))) continue;

        return true;
    }

    return false;
}

export function partTwo(filename: string): number {
    const input = getInput(filename);

    const towels = input[0].split(", ");

    let total = 0;

    const cache = new Map<string, number>();
    for (let i = 2; i < input.length; i++) {
        total += countArrangements(towels, cache, input[i]);
    }

    return total;
}

function countArrangements(towels: string[], cache: Map<string, number>, pattern: string): number {
    if (!pattern.length) return 1;

    let total = 0;
    for (const towel of towels) {
        if (!pattern.startsWith(towel)) continue;
        const remainingPattern = pattern.slice(towel.length);
        let count = cache.get(remainingPattern);
        if (count === undefined) {
            count = countArrangements(towels, cache, pattern.slice(towel.length));
            cache.set(remainingPattern, count);
        }
        total += count;
    }

    return total;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
