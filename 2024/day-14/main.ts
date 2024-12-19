export function partOne(filename: string, rows: number, cols: number): number {
    const input = getInput(filename);

    const map: number[][] = Array.from({ length: rows }, () => Array(cols).fill(0));

    for (const line of input) {
        const [posRaw, velRaw] = line.split(" ");
        let [colRaw, rowRaw] = posRaw.slice(2).split(",");
        let posRow = Number(rowRaw);
        let posCol = Number(colRaw);

        [colRaw, rowRaw] = velRaw.slice(2).split(",");
        const velRow = Number(rowRaw);
        const velCol = Number(colRaw);

        posRow += velRow * 100;
        posCol += velCol * 100;
        posRow %= rows;
        if (posRow < 0) {
            posRow += rows;
        }
        posCol %= cols;
        if (posCol < 0) {
            posCol += cols;
        }
        map[posRow][posCol]++;
    }

    const midRow = Math.floor(rows / 2);
    const midCol = Math.floor(cols / 2);

    const quadrants: [[number, number], [number, number]][] = [
        [[0, midRow], [0, midCol]],
        [[0, midRow], [midCol + 1, cols]],

        [[midRow + 1, rows], [0, midCol]],
        [[midRow + 1, rows], [midCol + 1, cols]],
    ];

    let total = 1;
    for (const quadrant of quadrants) {
        let robots = 0;
        for (let row = quadrant[0][0]; row < quadrant[0][1]; row++) {
            for (let col = quadrant[1][0]; col < quadrant[1][1]; col++) {
                robots += map[row][col];
            }
        }
        total *= robots;
    }
    return total;
}

class Position {
    constructor(public row: number, public col: number) {}
}

class Robot {
    constructor(public pos: Position, public vel: Position) {}
}

export function partTwo(filename: string, rows: number, cols: number): number {
    const input = getInput(filename);

    const robots: Robot[] = [];

    for (const line of input) {
        const [posRaw, velRaw] = line.split(" ");
        let [colRaw, rowRaw] = posRaw.slice(2).split(",");
        const posRow = Number(rowRaw);
        const posCol = Number(colRaw);

        [colRaw, rowRaw] = velRaw.slice(2).split(",");
        const velRow = Number(rowRaw);
        const velCol = Number(colRaw);

        robots.push(
            new Robot(
                new Position(posRow, posCol),
                new Position(velRow, velCol),
            ),
        );
    }

    let seconds = 1;
    aaaaaaaaaaaa: while (true) {
        const map: number[][] = Array.from({ length: rows }, () => Array(cols).fill(0));

        for (const robot of robots) {
            robot.pos.row += robot.vel.row;
            robot.pos.col += robot.vel.col;

            robot.pos.row %= rows;
            if (robot.pos.row < 0) {
                robot.pos.row += rows;
            }
            robot.pos.col %= cols;
            if (robot.pos.col < 0) {
                robot.pos.col += cols;
            }
            map[robot.pos.row][robot.pos.col] = 1;
        }

        let frame = 0;
        for (let row = 0; row < rows; row++) {
            for (let col = 0; col < cols; col++) {
                if (map[row][col] === 1) {
                    frame++;
                } else {
                    frame = 0;
                }
                if (frame > 30) {
                    // console.log(map.map((it) => it.join("")).join("\n"));
                    // console.log("Seconds:", seconds);
                    break aaaaaaaaaaaa;
                }
            }
        }

        seconds++;
    }

    return seconds;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
