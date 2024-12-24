import { assertEquals, assertLess } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example One", function () {
    const result = partOne("example_one");

    assertEquals(result, 7036);
});

Deno.test("Part One. Example Two", function () {
    const result = partOne("example_two");

    assertEquals(result, 11048);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input");

    assertEquals(result, 106512);
});

Deno.test("Part Two. Custom", function () {
    const result = partTwo("custom");

    assertEquals(result, 9);
});

Deno.test("Part Two. Example One", function () {
    const result = partTwo("example_one");

    assertEquals(result, 45);
});

Deno.test("Part Two. Example Two", function () {
    const result = partTwo("example_two");

    assertEquals(result, 64);
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input");

    assertEquals(result, 563);
});
