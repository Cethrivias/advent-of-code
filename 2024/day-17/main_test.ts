import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example");

    assertEquals(result.join(","), "4,6,3,5,6,3,5,2,1,0");
});

Deno.test("Part One. Input", function () {
    const result = partOne("input");

    assertEquals(result.join(","), "3,6,3,7,0,7,0,3,0");
});

Deno.test("Part Two. Example", function () {
    // actual input and example have different behaviour
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input");

    assertEquals(result, 136904920099226n);
});
