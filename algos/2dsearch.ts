function searchMatrix(matrix: number[][], target: number): boolean {
    // const midMat = Math.floor(matrix.length / 2)
    // if (matrix.length === 1) {
    //     return binarySearch(target, matrix[0])
    // } else if (matrix[midMat][0] > target) {
    //     return searchMatrix(matrix.slice(0, midMat), target)
    // } else if (matrix[midMat][0] < target && matrix[midMat + 1][0] < target) {
    //     return searchMatrix(matrix.slice(midMat + 1), target)
    // } else {
    //     return binarySearch(target, matrix[midMat])
    // }

    // a less stupid approach
    return binarySearch(target, matrix.reduce((p, n) => p.concat(n)))
};

function binarySearch(n: number, ns: number[]): boolean {
    const mid = Math.floor(ns.length / 2)
    if (n === ns[mid]) {
        return true
    }
    else if (ns.length === 0) {
        return false
    } else if (n > ns[mid]) {
        return binarySearch(n, ns.slice(mid + 1))
    } else {
        return binarySearch(n, ns.slice(0, mid))
    }
}

const matrix = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
console.log(searchMatrix(matrix, 1))
