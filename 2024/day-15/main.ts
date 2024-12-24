class Position {
    constructor(public row: number, public col: number) {}
}

const directions = {
    ["^"]: new Position(-1, 0),
    [">"]: new Position(0, 1),
    ["v"]: new Position(1, 0),
    ["<"]: new Position(0, -1),
};

export function partOne(filename: string): number {
    const input = getInput(filename);

    const [mapRaw, movesRaw] = input.split("\n\n");
    const map = mapRaw.split("\n").map((it) => it.split(""));
    const moves = movesRaw.split("\n").join("");

    let row = 0;
    const robot = new Position(0, 0);
    for (const line of map) {
        const col = line.indexOf("@");
        if (col >= 0) {
            robot.row = row;
            robot.col = col;
            break;
        }
        row++;
    }

    for (const move of moves) {
        const direction = directions[move as "^" | ">" | "v" | "<"];

        if (moveOne(map, robot, direction)) {
            robot.row += direction.row;
            robot.col += direction.col;
        }
    }

    let total = 0;
    for (let row = 0; row < map.length; row++) {
        for (let col = 0; col < map[row].length; col++) {
            if (map[row][col] === "O") {
                total += 100 * row + col;
            }
        }
    }

    return total;
}

function moveOne(map: string[][], pos: Position, direction: Position): boolean {
    if (pos.row < 0 || pos.row >= map.length || pos.col < 0 || pos.col >= map[pos.row].length) {
        return false;
    }

    if (map[pos.row][pos.col] === "#") {
        return false;
    }

    if (map[pos.row][pos.col] === ".") {
        return true;
    }

    const nextPos = new Position(pos.row + direction.row, pos.col + direction.col);
    if (moveOne(map, nextPos, direction)) {
        map[nextPos.row][nextPos.col] = map[pos.row][pos.col];
        map[pos.row][pos.col] = ".";
        return true;
    }

    return false;
}

enum ItemType {
    Wall = "#",
    Box = "O",
    Robot = "@",
    Empty = ".",
}

class Item {
    constructor(public type: ItemType, public positions: Position[]) {}
}

export function partTwo(filename: string): number {
    const input = getInput(filename).split("\n");

    const map: Item[][] = [];
    const robot = new Item(ItemType.Robot, []);
    // Expanding the map and replacing all elements with objects representing them
    // I decided to use objects because they are compared by reference
    // This will allow me to place the same "box" in two places
    // and figure out if it's the same box by comparing the reference
    //
    // It will also be useful for the item to know where it is.
    // This way I can have a method like `move(map, item, direction)`
    // and it will know where it is and how it should be moved
    for (let row = 0; input[row].length; row++) {
        map[row] = [];
        for (const char of input[row]) {
            switch (char) {
                case "@": {
                    robot.positions.push(new Position(row, map[row].length));
                    const nothing = new Item(
                        ItemType.Empty,
                        [new Position(row, map[row].length + 1)],
                    );
                    map[row].push(robot, nothing);
                    break;
                }
                case "O": {
                    const it = new Item(
                        ItemType.Box,
                        [new Position(row, map[row].length), new Position(row, map[row].length + 1)],
                    );
                    map[row].push(it, it);
                    break;
                }
                case "#": {
                    const it = new Item(
                        ItemType.Wall,
                        [new Position(row, map[row].length), new Position(row, map[row].length + 1)],
                    );
                    map[row].push(it, it);
                    break;
                }
                default:
                    map[row].push(
                        new Item(
                            ItemType.Empty,
                            [new Position(row, map[row].length)],
                        ),
                        new Item(
                            ItemType.Empty,
                            [new Position(row, map[row].length + 1)],
                        ),
                    );
            }
        }
    }

    let moves: string = "";
    for (let row = map.length + 1; row < input.length; row++) {
        moves += input[row];
    }

    for (const move of moves) {
        const direction = directions[move as "^" | ">" | "v" | "<"];

        // in Part 1 it was possible to move boxes as we were returning from the recursive chain
        // in Part 2 the chain of vertical movements can diverge
        // And it might be possible to move one branch but not the other
        // So we might end up in a situation where one branch was moved
        // and the other branch failed
        // To avoid it I decided to split the check and the movement into two separate actions
        if (canMoveV2(map, robot, direction)) {
            moveV2(map, robot, direction);
        }
    }

    let total = 0;
    for (let row = 0; row < map.length; row++) {
        for (let col = 0; col < map[row].length; col++) {
            if (map[row][col].type === ItemType.Box) {
                total += 100 * row + col;
                col++;
            }
        }
    }

    return total;
}

