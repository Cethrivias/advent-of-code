// 6349492251099

const input = Deno.readTextFileSync("input").trimEnd().split("").map(Number);

// Because in part 2 we cannot densely pack every segment
// we need to keep track of how much space is left in each segment
// to accurately count current block position
class Segment {
    constructor(public size: number, public taken: number) {}

    free() {
        return this.size - this.taken;
    }
}

let total = 0;
let block_pos = 0;

const disk_map: Segment[] = [];
for (let i = 0; i < input.length; i++) {
    const seg = i % 2 === 0 ? new Segment(input[i], input[i]) : new Segment(input[i], 0);
    disk_map.push(seg);
}

for (let i = 0; i < disk_map.length; i++) {
    if (i % 2 == 0) {
        // file
        const file_id = i / 2;
        // Going through the blocks occupied by the segment
        for (let j = 0; j < disk_map[i].size; j++) {
            if (disk_map[i].taken > 0) {
                // File was not moved
                total += block_pos * file_id;
                disk_map[i].taken--;
            }
            block_pos++;
        }
    } else {
        // empty space
        for (let j = disk_map.length - 1; j > i && disk_map[i].free() > 0; j -= 2) {
            // Searching for files if there is free space left
            if (disk_map[i].free() >= disk_map[j].taken) {
                // file fits into the free space
                const file_id = j / 2;
                while (disk_map[j].taken > 0) {
                    total += block_pos * file_id;
                    disk_map[j].taken--;
                    disk_map[i].taken++;
                    block_pos++;
                }
            }
        }
        // Free blocks left in the segment, but no file fits
        block_pos += disk_map[i].free();
        disk_map[i].taken = 0;
    }
}

console.log(`Total: ${total}`);
