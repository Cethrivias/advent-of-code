import { assertExists } from "@std/assert/exists";
import { assertArrayIncludes, assertEquals } from "@std/assert";

type Gate = [string | null, Op | null, string | null, bigint | null];

type Device = Map<string, Gate>;

enum Op {
    And = "AND",
    Or = "OR",
    Xor = "XOR",
}

export function partOne(filename: string): bigint {
    const input = getInput(filename);

    const device = new Map<string, Gate>();

    let gates = false;
    for (const line of input) {
        if (line === "") {
            gates = true;
        }

        if (!gates) {
            const wire = line.slice(0, 3);
            const val = line.slice(5);
            device.set(wire, [null, null, null, BigInt(val)]);
        } else {
            const [operationRaw, gate] = line.split(" -> ");
            const operation = operationRaw.split(" ");
            device.set(gate, [operation[0], operation[1] as Op, operation[2], null]);
        }
    }

    return solveDevice(device);
}

function solveDevice(device: Map<string, Gate>): bigint {
    let i = 0n;
    let result = 0n;
    while (true) {
        const wire = "z" + i.toString().padStart(2, "0");

        const gate = device.get(wire);
        if (!gate) break;
        const shortCircuit = new Set<string>();
        shortCircuit.add(wire);
        const value = resolveGate(device, gate);
        if (value === -1n) break;
        result |= value << i;
        i++;
    }

    return result;
}

function resolveGate(device: Map<string, Gate>, gate: Gate): bigint {
    if (gate[3] !== null) return gate[3];

    // assertExists(gate[0]);
    // assertExists(gate[1]);
    // assertExists(gate[2]);
    const a = device.get(gate[0]!)!;
    const b = device.get(gate[2]!)!;
    // assertExists(a);
    // assertExists(b);

    let val = 0n;

    if (gate[1] === "AND") {
        val = resolveGate(device, a) & resolveGate(device, b);
    } else if (gate[1] === "XOR") {
        val = resolveGate(device, a) ^ resolveGate(device, b);
    } else if (gate[1] === "OR") {
        val = resolveGate(device, a) | resolveGate(device, b);
    } else {
        throw new Error("Unknown operand");
    }

    return val;
}

export function partTwo(filename: string): string {
    const input = getInput(filename);

    const [device, _x, _y, z] = parseInput(input);

    const brokenResult = solveDevice(device);
    const diff = z ^ brokenResult;

    console.log(z.toString(2));
    console.log(brokenResult.toString(2));
    console.log(diff.toString(2).padStart(brokenResult.toString(2).length, " "));

    let num = 0n;
    while (true) {
        if (!verifyZ(device, makeName("z", num), num)) break;
        num++;
    }
    console.log("failed on", makeName("z", num));
    // rkf <-> z09
    // jgb <-> z20
    // vcg <-> z24
    // rvc <-> rrs

    return ["rkf", "z09", "jgb", "z20", "vcg", "z24", "rvc", "rrs"].sort().join(",");
}

function verifyZ(device: Map<string, Gate>, wire: string, pos: bigint) {
    console.log("verify Z", wire, pos);
    const [x, op, y] = device.get(wire)!;
    if (op !== "XOR") return false;
    assertExists(x);
    assertExists(y);

    if (pos === 0n) {
        const wires = [x, y].sort();
        return wires[0] === "x00" && wires[1] === "y00";
    }
    return (verifyDirectXor(device, x, pos) && verifyCarry(device, y, pos)) ||
        (verifyDirectXor(device, y, pos) && verifyCarry(device, x, pos));
}

function verifyDirectXor(device: Map<string, Gate>, wire: string, pos: bigint) {
    console.log("verify XOR", wire, pos);
    const [x, op, y] = device.get(wire)!;
    if (op !== "XOR") return false;
    return arrEq([x, y].sort(), [makeName("x", pos), makeName("y", pos)]);
}

function verifyCarry(device: Map<string, Gate>, wire: string, pos: bigint) {
    console.log("verify carry", wire, pos);
    const [x, op, y] = device.get(wire)!;
    assertExists(x);
    assertExists(y);

    if (pos === 1n) {
        return (op === Op.And && arrEq([x, y].sort(), ["x00", "y00"]));
    }
    if (op !== Op.Or) return false;
    return verifyDirectCarry(device, x, pos - 1n) && verifyRecarry(device, y, pos - 1n) ||
        verifyDirectCarry(device, y, pos - 1n) && verifyRecarry(device, x, pos - 1n);
}

function verifyDirectCarry(device: Device, wire: string, pos: bigint) {
    console.log("verify direct", wire, pos);
    const [x, op, y] = device.get(wire)!;
    if (op != Op.And) return false;
    assertExists(x);
    assertExists(y);
    return arrEq([x, y].sort(), [makeName("x", pos), makeName("y", pos)]);
}

