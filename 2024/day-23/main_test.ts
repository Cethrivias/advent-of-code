import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example");

    assertEquals(result, 7);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input");

    assertEquals(result, 1323);
});

Deno.test("Part Two. Example", function () {
    const result = partTwo("example_two");

    assertEquals(result, "co,de,ka,ta");
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input");

    assertEquals(result, "er,fh,fi,ir,kk,lo,lp,qi,ti,vb,xf,ys,yu");
});
