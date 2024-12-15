class Position {
    constructor(public Row: number, public Col: number) {}
}

const directions = [new Position(-1, 0), new Position(0, 1), new Position(1, 0), new Position(0, -1)];

export function partOne(filename: string): number {
    let total = 0;

    const input = getInput(filename);

    for (let row = 0; row < input.length; row++) {
        for (let col = 0; col < input[row].length; col++) {
            if (input[row][col] !== 0) {
                continue;
            }

            const visited: boolean[][] = input.map(() => []);
            total += countScore(input, new Position(row, col), visited);
        }
    }

    return total;
}

export function partTwo(filename: string): number {
    let total = 0;

    const input = getInput(filename);

    for (let row = 0; row < input.length; row++) {
        for (let col = 0; col < input[row].length; col++) {
            if (input[row][col] !== 0) {
                continue;
            }

            total += countRating(input, new Position(row, col));
        }
    }

    return total;
}

function countScore(input: number[][], pos: Position, visited: boolean[][]): number {
    // Been here
    if (visited[pos.Row][pos.Col]) {
        return 0;
    }

    // Found a trail
    if (input[pos.Row][pos.Col] === 9) {
        visited[pos.Row][pos.Col] = true;
        return 1;
    }

    let count = 0;
    // go through all accessible directions
    for (const direction of directions) {
        const nextPos = new Position(pos.Row + direction.Row, pos.Col + direction.Col);

        if (!withinBounds(input, nextPos)) {
            continue;
        }

        // Wrong elevation
        if (input[nextPos.Row][nextPos.Col] - input[pos.Row][pos.Col] !== 1) {
            continue;
        }

        count += countScore(input, nextPos, visited);
    }

    return count;
}

function countRating(input: number[][], pos: Position): number {
    // Found a trail
    if (input[pos.Row][pos.Col] === 9) {
        return 1;
    }

    let count = 0;
    // go through all accessible directions
    for (const direction of directions) {
        const nextPos = new Position(pos.Row + direction.Row, pos.Col + direction.Col);

        if (!withinBounds(input, nextPos)) {
            continue;
        }

        // Wrong elevation
        if (input[nextPos.Row][nextPos.Col] - input[pos.Row][pos.Col] !== 1) {
            continue;
        }

        count += countRating(input, nextPos);
    }

    return count;
}

function getInput(filename: string): number[][] {
    return Deno.readTextFileSync(filename).trimEnd().split("\n").map((row) => row.split("").map(Number));
}

function withinBounds(input: number[][], pos: Position): boolean {
    return pos.Row >= 0 && pos.Row < input.length && pos.Col >= 0 && pos.Col < input[pos.Row].length;
}
