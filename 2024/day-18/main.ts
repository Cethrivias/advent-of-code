function arr<T>(size: number, fill: () => T) {
    return Array.from({ length: size }, fill);
}

class Position {
    constructor(public row: number, public col: number) {}

    toString() {
        return `${this.col},${this.row}`;
    }
}

const directions = [new Position(-1, 0), new Position(0, 1), new Position(1, 0), new Position(0, -1)];

export function partOne(filename: string, size: number, bytes: number): number {
    const input = getInput(filename);

    const map = arr(size, () => arr(size, () => "."));
    const scores = arr(size, () => arr(size, () => Number.MAX_SAFE_INTEGER));

    for (let byte = 0; byte < bytes; byte++) {
        const [col, row] = input[byte].split(",").map(Number);
        map[row][col] = "#";
    }

    const queue = [new Position(0, 0)];
    scores[0][0] = 0;

    while (queue.length) {
        const pos = queue.pop()!;
        const scoreCurr = scores[pos.row][pos.col];

        for (const dir of directions) {
            const posNext = new Position(pos.row + dir.row, pos.col + dir.col);
            const scoreNext = scoreCurr + 1;

            if (posNext.row < 0 || posNext.row >= size || posNext.col < 0 || posNext.col >= size) continue;
            if (map[posNext.row][posNext.col] === "#") continue;
            if (scoreNext >= scores[posNext.row][posNext.col]) continue;

            scores[posNext.row][posNext.col] = scoreNext;
            queue.push(posNext);
        }
    }

    return scores[size - 1][size - 1];
}

export function partTwo(filename: string, size: number, bytes: number): string {
    const input = getInput(filename);

    const map = arr(size, () => arr(size, () => "."));

    for (let byte = 0; byte < bytes; byte++) {
        const [col, row] = input[byte].split(",").map(Number);
        map[row][col] = "#";
    }

    let answer = "0,0";
    for (let byte = bytes; byte < input.length; byte++) {
        const [col, row] = input[byte].split(",").map(Number);
        map[row][col] = "#";

        if (findScore(map) === Number.MAX_SAFE_INTEGER) {
            answer = `${col},${row}`;
            break;
        }
    }

    return answer;
}

function findScore(map: string[][]): number {
    const scores = arr(map.length, () => arr(map[0].length, () => Number.MAX_SAFE_INTEGER));
    const queue = [new Position(0, 0)];
    scores[0][0] = 0;

    while (queue.length) {
        const pos = queue.pop()!;
        const scoreCurr = scores[pos.row][pos.col];

        for (const dir of directions) {
            const posNext = new Position(pos.row + dir.row, pos.col + dir.col);
            const scoreNext = scoreCurr + 1;

            if (posNext.row < 0 || posNext.row >= map.length || posNext.col < 0 || posNext.col >= map[0].length) {
                continue;
            }
            if (map[posNext.row][posNext.col] === "#") continue;
            if (scoreNext >= scores[posNext.row][posNext.col]) continue;

            scores[posNext.row][posNext.col] = scoreNext;
            queue.push(posNext);
        }
    }

    return scores[scores.length - 1][scores[0].length - 1];
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