function verifyRecarry(device: Device, wire: string, pos: bigint): boolean {
    console.log("verify recarry", wire, pos);
    const [x, op, y] = device.get(wire)!;
    if (op !== Op.And) return false;
    assertExists(x);
    assertExists(y);

    return verifyDirectXor(device, x, pos) && verifyCarry(device, y, pos) ||
        verifyDirectXor(device, y, pos) && verifyCarry(device, x, pos);
}

function makeName(wire: string, pos: bigint): string {
    return `${wire}${pos.toString().padStart(2, "0")}`;
}

function arrEq(x: unknown[], y: unknown[]): boolean {
    if (x.length !== y.length) return false;

    for (let i = 0; i < x.length; i++) {
        if (x[i] === y[i]) continue;
        return false;
    }

    return true;
}

function checkBrokenWires(device: Map<string, Gate>, brokenWires: bigint[], expected: bigint): boolean {
    try {
        for (const bWire of brokenWires) {
            const wire = `z${bWire.toString().padStart(2, "0")}`;
            const gate = device.get(wire)!;
            const res = resolveGate(device, gate);

            if (res !== ((expected >> bWire) & 1n)) {
                return false;
            }
        }
    } catch (_error) {
        return false;
    }

    return true;
}

function canBeSwapped(wire: string): boolean {
    return wire[0] !== "x" && wire[0] !== "y";
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}

function parseInput(input: string[]): [device: Map<string, Gate>, x: bigint, y: bigint, z: bigint] {
    const device = new Map<string, Gate>();

    let parsingGates = false;
    let x = 0n;
    let y = 0n;
    for (const line of input) {
        if (line === "") {
            parsingGates = true;
            continue;
        }

        if (!parsingGates) {
            const wire = line.slice(0, 3);
            const val = BigInt(line.slice(5));
            if (wire[0] === "x") {
                const pos = BigInt(wire.slice(1, 3));
                x |= val << pos;
            } else if (wire[0] === "y") {
                const pos = BigInt(wire.slice(1, 3));
                y |= val << pos;
            }
            device.set(wire, [null, null, null, val]);
        } else {
            const [operationRaw, gate] = line.split(" -> ");
            const operation = operationRaw.split(" ");
            device.set(gate, [operation[0], operation[1] as Op, operation[2], null]);
        }
    }

    return [device, x, y, x + y];
}

Deno.test(parseInput.name, async function (t) {
    await t.step("Example 1", function () {
        const input = [
            "x00: 1",
            "x01: 1",
            "x02: 0",
            "x03: 1",
            "y00: 1",
            "y01: 0",
            "y02: 1",
            "y03: 1",
        ];

        const [_, x, y, z] = parseInput(input);

        assertEquals(x, 11n, "Wrong X");
        assertEquals(y, 13n, "Wrong Y");
        assertEquals(z, 24n, "Wrong Z");
    });
});

function getBrokenPositions(brokenResult: bigint, z: bigint): bigint[] {
    const diff = z ^ brokenResult;

    console.log(z.toString(2));
    console.log(brokenResult.toString(2));
    console.log(diff.toString(2).padStart(brokenResult.toString(2).length, " "));

    const brokenWires: bigint[] = [];
    for (let i = 0n; i < 64n; i++) {
        if (!(diff & (1n << i))) continue;
        if (i === 9n) continue; // specific to my input

        brokenWires.push(i);

        // console.log(`Pos: ${i}`);
        // const gates: Gate[] = [device.get(`z${i.toString().padStart(2, "0")}`)!];
        // while (gates.length) {
        //     const gate = gates.shift()!;
        //     console.log(gate);
        //     if (gate[0] !== null) gates.push(device.get(gate[0])!);
        //     if (gate[2] !== null) gates.push(device.get(gate[2])!);
        // }
    }

    return brokenWires;
}

Deno.test(getBrokenPositions.name, async function (t) {
    await t.step("Example 1", function () {
        const z = 25n;
        const brokenResult = 24n;

        const pos = getBrokenPositions(brokenResult, z);
        assertEquals(pos.length, 1, "Too many broken positions");
        assertEquals(pos[0], 0n, "Wrong broken position");
    });
    await t.step("Example 2", function () {
        const z = 45920934306510n;
        const brokenResult = 45923082839246n;

        const pos = getBrokenPositions(brokenResult, z);
        // assertEquals(pos.length, 11, "Wrong count");
        // assertArrayIncludes(pos, [9n, 10n, 20n, 21n, 22n, 23n, 24n, 31n, 32n, 33n, 34n], "Wrong positions");
        assertEquals(pos.length, 10, "Wrong count");
        assertArrayIncludes(pos, [10n, 20n, 21n, 22n, 23n, 24n, 31n, 32n, 33n, 34n], "Wrong positions");
    });
});

