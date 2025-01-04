import { assertEquals, assertGreater, assertLess, assertNotEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example");

    assertEquals(result, 126384);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input");

    assertLess(result, 161472);
    assertGreater(result, 156544);
    assertEquals(result, 157908);
});

Deno.test("Part Two. Example", async function (t) {
    await t.step("2 keypads", function () {
        const result = partTwo("example", 2);

        assertEquals(result, 126384n);
    });
});

Deno.test("Part Two. Input", async function (t) {
    await t.step("2 keypads", function () {
        const result = partTwo("input", 2);

        assertEquals(result, 157908n);
    });

    await t.step("25 keypads", function () {
        const result = partTwo("input", 25);

        assertLess(result, 294184272873948n);
        assertGreater(result, 116281085558300n);
        assertEquals(result, 196910339808654n);
    });
});
