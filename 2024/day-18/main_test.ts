import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example", 7, 12);

    assertEquals(result, 22);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input", 71, 1024);

    assertEquals(result, 380);
});

Deno.test("Part Two. Example One", function () {
    const result = partTwo("example", 7, 12);

    assertEquals(result, "6,1");
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input", 71, 1024);

    assertEquals(result, "26,50");
});
