import { assertEquals, assertExists, assertInstanceOf, assertStringIncludes } from "@std/assert";

function arr<T>(length: number, fill: () => T): T[] {
    return Array.from({ length }, fill);
}

function last<T>(arr: T[]): T {
    return arr[arr.length - 1];
}

export function partOne(filename: string): bigint {
    const input = getInput(filename);

    let total = 0n;
    for (const line of input) {
        const secret = BigInt(line);
        total += generateSecret(secret);
    }

    return total;
}

function generateSecret(secret: bigint): bigint {
    for (let i = 0; i < 2000; i++) {
        secret = nextSecret(secret);
    }

    return secret;
}

function nextSecret(secret: bigint): bigint {
    // multiplying the secret number by 64 (x << 6) then XOR with the secret number
    // prune (x % 16777216) (0b1000000000000000000000000) (1 << 24)
    // dividing the secret number by 32 (x >> 5) then XOR with the secret number
    // prune (x % 16777216) (0b1000000000000000000000000) (1 << 24)
    // multiplying the secret number by 2048 (x << 11) then mix into the secret number
    // prune
    // bop it, push it, twist it, pull it, shake it
    secret ^= secret << 6n;
    secret %= 16777216n;
    secret ^= secret >> 5n;
    secret %= 16777216n;
    secret ^= secret << 11n;
    secret %= 16777216n;

    return secret;
}

Deno.test("generateSecret", async function (t) {
    const examples: [seed: bigint, expected: bigint][] = [
        [1n, 8685429n],
        [10n, 4700978n],
        [100n, 15273692n],
        [2024n, 8667524n],
    ];
    for (const [seed, expected] of examples) {
        await t.step(seed.toString(), function () {
            assertEquals(generateSecret(seed), expected);
        });
    }
});

export function partTwo(filename: string): bigint {
    const input = getInput(filename);
    const steps = 2000;

    const cache = new Map<string, bigint>();

    for (let buyerId = 0; buyerId < input.length; buyerId++) {
        let secret = BigInt(input[buyerId]);
        let price = secret % 10n;
        const sequence: bigint[] = [];
        const sequences = new Set<string>();

        for (let i = 1; i < steps + 1; i++) {
            const secretN = nextSecret(secret);
            const priceN = secretN % 10n;
            const change = priceN - price;
            secret = secretN;
            price = priceN;

            sequence.push(change);
            if (sequence.length > 4) sequence.shift();

            const key = sequence.join(",");
            if (sequences.has(key)) continue;
            sequences.add(key);

            const bananas = cache.get(key) || 0n;
            cache.set(key, bananas + priceN);
        }
    }

    let bestBananas = 0n;

    for (const bananas of cache.values()) {
        if (bananas > bestBananas) {
            bestBananas = bananas;
        }
    }

    return bestBananas;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
