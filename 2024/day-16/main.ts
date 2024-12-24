class Position {
    constructor(public row: number, public col: number) {}
}

const directions = [new Position(-1, 0), new Position(0, 1), new Position(1, 0), new Position(0, -1)];

export function partOne(filename: string): number {
    const input = getInput(filename);
    const weights: number[][] = Array.from({ length: input.length }, () => []);
    const start = find(input, "S");
    const end = find(input, "E");

    weights[start.row][start.col] = 0;
    solveMaze(input, weights, start, 1);

    return weights[end.row][end.col];
}

function find(map: string[][], el: string): Position {
    for (let row = 0; row < map.length; row++) {
        for (let col = 0; col < map[row].length; col++) {
            if (map[row][col] === el) {
                return new Position(row, col);
            }
        }
    }

    throw new Error("Start not found");
}

function solveMaze(map: string[][], weights: number[][], pos: Position, direction: number) {
    const score = weights[pos.row][pos.col];

    for (let offset = -1; offset <= 1; offset++) {
        let dirNext = direction + offset;
        dirNext = dirNext < 0 ? directions.length + dirNext : dirNext % directions.length;

        const posNext = new Position(
            pos.row + directions[dirNext].row,
            pos.col + directions[dirNext].col,
        );

        if (map[posNext.row][posNext.col] === "#") {
            continue;
        }

        const scoreNext = score + (offset === 0 ? 1 : 1001);
        if (weights[posNext.row][posNext.col] === undefined || weights[posNext.row][posNext.col] > scoreNext) {
            weights[posNext.row][posNext.col] = scoreNext;
            solveMaze(map, weights, posNext, dirNext);
        }
    }
}

class Placement {
    constructor(public pos: Position, public dir: number) {}
}

export function partTwo(filename: string): number {
    const input = getInput(filename);
    const start = find(input, "S");

    // (row, col, direction) => score
    const scores: number[][][] = Array.from(
        { length: input.length },
        () =>
            Array.from(
                { length: input[0].length },
                () => Array.from({ length: directions.length }, () => Number.MAX_SAFE_INTEGER),
            ),
    );
    // (row, col, direction) => Placement[]
    const paths: Placement[][][][] = Array.from(
        { length: input.length },
        () =>
            Array.from(
                { length: input[0].length },
                () =>
                    Array.from(
                        { length: directions.length },
                        () => [],
                    ),
            ),
    );
    scores[start.row][start.col][1] = 0;

    const queue: [prev: Placement, curr: Placement][] = [[new Placement(start, 1), new Placement(start, 1)]];
    let scoreLowest = Number.MAX_SAFE_INTEGER;

    while (queue.length) {
        const [placePrev, placeCurr] = queue.shift()!;
        const score = scores[placeCurr.pos.row][placeCurr.pos.col][placeCurr.dir];

        for (const turn of [0, -1, 1]) {
            let dirNext = placeCurr.dir + turn;
            dirNext = dirNext < 0 ? directions.length + dirNext : dirNext % directions.length;

            const posNext = new Position(
                placeCurr.pos.row + directions[dirNext].row,
                placeCurr.pos.col + directions[dirNext].col,
            );

            if (input[posNext.row][posNext.col] === "#") {
                continue;
            }

            const scoreNext = score + (turn === 0 ? 1 : 1001);

            if (scores[posNext.row][posNext.col][dirNext] < scoreNext) {
                continue;
            }

            if (scores[posNext.row][posNext.col][dirNext] > scoreNext) {
                scores[posNext.row][posNext.col][dirNext] = scoreNext;
                paths[placeCurr.pos.row][placeCurr.pos.col][placeCurr.dir] = [];
            }

            if (input[posNext.row][posNext.col] === "E") {
                if (scoreNext > scoreLowest) {
                    break;
                }
                if (scoreNext < scoreLowest) {
                    scoreLowest = scoreNext;
                    paths[posNext.row][posNext.col][dirNext] = [];
                }
                paths[posNext.row][posNext.col][dirNext].push(placeCurr);
            }

            paths[placeCurr.pos.row][placeCurr.pos.col][placeCurr.dir].push(placePrev);
            queue.push([placeCurr, new Placement(posNext, dirNext)]);
            queue.sort(([, a], [, b]) => scores[a.pos.row][a.pos.col][a.dir] - scores[b.pos.row][b.pos.col][b.dir]);
        }
    }

    const seats: boolean[][] = Array.from({ length: input.length }, () => Array.from({ length: input[0].length }));
    const end = find(input, "E");
    const dir = scores[end.row][end.col].indexOf(scoreLowest);

    const queue2 = [new Placement(end, dir)];

    let total = 1;
    while (queue2.length) {
        const place = queue2.shift()!;

        while (paths[place.pos.row][place.pos.col][place.dir].length) {
            const path = paths[place.pos.row][place.pos.col][place.dir].pop()!;
            if (!seats[path.pos.row][path.pos.col]) {
                total++;
            }
            seats[path.pos.row][path.pos.col] = true;
            queue2.push(path);
        }
    }

    return total;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n").map((it) => it.split(""));
}