function canMoveV2(map: Item[][], item: Item, direction: Position): boolean {
    const nextItems: Item[] = [];

    for (let i = 0; i < item.positions.length; i++) {
        const pos = item.positions[i];
        const nextPos = new Position(pos.row + direction.row, pos.col + direction.col);
        const nextItem = map[nextPos.row][nextPos.col];

        if (nextItem.type === ItemType.Wall) {
            return false;
        }

        if (nextItem.type === ItemType.Empty) {
            continue;
        }

        if (nextItem !== item && !nextItems.includes(nextItem)) {
            nextItems.push(nextItem);
        }
    }

    for (const nextItem of nextItems) {
        if (!canMoveV2(map, nextItem, direction)) {
            return false;
        }
    }
    return true;
}

function moveV2(map: Item[][], item: Item, direction: Position) {
    // assertNotEquals(item.type, ItemType.Wall);
    // assertNotEquals(item.type, ItemType.Empty);

    const nextItems: Item[] = [];

    // Base case. Moving the robot itself
    if (item.type === ItemType.Robot) {
        const pos = item.positions[0];
        const nextPos = new Position(pos.row + direction.row, pos.col + direction.col);

        // assertNotEquals(map[nextPos.row][nextPos.col].type, ItemType.Wall);
        // assertNotEquals(map[nextPos.row][nextPos.col].type, ItemType.Robot);
        if (map[nextPos.row][nextPos.col].type === ItemType.Box) {
            moveV2(map, map[nextPos.row][nextPos.col], direction);
        }

        map[nextPos.row][nextPos.col].positions[0].row -= direction.row;
        map[nextPos.row][nextPos.col].positions[0].col -= direction.col;
        map[pos.row][pos.col] = map[nextPos.row][nextPos.col];
        map[nextPos.row][nextPos.col] = item;
        pos.row += direction.row;
        pos.col += direction.col;

        return;
    }

    // Moving boxes now

    // vertical
    if (direction.row != 0) {
        for (let i = 0; i < item.positions.length; i++) {
            const pos = item.positions[i];
            const nextPos = new Position(pos.row + direction.row, pos.col);
            const nextItem = map[nextPos.row][nextPos.col];

            if (
                nextItem.type !== ItemType.Empty && // something blocks our path
                !nextItems.includes(nextItem) // we haven't moved it yet (TODO: might not need it)
            ) {
                nextItems.push(nextItem);
                // First move the blocking item
                moveV2(map, nextItem, direction);
            }

            // Everything was moved out of the way. Next position is now empty
            map[nextPos.row][nextPos.col].positions[0].row -= direction.row;
            map[pos.row][pos.col] = map[nextPos.row][nextPos.col];
            map[nextPos.row][nextPos.col] = item;
            pos.row += direction.row;
        }

        return;
    }

    // horizontal
    //
    // Instead of pushing the box I'm going to flip it
    // Pushing left-to-right '->[]'
    // After this is done it's going to be '->]['

    // Positions are always left-to-right.
    // So if we a moving right we need to pick the left side of the box and flip it with whatever is after the box
    // and the opposite if we are moving left
    const pos = direction.col > 0 ? item.positions[0] : item.positions[1];

    const nextPos = new Position(pos.row, pos.col + direction.col * 2);
    const nextItem = map[nextPos.row][nextPos.col];

    if (nextItem.type !== ItemType.Empty) {
        moveV2(map, nextItem, direction);
    }

    // Everything was moved out of the way. Next position is now empty
    map[nextPos.row][nextPos.col].positions[0].col -= direction.col * 2;
    map[pos.row][pos.col] = map[nextPos.row][nextPos.col];
    map[nextPos.row][nextPos.col] = item;
    pos.col += direction.col * 2;

    // We need to maintain position in a left-to-right order
    if (item.positions[0].col > item.positions[1].col) {
        const pos = item.positions[0];
        item.positions[0] = item.positions[1];
        item.positions[1] = pos;
    }
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd();
}
