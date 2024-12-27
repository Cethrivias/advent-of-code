import { assertEquals, assertFalse } from "@std/assert";

class State {
    A = 0;
    B = 0;
    C = 0;
    ptr = 0;
    out: number[] = [];
}

function parseComboOperand(s: State, operand: number): number {
    switch (operand) {
        case 0:
        case 1:
        case 2:
        case 3:
            return operand;
        case 4:
            return s.A;
        case 5:
            return s.B;
        case 6:
            return s.C;
        default:
            throw new Error(`Unknown operand: ${operand}`);
    }
}

function executeInstruction(s: State, opcode: number, operand: number): number {
    let nextPtr = s.ptr + 2;

    // console.log(`opcode: ${opcode}, operand: ${operand}`);
    switch (opcode) {
        case 0: // adv - division
            s.A = Math.floor(s.A / Math.pow(2, parseComboOperand(s, operand)));
            break;
        case 1: // bxl - bitwise XOR
            s.B ^= operand;
            break;
        case 2: // bst - lowest 3 bits
            s.B = parseComboOperand(s, operand) % 8;
            break;
        case 3: // jnz - jump - goto A
            if (s.A !== 0) {
                nextPtr = operand;
            }
            break;
        case 4: // bxc
            s.B ^= s.C;
            break;
        case 5: // out
            s.out.push(parseComboOperand(s, operand) % 8);
            break;
        case 6: // bdv - division
            s.B = Math.floor(s.A / Math.pow(2, parseComboOperand(s, operand)));
            break;
        case 7: // cdv - division
            s.C = Math.floor(s.A / Math.pow(2, parseComboOperand(s, operand)));
            break;
    }

    return nextPtr;
}

export function partOne(filename: string): number[] {
    const input = getInput(filename);

    const state = new State();

    state.A = Number(input[0].slice(12));
    state.B = Number(input[1].slice(12));
    state.C = Number(input[2].slice(12));

    const program = input[4].slice(9).split(",").map(Number);

    run(state, program);

    return state.out;
}

function run(s: State, program: number[]) {
    while (s.ptr + 1 < program.length) {
        s.ptr = executeInstruction(s, program[s.ptr], program[s.ptr + 1]);
    }
}

export function partTwo(filename: string): bigint {
    const input = getInput(filename);

    const program = input[4].slice(9).split(",").map(BigInt);

    const find = (target: bigint[], ans: bigint): bigint | null => {
        if (target.length === 0) return ans;

        for (let t = 0n; t < 8n; t++) {
            const a: bigint = (ans << 3n) | t;
            let b = 0n;
            let c = 0n;
            let output: bigint | null = null;
            let adv = false;

            const combo = (operand: bigint): bigint => {
                if (0 <= operand && operand <= 3) return operand;
                if (operand === 4n) return a;
                if (operand === 5n) return b;
                if (operand === 6n) return c;
                throw new Error(`Unknown combo operand: ${operand}`);
            };

            for (let pointer = 0; pointer < program.length - 2; pointer += 2) {
                const ins = program[pointer];
                const operand = program[pointer + 1];

                if (ins === 0n) {
                    assertFalse(adv, "ADV can happen only once per segment");
                    assertEquals(operand, 3n, "Can only shift A by one segment");
                    adv = true;
                } else if (ins === 1n) {
                    b ^= operand;
                } else if (ins === 2n) {
                    b = combo(operand) & 0b111n;
                } else if (ins === 3n) {
                    // noop.
                } else if (ins === 4n) {
                    b ^= c;
                } else if (ins === 5n) {
                    output = combo(operand) & 0b111n;
                } else if (ins === 6n) {
                    b = a >> combo(operand);
                } else if (ins === 7n) {
                    c = a >> combo(operand);
                }

                if (output === target[target.length - 1]) {
                    // found a potential A value
                    const sub = find(target.slice(0, target.length - 1), a);
                    // this value is a dead-end. Trying next
                    if (sub === null) continue;

                    return sub;
                }
            }
        }

        return null;
    };

    return find(program, 0n)!;
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}

Deno.test("Run. Instructions", async function (t) {
    await t.step("Example One", function () {
        const state = new State();
        state.C = 9;
        const program = [2, 6];
        run(state, program);

        assertEquals(state.B, 1);
    });
    await t.step("Example Two", function () {
        const state = new State();
        state.A = 10;
        const program = [5, 0, 5, 1, 5, 4];
        run(state, program);

        assertEquals(state.out.join(","), "0,1,2");
    });
    await t.step("Example Three", function () {
        // If register A contains 2024,
        // the program 0,1,5,4,3,0
        // would output 4,2,5,6,7,7,7,7,3,1,0
        // and leave 0 in register A.
        const state = new State();
        state.A = 2024;
        const program = [0, 1, 5, 4, 3, 0];
        run(state, program);

        assertEquals(state.out.join(","), "4,2,5,6,7,7,7,7,3,1,0");
        assertEquals(state.A, 0);
    });
    await t.step("Example Four", function () {
        // If register B contains 29,
        // the program 1,7
        // would set register B to 26.
        const state = new State();
        state.B = 29;
        const program = [1, 7];
        run(state, program);

        assertEquals(state.B, 26);
    });
    await t.step("Example Five", function () {
        // If register B contains 2024
        // and register C contains 43690,
        // the program 4,0
        // would set register B to 44354
        const state = new State();
        state.B = 2024;
        state.C = 43690;
        const program = [4, 0];
        run(state, program);

        assertEquals(state.B, 44354);
    });
});
