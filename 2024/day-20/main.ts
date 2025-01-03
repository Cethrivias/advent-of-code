import { assertExists } from "@std/assert";

function arr<T>(size: number, fill: () => T) {
    return Array.from({ length: size }, fill);
}

class Position {
    constructor(public row: number, public col: number) {}
    toString() {
        return `${this.row}:${this.col}`;
    }
}

const directions = [new Position(-1, 0), new Position(0, 1), new Position(1, 0), new Position(0, -1)];

function find<T>(arr: T[][], it: T): Position | null {
    for (let row = 0; row < arr.length; row++) {
        for (let col = 0; col < arr[row].length; col++) {
            if (arr[row][col] === it) {
                return new Position(row, col);
            }
        }
    }

    return null;
}

export function partOne(filename: string, savedTime: number): number {
    const input = getInput(filename);

    const start = find(input, "S");
    assertExists(start, "No start");
    const end = find(input, "E");
    assertExists(end, "No end");

    const scores = arr(input.length, () => arr(input[0].length, () => 0));
    const cheats = new Map<string, Position[]>();

    const queue = [start];
    let score = 0;

    while (queue.length) {
        const pos = queue.pop()!;
        score++;
        scores[pos.row][pos.col] = score;

        for (const dir of directions) {
            const posN = new Position(
                pos.row + dir.row,
                pos.col + dir.col,
            );

            if (outOfBounds(input, posN)) continue;
            if (scores[posN.row][posN.col] !== 0) continue;

            if (input[posN.row][posN.col] === "#") {
                // check if we can jump across it
                const posNN = new Position(posN.row + dir.row, posN.col + dir.col);

                if (outOfBounds(input, posNN)) continue;
                if (input[posNN.row][posNN.col] !== "#" && scores[posNN.row][posNN.col] === 0) {
                    let from = cheats.get(posNN.toString());
                    if (from === undefined) {
                        from = [];
                        cheats.set(posNN.toString(), from);
                    }
                    from.push(pos);
                }

                continue;
            }

            queue.push(posN);
        }
    }

    let total = 0;

    const cheatTime = 2; // picoseconds to use the cheat
    for (const [to, from] of cheats) {
        const [toRow, toCol] = to.split(":").map(Number);
        for (const pos of from) {
            if (scores[toRow][toCol] - scores[pos.row][pos.col] >= savedTime + cheatTime) {
                total++;
            }
        }
    }

    return total;
}

export function partTwo(filename: string, savedTime: number): number {
    const input = getInput(filename);

    const start = find(input, "S");
    assertExists(start, "No start");
    const end = find(input, "E");
    assertExists(end, "No end");

    const scores = arr(input.length, () => arr(input[0].length, () => 0));
    const path: Position[] = [];

    // Map out the path from start to end
    const queue = [start];
    let score = 0;
    while (queue.length) {
        const pos = queue.pop()!;
        score++;
        scores[pos.row][pos.col] = score;
        path.push(pos);

        for (const dir of directions) {
            const posN = new Position(
                pos.row + dir.row,
                pos.col + dir.col,
            );

            if (outOfBounds(input, posN)) continue;
            if (scores[posN.row][posN.col] !== 0) continue;
            if (input[posN.row][posN.col] === "#") continue;

            queue.push(posN);
            break;
        }
    }

    // Go through the path to check which coordinates are close enough to each other
    let total = 0;
    for (let i = 0; i < path.length; i++) {
        const to = path[i];
        for (let j = 0; j < i; j++) {
            const from = path[j];

            const distance = Math.abs(from.row - to.row) + Math.abs(from.col - to.col);
            if (distance > 20) continue;

            if (scores[to.row][to.col] - scores[from.row][from.col] < savedTime + distance) continue;
            total++;
        }
    }

    return total;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n").map((row) => row.split(""));
}

function outOfBounds(input: string[][], pos: Position) {
    return pos.row < 0 || pos.row >= input.length || pos.col < 0 || pos.col >= input[pos.row].length;
}
