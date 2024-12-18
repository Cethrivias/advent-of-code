class Position {
    constructor(public x: number, public y: number) {}
}

class Button {
    count = 0;
    constructor(public pos: Position, public cost: number) {}
}

class Machine {
    constructor(public buttons: Button[], public prize: Position) {}
}

export function partOne(filename: string): number {
    let total = 0;

    const input = getInput(filename);

    for (let i = 0; i < input.length; i += 4) {
        const commaIdx = input[i + 2].indexOf(",");
        const machine = new Machine(
            [
                new Button(
                    new Position(
                        Number(input[i].slice(12, 14)),
                        Number(input[i].slice(18, 20)),
                    ),
                    3,
                ),
                new Button(
                    new Position(
                        Number(input[i + 1].slice(12, 14)),
                        Number(input[i + 1].slice(18, 20)),
                    ),
                    1,
                ),
            ],
            new Position(
                Number(input[i + 2].slice(9, commaIdx)),
                Number(input[i + 2].slice(commaIdx + 4)),
            ),
        );

        const cache = new Map<bigint, number>();
        total += countTokensOne(cache, machine, new Position(0, 0));
    }

    return total;
}

function countTokensOne(cache: Map<bigint, number>, machine: Machine, pos: Position): number {
    if (pos.x > machine.prize.x || pos.y > machine.prize.y) {
        return 0;
    }

    if (machine.buttons[0].count > 100 || machine.buttons[1].count > 100) {
        return 0;
    }

    let solution = 0;
    if (machine.prize.x === pos.x && machine.prize.y === pos.y) {
        for (const button of machine.buttons) {
            solution += button.cost * button.count;
        }

        return solution;
    }

    const x = BigInt(machine.buttons[0].count);
    const y = BigInt(machine.buttons[1].count);
    const key = x >= y ? x * x + x + y : x + y * y;
    if (cache.has(key)) {
        return cache.get(key)!;
    }

    for (const button of machine.buttons) {
        const nextPos = new Position(pos.x + button.pos.x, pos.y + button.pos.y);
        button.count++;
        const currentSolution = countTokensOne(cache, machine, nextPos);
        button.count--;
        if (!currentSolution) {
            continue;
        }
        if (solution === 0 || solution > currentSolution) {
            solution = currentSolution;
        }
    }

    cache.set(key, solution);

    return solution;
}

export function partTwo(filename: string): number {
    let total = 0;

    const input = getInput(filename);

    for (let i = 0; i < input.length; i += 4) {
        const commaIdx = input[i + 2].indexOf(",");
        const machine = new Machine(
            [
                new Button(
                    new Position(
                        Number(input[i].slice(12, 14)),
                        Number(input[i].slice(18, 20)),
                    ),
                    3,
                ),
                new Button(
                    new Position(
                        Number(input[i + 1].slice(12, 14)),
                        Number(input[i + 1].slice(18, 20)),
                    ),
                    1,
                ),
            ],
            new Position( //          MAX_SAFE_INTEGER 9_007_199_254_740_991
                Number(input[i + 2].slice(9, commaIdx)) + 10_000_000_000_000,
                Number(input[i + 2].slice(commaIdx + 4)) + 10_000_000_000_000,
            ),
        );

        total += countTokensTwo(machine);
    }

    return total;
}

function countTokensTwo(machine: Machine): number {
    const AX = machine.buttons[0].pos.x;
    const AY = machine.buttons[0].pos.y;

    const BX = machine.buttons[1].pos.x;
    const BY = machine.buttons[1].pos.y;

    // a * AX + b * BX = machine.prize.x
    // a * AX = machine.prize.x - b * BX
    // a = (machine.prize.x - b * BX) / AX
    //
    // a * AY + b * BY = machine.prize.y
    // ((machine.prize.x - b * BX) / AX) * AY + b * BY = machine.prize.y | multiply by AX
    // (machine.prize.x - b * BX) * AY + b * BY * AX = machine.prize.y * AX | open brackets with AY
    // machine.prize.x * AY - b * BX * AY + b * BY * AX = machine.prize.y * AX | group b multiplier
    // machine.prize.x * AY - b * (BY * AX - BX * AY) = machine.prize.y * AX | move known constants to the right
    // b * (BY * AX - BX * AY) = machine.prize.y * AX - machine.prize.x * AY | divide everything by b multiplier
    // b = (machine.prize.y * AX - machine.prize.x * AY) / (BY * AX - BX * AY)

    const b = (machine.prize.y * AX - machine.prize.x * AY) / (BY * AX - BX * AY);
    if (b < 0 || b % 1 != 0) {
        return 0;
    }
    const a = (machine.prize.x - b * BX) / AX;
    if (a < 0 || a % 1 != 0) {
        return 0;
    }

    return a * 3 + b;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
