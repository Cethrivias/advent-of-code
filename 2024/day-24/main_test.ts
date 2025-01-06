import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part One. Example", function () {
    const result = partOne("example");

    assertEquals(result, 4n);
});

Deno.test("Part One. Example Two", function () {
    const result = partOne("example_two");

    assertEquals(result, 2024n);
});

Deno.test("Part One. Input", function () {
    const result = partOne("input");

    assertEquals(result, 45923082839246n);
});

Deno.test("Part Two. Input", function () {
    const result = partTwo("input_modified");

    assertEquals(result, "jgb,rkf,rrs,rvc,vcg,z09,z20,z24");
});
