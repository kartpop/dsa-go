// Priority Queue implementation
// Make a priority queue of points on the x-y plane.
// Points closest to the origin have highest priority,
// if two points are equidistant from origin, one with lower x coordinate wins,
// if they have the same x coordinates, one with lower y coordinate wins.

package priorityqueue

import (
	"math"
)

type Point struct {
	x, y int
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	idist, jdist := distance(*pq[i]), distance(*pq[j])
	if idist != jdist {
		return idist < jdist
	} else if pq[i].x != pq[j].x {
		return pq[i].x < pq[j].x
	}
	return pq[i].y < pq[j].y
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	p := x.(*Point)
	*pq = append(*pq, p)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	retval := old[n-1]
	*pq = old[:n-1]
	return retval
}

func distance(p Point) float64 {
	return math.Sqrt(float64(powint(p.x, 2) + powint(p.y, 2)))
}

func powint(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
