function numIslands(grid: string[][]): number {
    let n = 0;
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[0].length; j++) {
            if (grid[i][j] === '1' && !(neighbors(grid, i, j).includes('2'))) {
                n++
            }
            if (grid[i][j] === '1') {
                grid[i][j] = '2'
            }
        }
    }
    return n
}

function neighbors(grid: string[][], i: number, j: number): string[] {
    let ns = []
    ns.push(grid[i][j - 1], grid[i][j + 1])
    if (grid[i - 1] !== undefined) {
        ns.push(grid[i - 1][j])
    }
    if (grid[i + 1] !== undefined) {
        ns.push(grid[i + 1][j])
    }
    return ns.filter(s => s !== undefined)
}

