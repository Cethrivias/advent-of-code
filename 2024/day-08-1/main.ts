const input = Deno.readTextFileSync("input").split("\n").filter((it) => it).map((it) => it.split(""));

class Position {
    constructor(public row: number, public col: number) {}
}

const antennas = new Map<string, Position[]>();
let total = 0;

for (let row = 0; row < input.length; row++) {
    for (let col = 0; col < input[row].length; col++) {
        if (input[row][col] === ".") {
            continue;
        }

        let pos = antennas.get(input[row][col]);
        if (!pos) {
            pos = [];
            antennas.set(input[row][col], pos);
        }
        pos.push(new Position(row, col));
    }
}

for (const positions of antennas.values()) {
    for (let i = 0; i < positions.length - 1; i++) {
        for (let j = i + 1; j < positions.length; j++) {
            const row = positions[i].row - positions[j].row;
            const col = positions[i].col - positions[j].col;

            const pos = new Position(positions[i].row + row, positions[i].col + col);
            if (withinBounds(input, pos) && input[pos.row][pos.col] !== "#") {
                input[pos.row][pos.col] = "#";
                total++;
            }

            pos.row = positions[j].row - row;
            pos.col = positions[j].col - col;
            if (withinBounds(input, pos) && input[pos.row][pos.col] !== "#") {
                input[pos.row][pos.col] = "#";
                total++;
            }
        }
    }
}

console.log(`Total: ${total}`);

function withinBounds(input: string[][], pos: Position) {
    return pos.row >= 0 && pos.row < input.length && pos.col >= 0 && pos.col < input[0].length;
}
