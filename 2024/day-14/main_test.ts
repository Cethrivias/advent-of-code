import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test(function partOneExample() {
    const result = partOne("example", 7, 11);

    assertEquals(result, 12);
});

Deno.test(function partOneInput() {
    const result = partOne("input", 103, 101);

    assertEquals(result, 218619324);
});

Deno.test(function partTwoInput() {
    const result = partTwo("input", 103, 101);

    assertEquals(result, 6446);
});
