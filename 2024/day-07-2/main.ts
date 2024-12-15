// @deno-types="npm:@types/node"
import { createInterface as rl } from "node:readline/promises";
import { createReadStream } from "node:fs";

let total = 0;

const lines = rl({ input: createReadStream("input") });

for await (const line of lines) {
    const [test_value_string, numbers_strings] = line.split(": ");
    const test_value = parseInt(test_value_string);
    const numbers = numbers_strings.split(" ").map((it) => parseInt(it));
    if (attempt(test_value, numbers, 1, numbers[0])) {
        total += test_value;
    }
}

console.log("total:", total);

function attempt(
    test_value: number,
    numbers: number[],
    pos: number,
    current: number,
) {
    if (numbers.length === pos) {
        return test_value == current;
    }
    if (current > test_value) {
        return false;
    }

    if (attempt(test_value, numbers, pos + 1, current * numbers[pos])) {
        return true;
    }

    if (attempt(test_value, numbers, pos + 1, current + numbers[pos])) {
        return true;
    }

    if (attempt(test_value, numbers, pos + 1, parseInt(`${current}${numbers[pos]}`))) {
        return true;
    }

    return false;
}