function getCombinations<T>(elements: T[], sample: number, reqEl: T | null = null): T[][] {
    const combos: T[][] = [];

    const getCombosInner = (elements: T[], sample: number, start: number, tmp: T[]) => {
        if (tmp.length === sample) {
            if (reqEl !== null && !tmp.includes(reqEl)) return;
            combos.push(tmp.slice());
            return;
        }

        for (let i = start; i < elements.length; i++) {
            tmp.push(elements[i]);
            getCombosInner(elements, sample, i + 1, tmp);
            tmp.pop();
        }
    };

    // for (let i = 1; i <= elements.length; i++) {
    //     getCombosInner(elements, i, 0, []);
    // }
    getCombosInner(elements, sample, 0, []);

    return combos;
}

Deno.test(getCombinations.name, async function (t) {
    await t.step("Example 1. No req el", function () {
        const a = Array.from({ length: 5 }, (_, i) => i);
        const res = getCombinations(a, 3);
        assertEquals(res.length, 10);
    });
    await t.step("Example 2. No req el", function () {
        const a = Array.from({ length: 4 }, (_, i) => i);
        const res = getCombinations(a, 3);
        assertEquals(res.length, 4);
    });
    await t.step("Example 3. No req el", function () {
        const a = Array.from({ length: 8 }, (_, i) => i);
        const res = getCombinations(a, 2);
        assertEquals(res.length, 28);
    });
    await t.step("Example 4. No req el", function () {
        const a = Array.from({ length: 28 }, (_, i) => i);
        const res = getCombinations(a, 4);
        assertEquals(res.length, 20475);
    });
    await t.step("Example 5. No req el", function () {
        const a = Array.from({ length: 9 }, (_, i) => i);
        const res = getCombinations(a, 4);
        assertEquals(res.length, 126);
    });
    // await t.step("Input sized. No req el", function () {
    //     const a = Array.from({ length: 222 }, (_, i) => i);
    //     const res = getCombinations(a, 8);
    //     assertEquals(res.length, 80440);
    // });
    // await t.step("Example 1", function () {
    //     const a = Array.from({ length: 5 }, (_, i) => i);
    //     const res = getCombinations(a, 3, 1);
    //     assertEquals(res.length, 6);
    // });
    // await t.step("Example 2", function () {
    //     const a = Array.from({ length: 4 }, (_, i) => i);
    //     const res = getCombinations(a, 3, 1);
    //     assertEquals(res.length, 3);
    // });
    // await t.step("Input sized", function () {
    //     const a = Array.from({ length: 222 }, (_, i) => i);
    //     const res = getCombinations(a, 8, 100);
    //     assertEquals(res.length, 80440);
    // });
});

function getSwappableWires(device: Map<string, Gate>, wire: string, expected: bigint): Set<string> {
    const result = new Set<string>();

    const gate = device.get(wire);
    assertExists(gate, "No gate");
    const wireA = gate[0];
    const wireB = gate[2];
    const op = gate[1];

    if (wireA === null || wireB === null) return result;

    const gateA = device.get(wireA);
    assertExists(gateA, "No gate A");

    const gateB = device.get(wireB);
    assertExists(gateB, "No gate B");

    const a = resolveGate(device, gateA);
    const b = resolveGate(device, gateB);

    assertExists(op, "No op");

    if (op === "AND") {
        if (expected === 1n) {
            if (a !== 1n) result.add(wireA);
            if (b !== 1n) result.add(wireB);
        } else if (expected === 0n) {
            if (a & b) result.add(wireA).add(wireB);
        }
    } else if (op === "XOR") {
        if (expected === 1n) {
            if (!(a ^ b)) result.add(wireA).add(wireB);
        } else if (expected === 0n) {
            if (a ^ b) result.add(wireA).add(wireB);
        }
    } else if (op === "OR") {
        if (expected === 1n) {
            if (!(a | b)) result.add(wireA).add(wireB);
        } else if (expected === 0n) {
            if (a === 1n) result.add(wireA);
            if (b === 1n) result.add(wireB);
        }
    } else {
        throw new Error("Unknown operand");
    }

    return result;
}

Deno.test(getSwappableWires.name, async function (t) {
    await t.step("AND", async function (t) {
        await t.step("aaa wrong", function () {
            const device = new Map<string, Gate>();
            device.set("zzz", ["aaa", Op.And, "bbb", null]);
            device.set("aaa", [null, null, null, 0n]);
            device.set("bbb", [null, null, null, 1n]);
            const wires = getSwappableWires(device, "zzz", 1n);
            assertEquals(wires.size, 1);
            assertEquals(wires.has("aaa"), true);
        });
        await t.step("bbb wrong", function () {
            const device = new Map<string, Gate>();
            device.set("zzz", ["aaa", Op.And, "bbb", null]);
            device.set("aaa", [null, null, null, 1n]);
            device.set("bbb", [null, null, null, 0n]);
            const wires = getSwappableWires(device, "zzz", 1n);
            assertEquals(wires.size, 1);
            assertEquals(wires.has("bbb"), true);
        });
    });
});
