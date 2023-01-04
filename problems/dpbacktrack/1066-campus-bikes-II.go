/*
https://leetcode.com/problems/campus-bikes-ii/

On a campus represented as a 2D grid, there are n workers and m bikes, with n <= m. Each worker and bike is 
a 2D coordinate on this grid.

We assign one unique bike to each worker so that the sum of the Manhattan distances between each worker and 
their assigned bike is minimized.

Return the minimum possible sum of Manhattan distances between each worker and their assigned bike.

The Manhattan distance between two points p1 and p2 is Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|.

 

Example 1:


Input: workers = [[0,0],[2,1]], bikes = [[1,2],[3,3]]
Output: 6
Explanation: 
We assign bike 0 to worker 0, bike 1 to worker 1. The Manhattan distance of both assignments is 3, 
so the output is 6.
Example 2:


Input: workers = [[0,0],[1,1],[2,0]], bikes = [[1,0],[2,2],[2,1]]
Output: 4
Explanation: 
We first assign bike 0 to worker 0, then assign bike 1 to worker 1 or worker 2, bike 2 to worker 2 or 
worker 1. Both assignments lead to sum of the Manhattan distances as 4.
Example 3:

Input: workers = [[0,0],[1,0],[2,0],[3,0],[4,0]], bikes = [[0,999],[1,999],[2,999],[3,999],[4,999]]
Output: 4995
 

Constraints:

n == workers.length
m == bikes.length
1 <= n <= m <= 10
workers[i].length == 2
bikes[i].length == 2
0 <= workers[i][0], workers[i][1], bikes[i][0], bikes[i][1] < 1000
All the workers and the bikes locations are unique.
*/

package dbbacktrack

import "math"

// Backtracking with memoization
// Memoize bikes allocated using 10 element array of 0s and 1s --> 10 bikes maximum

func AssignBikes(workers [][]int, bikes [][]int) int {
    dist := createDistanceMatrix(workers, bikes)
    m, n := len(bikes), len(workers)
    var visited [10]bool
    memo := make(map[[10]bool]int)
    
    var assign func(w int) int
    assign = func(w int) int {
        if w == n { 
            return 0
        }
        if sum, ok := memo[visited]; ok {
            return sum
        }
        
        sum := math.MaxInt
        for b := 0; b < m; b++ {
            if !visited[b] {
                visited[b] = true
                sum = min(dist[w][b] + assign(w+1), sum)
                visited[b] = false
            }
        }
        
        memo[visited] = sum
        return sum
    }
    
    return assign(0)
}

func createDistanceMatrix(workers, bikes [][]int) [][]int {
    abs := func(i int) int {
        if i < 0 { return -i }
        return i
    }
    
    distance := func(w, b int) int {
        wx, wy := workers[w][0], workers[w][1]
        bx, by := bikes[b][0], bikes[b][1]
        return abs(wx - bx) + abs(wy - by)
    }
    
    dist := make([][]int, len(workers))
    for w := 0; w < len(workers); w++ {
        dist[w] = make([]int, len(bikes))
        for b := 0; b < len(bikes); b++ {
            dist[w][b] = distance(w, b)
        }
    }
    
    return dist
}

func min(i, j int) int {
    if i < j { return i }
    return j
}