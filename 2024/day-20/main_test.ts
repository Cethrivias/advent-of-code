import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example", 4);

    assertEquals(result, 30);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input", 100);

    assertEquals(result, 1438);
});

Deno.test("Part Two. Example", function () {
    const result = partTwo("example", 50);

    assertEquals(result, 285);
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input", 100);

    assertEquals(result, 1026446);
});
