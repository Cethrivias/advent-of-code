import { assertEquals, assertGreater } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example");

    assertEquals(result, 37327623n);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input");

    assertEquals(result, 12759339434n);
});

Deno.test("Part Two. Example", function () {
    const result = partTwo("example_two");

    assertEquals(result, 23n);
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input");

    assertGreater(result, 1332n);
    assertEquals(result, 1405n);
});
