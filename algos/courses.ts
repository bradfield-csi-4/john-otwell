function canFinish(numCourses: number, prerequisites: number[][]): boolean {
    const g = makeGraph(prerequisites, numCourses)
    console.log(g)
    for (let i = 0; i < numCourses; i++) {
        if (isThereCycle(g, i)) {
            return false
        }
    }
    return true
};


function makeGraph(prereqs: number[][], numCourses: number): number[][] {
    let graph: number[][] = []
    for (let i = 0; i < prereqs.length; i++) {
        let courseName = prereqs[i][0]
        if (graph[courseName] === undefined) {
            graph[courseName] = [prereqs[i][1]]
        } else {
            graph[courseName] = graph[courseName].concat(prereqs[i][1])
        }
    }
    for (let j = 0; j < numCourses; j++) {
        if (graph[j] === undefined) {
            graph[j] = []
        }
    }
    return graph
}

function isThereCycle(graph: number[][], curr: number): boolean {
    function tramp(current: number, visited: number[]): boolean {
        if (visited.includes(current)) {
            return true
        } else {
            for (let i = 0; i < graph[current].length; i++) {
                if (tramp(graph[current][i], visited.concat(current))) {
                    return true
                }
            }
            return false
        }
    }
    return tramp(curr, [])
}

console.log(canFinish(2, [[0, 1], [1, 0]]))
