class Position {
    constructor(public row: number, public col: number) {}
}

const numericKeypad: Record<string, Position> = {
    // 7, 8, 9
    "7": new Position(0, 0),
    "8": new Position(0, 1),
    "9": new Position(0, 2),
    // 4, 5, 6
    "4": new Position(1, 0),
    "5": new Position(1, 1),
    "6": new Position(1, 2),
    // 1, 2, 3
    "1": new Position(2, 0),
    "2": new Position(2, 1),
    "3": new Position(2, 2),
    // _, 0, A
    "_": new Position(3, 0),
    "0": new Position(3, 1),
    "A": new Position(3, 2),
};

const directionalKeypad: Record<string, Position> = {
    "_": new Position(0, 0),
    "^": new Position(0, 1),
    "A": new Position(0, 2),

    "<": new Position(1, 0),
    "v": new Position(1, 1),
    ">": new Position(1, 2),
};

export function partOne(filename: string): number {
    const input = getInput(filename);

    let total = 0;
    for (const line of input) {
        let keys: string[] = line;
        for (const keypad of [numericKeypad, directionalKeypad, directionalKeypad]) {
            let key: string = "A";
            const keysN: string[] = [];

            for (const keyN of keys) {
                const path = getPath(keypad, key, keyN);
                keysN.push(...path);
                key = keyN;
            }

            keys = keysN;
        }

        total += keys.length * Number(line.slice(0, -1).join(""));
    }

    return total;
}

function getPath(keypad: Record<string, Position>, key: string, keyN: string): string[] {
    const path: string[] = [];

    const emptyPos = keypad["_"];

    const keyPos = keypad[key];
    const keyPosN = keypad[keyN];

    const rows = keyPosN.row - keyPos.row;
    const cols = keyPosN.col - keyPos.col;

    if (keyPos.col + cols === emptyPos.col && keyPos.row === emptyPos.row) {
        // We are on a row with an empty button and will hit it if we move through cols first
        for (let i = 0; i < Math.abs(rows); i++) {
            path.push(rows < 0 ? "^" : "v");
        }
        for (let i = 0; i < Math.abs(cols); i++) {
            path.push(cols < 0 ? "<" : ">");
        }
    } else if (keyPos.row + rows === emptyPos.row && keyPos.col === emptyPos.col) {
        // We are on a col with an empty button and will hit it if we move through rows first
        for (let i = 0; i < Math.abs(cols); i++) {
            path.push(cols < 0 ? "<" : ">");
        }
        for (let i = 0; i < Math.abs(rows); i++) {
            path.push(rows < 0 ? "^" : "v");
        }
    } else {
        // I think this entire `else` block can be simplified but I can't be bothered
        // We are safe to do whatever we want
        if (cols > 0) {
            // Best to end on '>' because it's close to A
            for (let i = 0; i < Math.abs(rows); i++) {
                path.push(rows < 0 ? "^" : "v");
            }
            for (let i = 0; i < Math.abs(cols); i++) {
                path.push(cols < 0 ? "<" : ">");
            }
        } else if (rows < 0) {
            // Ending on '^' is also good
            for (let i = 0; i < Math.abs(cols); i++) {
                path.push(cols < 0 ? "<" : ">");
            }
            for (let i = 0; i < Math.abs(rows); i++) {
                path.push(rows < 0 ? "^" : "v");
            }
        } else {
            // cols < 0. We need to make sure that '<' (the farthest button from A) is pressed first
            for (let i = 0; i < Math.abs(cols); i++) {
                path.push(cols < 0 ? "<" : ">");
            }
            for (let i = 0; i < Math.abs(rows); i++) {
                path.push(rows < 0 ? "^" : "v");
            }
        }
    }

    path.push("A");

    return path;
}

export function partTwo(filename: string, keypads: number): bigint {
    const input = getInput(filename);

    let total = 0n;
    const cache = new Map<string, bigint>();
    for (const line of input) {
        // console.log("line:", line);

        const numericKeypadMoves = [];

        let key: string = "A";
        for (const keyN of line) {
            const path = getPath(numericKeypad, key, keyN);
            numericKeypadMoves.push(...path);
            key = keyN;
        }

        let keysCount = 0n;
        let numericKeypadMove = "A";
        for (const numericKeypadMoveN of numericKeypadMoves) {
            // console.log(`Moving: ${numericKeypadMoveN}`);
            keysCount += countMoves(keypads, cache, numericKeypadMove, numericKeypadMoveN);
            numericKeypadMove = numericKeypadMoveN;
        }

        total += keysCount * BigInt(line.slice(0, -1).join(""));

        // return total;
    }

    return total;
}

function countMoves(keypadCount: number, cache: Map<string, bigint>, from: string, to: string): bigint {
    if (keypadCount === 0) return 1n;

    const cacheKey = `${from}:${to}:${keypadCount}`;
    let total = cache.get(cacheKey);
    if (total !== undefined && total > 0n) {
        return total;
    }
    total = 0n;

    const path = getPath(directionalKeypad, from, to);
    let key = "A";
    for (const keyN of path) {
        total += countMoves(keypadCount - 1, cache, key, keyN);
        key = keyN;
    }

    cache.set(cacheKey, total);

    return total;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n").map((row) => row.split(""));
}
