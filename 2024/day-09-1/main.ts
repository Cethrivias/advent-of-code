const disk_map = Deno.readTextFileSync("input").trimEnd().split("").map(Number);

let total = 0;
let block_pos = 0;

for (let i = 0; i < disk_map.length; i++) {
    if (i % 2 == 0) {
        // file
        const file_id = i / 2;
        while (disk_map[i] > 0) {
            total += block_pos * file_id;
            disk_map[i]--;
            block_pos++;
        }
    } else {
        // empty space
        for (let j = disk_map.length - 1; j > i && disk_map[i] > 0; j -= 2) {
            const file_id = j / 2;
            while (disk_map[j] > 0 && disk_map[i] > 0) {
                total += block_pos * file_id;
                disk_map[j]--;
                disk_map[i]--;
                block_pos++;
            }
        }
    }
}

console.log(`Total: ${total}`);
