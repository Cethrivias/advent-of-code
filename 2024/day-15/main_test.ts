import { assertEquals } from "@std/assert";
import { partOne, partTwo } from "./main.ts";

Deno.test("Part one. Example", () => {
    const result = partOne("example");

    assertEquals(result, 10092);
});

Deno.test("Part one. Input", () => {
    const result = partOne("input");

    assertEquals(result, 1552879);
});

Deno.test("Part two. Example", () => {
    const result = partTwo("example");

    assertEquals(result, 9021);
});

Deno.test("Part two. Input", () => {
    const result = partTwo("input");

    assertEquals(result, 1561175);
});
