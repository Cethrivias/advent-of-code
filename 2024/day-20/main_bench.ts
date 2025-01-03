import { partOne, partTwo } from "./main.ts";

Deno.bench({
    name: "Part One",
    fn() {
        partOne("input", 100);
    },
});

Deno.bench({
    name: "Part Two",
    fn() {
        partTwo("input", 100);
    },
});
