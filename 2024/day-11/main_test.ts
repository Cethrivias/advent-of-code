import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test(function partOneExample() {
    const result = partOne("example");

    assertEquals(result, 55312);
});

Deno.test(function partOneInput() {
    const result = partOne("input");

    assertEquals(result, 186203);
});


Deno.test(function partTwoExample() {
    const result = partTwo("example");

    assertEquals(result, 65601038650482);
});

Deno.test(function partTwoInput() {
    const result = partTwo("input");

    assertEquals(result, 221291560078593);
});
