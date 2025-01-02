import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example");

    assertEquals(result, -1);
});

// Deno.test("Part One. Input", function () {
//     const result = partOne("input");
//
//     assertEquals(result, -1);
// });
//
//
// Deno.test("Part Two. Example", function () {
//     const result = partTwo("example");
//
//     assertEquals(result, -1);
// });
//
// Deno.test("Part Two. Input", function () {
//     const result = partTwo("input");
//
//     assertEquals(result, -1);
// });
