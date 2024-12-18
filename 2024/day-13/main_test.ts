import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test(function partOneExample() {
    const result = partOne("example");

    assertEquals(result, 480);
});

Deno.test(function partOneInput() {
    const result = partOne("input");

    assertEquals(result, 40069);
});


Deno.test(function partTwoExample() {
    const result = partTwo("example");

    assertEquals(result, 875318608908);
});

Deno.test(function partTwoInput() {
    const result = partTwo("input");

    assertEquals(result, 71493195288102);
});
