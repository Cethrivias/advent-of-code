class Position {
    constructor(public row: number, public col: number) {}
}

class Plot {
    constructor(public plant: string, public area: number, public perimeter: number, public sides: number) {}
}

const directions = [new Position(-1, 0), new Position(0, 1), new Position(1, 0), new Position(0, -1)];

export function partOne(filename: string): number {
    let total = 0;

    const input = getInput(filename);

    const visited: boolean[][] = [];
    for (let i = 0; i < input.length; i++) {
        visited[i] = [];
    }

    for (let row = 0; row < input.length; row++) {
        for (let col = 0; col < input[row].length; col++) {
            if (visited[row][col]) {
                continue;
            }
            const plot = new Plot(input[row][col], 0, 0, 0);
            const pos = new Position(row, col);
            calculatePrice(input, visited, plot, pos);
            total += plot.area * plot.perimeter;
        }
    }

    return total;
}

function calculatePrice(field: string[][], visited: boolean[][], plot: Plot, pos: Position): void {
    if (!withinBounds(field, pos)) {
        plot.perimeter++;
        return;
    }

    if (field[pos.row][pos.col] !== plot.plant) {
        plot.perimeter++;
        return;
    }

    if (visited[pos.row][pos.col]) {
        return;
    }

    visited[pos.row][pos.col] = true;

    plot.area++;

    for (const direction of directions) {
        const nextPos = new Position(pos.row + direction.row, pos.col + direction.col);
        calculatePrice(field, visited, plot, nextPos);
    }

    return;
}

export function partTwo(filename: string): number {
    const input: (string | number)[][] = getInput(filename);

    const visited: boolean[][] = [];
    for (let i = 0; i < input.length; i++) {
        visited[i] = [];
    }

    // assign unique ids to each plot and mark them on the field
    const plots: Plot[] = []; // references to the plots (id === idx)
    for (let row = 0; row < input.length; row++) {
        for (let col = 0; col < input[row].length; col++) {
            if (visited[row][col]) {
                continue;
            }
            const plot = new Plot(input[row][col] as string, 0, 0, 0);
            const pos = new Position(row, col);
            mapPlot(input, visited, plots, null, plot, pos);
        }
    }

    // Now all elements of the field were updated in place to be numbers (plot ids)
    // and we can drop string from the types
    const mappedField = input as number[][];

    // Going left to right. Checking if the top or bottom are the sides of a current plot
    let directions = [new Position(-1, 0), new Position(1, 0)]; // above and below
    let currentPlot = plots[0]; // 0:0 will always be the first plot
    const sides: boolean[] = []; // flags marking if there is a side above or below
    for (let row = 0; row < mappedField.length; row++) {
        for (let col = 0; col < mappedField[row].length; col++) {
            const nextPlot = plots[mappedField[row][col]];

            // Next plot has started meaning that any sides that we were tracing ended
            if (currentPlot !== nextPlot) {
                for (let i = 0; i < sides.length; i++) {
                    if (sides[i]) {
                        currentPlot.sides++;
                        sides[i] = false;
                    }
                }
                currentPlot = nextPlot;
            }

            for (let i = 0; i < directions.length; i++) {
                const sidePosition = new Position(row + directions[i].row, col + directions[i].col);

                // Next to the edge of the field. Always a side
                if (!withinBounds(mappedField, sidePosition)) {
                    sides[i] = true;
                    continue;
                }

                // Next to another plot
                const sidePlotId = mappedField[sidePosition.row][sidePosition.col];
                if (currentPlot !== plots[sidePlotId]) {
                    sides[i] = true;
                    continue;
                }

                // Were tracing a side, but the plot expanded in that direction. So the side ended
                if (sides[i]) {
                    currentPlot.sides++;
                    sides[i] = false;
                }
            }
        }

        // row done
        for (let i = 0; i < sides.length; i++) {
            if (sides[i]) {
                currentPlot.sides++;
                sides[i] = false;
            }
        }
    }

    // Going top to bottom. Checking if the top or bottom are the sides of a current plot
    directions = [new Position(0, 1), new Position(0, -1)];
    currentPlot = plots[0]; // 0:0 will always be the first plot
    sides[0] = false;
    sides[1] = false;
    for (let col = 0; col < mappedField[0].length; col++) {
        for (let row = 0; row < mappedField.length; row++) {
            const nextPlot = plots[mappedField[row][col]];

            // Next plot has started meaning that any sides that we were tracing ended
            if (currentPlot !== nextPlot) {
                for (let i = 0; i < sides.length; i++) {
                    if (sides[i]) {
                        currentPlot.sides++;
                        sides[i] = false;
                    }
                }
                currentPlot = nextPlot;
            }

            for (let i = 0; i < directions.length; i++) {
                const sidePosition = new Position(row + directions[i].row, col + directions[i].col);

                // Next to the edge of the field. Always a side
                if (!withinBounds(mappedField, sidePosition)) {
                    sides[i] = true;
                    continue;
                }

                // Next to another plot
                const sidePlotId = mappedField[sidePosition.row][sidePosition.col];
                if (currentPlot !== plots[sidePlotId]) {
                    sides[i] = true;
                    continue;
                }

                // Only one case left - same plot in the adjacent position meaning
                // We were tracing a side, but the plot expanded in that direction and the side ended
                if (sides[i]) {
                    currentPlot.sides++;
                    sides[i] = false;
                }
            }
        }

        // column done
        for (let i = 0; i < sides.length; i++) {
            if (sides[i]) {
                currentPlot.sides++;
                sides[i] = false;
            }
        }
    }

    let total = 0;
    for (const plot of plots) {
        total += plot.area * plot.sides;
    }

    return total;
}

function mapPlot(
    field: (string | number)[][],
    visited: boolean[][],
    plots: Plot[],
    plotId: number | null,
    plot: Plot,
    pos: Position,
): void {
    if (!withinBounds(field, pos)) {
        return;
    }

    if (field[pos.row][pos.col] !== plot.plant) {
        return;
    }

    if (visited[pos.row][pos.col]) {
        return;
    }

    visited[pos.row][pos.col] = true;

    if (plotId === null) {
        plotId = plots.length;
        plots[plotId] = plot;
    }

    plot.area++;
    field[pos.row][pos.col] = plotId;

    for (const direction of directions) {
        const nextPos = new Position(pos.row + direction.row, pos.col + direction.col);
        mapPlot(field, visited, plots, plotId, plot, nextPos);
    }
}

function getInput(filename: string): string[][] {
    return Deno.readTextFileSync(filename).trimEnd().split("\n").map((it) => it.split(""));
}

function withinBounds(input: unknown[][], pos: Position): boolean {
    return pos.row >= 0 && pos.row < input.length && pos.col >= 0 && pos.col < input[pos.row].length;
}
