import { assertEquals } from "@std/assert";

export function partOne(filename: string): number {
    const input = getInput(filename);

    const locks: number[][] = [];
    const keys: number[][] = [];

    for (const block of input) {
        console.log(block);
        const schematic = block.split("\n");

        if (schematic[0][0] === "#") {
            locks.push(parseLock(schematic));
        } else {
            keys.push(parseKey(schematic));
        }
    }

    let total = 0;
    for (const lock of locks) {
        for (const key of keys) {
            if (fit(lock, key)) total++;
        }
    }

    return total;
}

function parseLock(schematic: string[]): number[] {
    const lock = Array.from({ length: schematic[0].length }, () => 0);

    for (let col = 0; col < schematic[0].length; col++) {
        for (let row = 0; row < schematic.length; row++) {
            if (schematic[row][col] === "#") continue;
            lock[col] = row - 1;
            break;
        }
    }

    return lock;
}

Deno.test("Parse Lock", function () {
    const schematic = [
        "#####",
        ".####",
        ".####",
        ".####",
        ".#.#.",
        ".#...",
        ".....",
    ];

    const lock = parseLock(schematic);
    assertEquals(lock.length, 5, "Wrong lock length");
    assertEquals(lock, [0, 5, 3, 4, 3], "Wrong lock");
});

function parseKey(schematic: string[]): number[] {
    const key = Array.from({ length: schematic[0].length }, () => 0);

    for (let col = 0; col < schematic[0].length; col++) {
        for (let row = schematic.length - 1; row >= 0; row--) {
            if (schematic[row][col] === "#") continue;
            key[col] = schematic[0].length - row;
            break;
        }
    }

    return key;
}

Deno.test("Parse Key", function () {
    const schematic = [
        ".....",
        "#....",
        "#....",
        "#...#",
        "#.#.#",
        "#.###",
        "#####",
    ];

    const key = parseKey(schematic);
    assertEquals(key.length, 5, "Wrong key length");
    assertEquals(key, [5, 0, 2, 1, 3], "Wrong key");
});

function fit(lock: number[], key: number[]): boolean {
    for (let i = 0; i < lock.length; i++) {
        if (lock[i] + key[i] < 6) continue;
        // console.log(`'${lock}' '${key}', i: ${i}`);
        return false;
    }
    return true;
}

Deno.test("Fit", async function (t) {
    await t.step("ex 1", function () {
        const lock = [0, 5, 3, 4, 3];
        const key = [5, 0, 2, 1, 3];
        assertEquals(fit(lock, key), false, `${lock} shouldn't fit ${key}`);
    });
    await t.step("ex 2", function () {
        const lock = [0, 5, 3, 4, 3];
        const key = [4, 3, 4, 0, 2];
        assertEquals(fit(lock, key), false, `${lock} shouldn't fit ${key}`);
    });
    await t.step("ex 3", function () {
        const lock = [0, 5, 3, 4, 3];
        const key = [3, 0, 2, 0, 1];
        assertEquals(fit(lock, key), true, `${lock} should fit ${key}`);
    });
});

export function partTwo(filename: string): number {
    const input = getInput(filename);

    for (const line of input) {
        console.log(line);
    }

    return 0;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n\n");
}
