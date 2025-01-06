import { assertExists } from "@std/assert/exists";

export function partOne(filename: string): number {
    const input = getInput(filename);

    const pcIds = new Map<string, number>();
    const pcs: string[] = [];

    const links: [string, string][] = [];
    for (const line of input) {
        const link = line.split("-") as [string, string];
        links.push(link);

        for (const pc of link) {
            let pcId = pcIds.get(pc);
            if (pcId !== undefined) continue;
            pcId = pcs.length;
            pcs.push(pc);
            pcIds.set(pc, pcId);
        }
    }

    const network: boolean[][] = Array.from(
        { length: pcs.length },
        () => Array.from({ length: pcs.length }, () => false),
    );

    for (const link of links) {
        const pc1id = pcIds.get(link[0]);
        assertExists(pc1id, "pc1id does not exist");
        const pc2id = pcIds.get(link[1]);
        assertExists(pc2id, "pc2id does not exist");

        network[pc1id][pc2id] = true;
        network[pc2id][pc1id] = true;
    }

    let total = 0;
    const processed = new Set<string>();
    for (const [pc1, pc1id] of pcIds) {
        if (pc1[0] !== "t") continue;

        for (let pc2id = 0; pc2id < network[pc1id].length; pc2id++) {
            if (!network[pc1id][pc2id]) continue;

            let pc2 = "";
            for (const [pc, id] of pcIds) {
                if (pc2id !== id) continue;
                pc2 = pc;
                break;
            }
            assertExists(pc2, "pc2 is empty");

            for (let pc3id = 0; pc3id < pcs.length; pc3id++) {
                if (!network[pc3id][pc1id] || !network[pc3id][pc2id]) continue;

                const link = [pc1id, pc2id, pc3id].sort().join(",");
                if (processed.has(link)) continue;
                processed.add(link);

                total++;
            }
        }
    }

    return total;
}

export function partTwo(filename: string): string {
    const input = getInput(filename);

    const pcIds = new Map<string, number>();
    const pcs: string[] = [];

    const links: [string, string][] = [];
    for (const line of input) {
        const link = line.split("-") as [string, string];
        links.push(link);

        for (const pc of link) {
            let pcId = pcIds.get(pc);
            if (pcId !== undefined) continue;
            pcId = pcs.length;
            pcs.push(pc);
            pcIds.set(pc, pcId);
        }
    }

    const network: boolean[][] = Array.from(
        { length: pcs.length },
        () => Array.from({ length: pcs.length }, () => false),
    );

    for (const link of links) {
        const pc1id = pcIds.get(link[0]);
        assertExists(pc1id, "pc1id does not exist");
        const pc2id = pcIds.get(link[1]);
        assertExists(pc2id, "pc2id does not exist");

        network[pc1id][pc2id] = true;
        network[pc2id][pc1id] = true;
    }

    // console.log(network.map((row) => row.map((val) => val ? "#" : " ").join("")).join("\n"));

    let biggestLink: number[] = [];
    for (let pc1id = 0; pc1id < network.length; pc1id++) {
        for (let offset = 0; offset < network[pc1id].length; offset++) {
            const link: number[] = [];

            for (let pc2id = offset; pc2id < network[pc1id].length; pc2id++) {
                // is this pc connected?
                if (!network[pc1id][pc2id]) continue;

                // is this connected to every other pc in the link?
                let connected = true;
                for (const pcId of link) {
                    if (network[pc2id][pcId]) continue;
                    connected = false;
                    break;
                }
                if (!connected) continue;
                link.push(pc2id);
            }

            link.push(pc1id);

            if (biggestLink.length < link.length) {
                biggestLink = link;
            }
        }
    }

    return biggestLink.map((pcId) => pcs[pcId]).sort().join(",");
}

function getInput(filename: string) {
    return Deno.readTextFileSync(filename).trimEnd().split("\n");
}
