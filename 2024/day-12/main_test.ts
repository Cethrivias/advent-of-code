import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test(function partOneExample() {
    const result = partOne("example");

    assertEquals(result, 1930);
});

Deno.test(function partOneInput() {
    const result = partOne("input");

    assertEquals(result, 1377008);
});

Deno.test(function partTwoExample() {
    const result = partTwo("example");

    assertEquals(result, 1206);
});

Deno.test(function partTwoExampleOne() {
    const result = partTwo("part-two-example-one");

    assertEquals(result, 236);
});

Deno.test(function partTwoExampleTwo() {
    const result = partTwo("part-two-example-two");

    assertEquals(result, 368);
});

Deno.test(function partTwoInput() {
    const result = partTwo("input");

    assertEquals(result, 815788);
});
