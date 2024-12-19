import { partOne, partTwo } from "./main.ts";

Deno.bench({
    name: "Part One",
    fn() {
        partOne("input", 103, 101);
    },
});

Deno.bench({
    name: "Part Two",
    fn() {
        partTwo("input", 103, 101);
    },
});
