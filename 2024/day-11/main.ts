export function partOne(filename: string): number {
    let total = 0;

    const stones = getInput(filename);

    const cache = new Map<bigint, number>();
    for (const stone of stones) {
        total += blink(cache, stone, 25);
    }

    return total;
}

export function partTwo(filename: string): number {
    let total = 0;

    const stones = getInput(filename);

    const cache = new Map<bigint, number>();
    for (const stone of stones) {
        total += blink(cache, stone, 75);
    }

    return total;
}

function blink(cache: Map<bigint, number>, stone: number, count: number) {
    if (count === 0) {
        return 1;
    }

    // Dont ask. But it's faster than using a string as a key (shaved like 25% of total runtime)
    // It's not suppose to overflow the int size of the pair,
    // but js `number` is actually a double and loses precision at high values so I cast it to `BigInt`
    // https://en.wikipedia.org/wiki/Pairing_function#Other_pairing_functions
    // http://szudzik.com/ElegantPairing.pdf
    const x = BigInt(stone);
    const y = BigInt(count);
    const key = x >= y ? x * x + x + y : x + y * y;

    const result = cache.get(key);
    if (result !== undefined) {
        return result;
    }

    let total = 0;

    if (stone === 0) {
        total += blink(cache, 1, count - 1);
        cache.set(key, total);

        return total;
    }

    const digits = countDigits(stone);
    if (digits % 2 === 0) {
        const rightDigits = Math.pow(10, digits / 2);
        const leftStone = Math.floor(stone / rightDigits);
        const rightStone = stone % (leftStone * rightDigits);

        total += blink(cache, leftStone, count - 1);
        total += blink(cache, rightStone, count - 1);

        cache.set(key, total);

        return total;
    }

    total += blink(cache, stone * 2024, count - 1);
    cache.set(key, total);

    return total;
}

function countDigits(stone: number): number {
    let digits = 0;

    while (stone >= 1) {
        stone = stone / 10;
        digits++;
    }

    return digits;
}

function getInput(filename: string): number[] {
    return Deno.readTextFileSync(filename).trimEnd().split(" ").map(Number);
}
