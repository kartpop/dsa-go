package priorityqueue

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQueue{}

	p := &Point{x: 2, y: -3}
	heap.Push(&pq, p) // pq -> [{2,-3}]

	p = &Point{x: 1, y: 3}
	heap.Push(&pq, p) // pq -> [{1,3}, {2,-3}]

	popped := heap.Pop(&pq) // pq -> [{2,-3}]
	p, ok := popped.(*Point)
	if !ok {
		t.Errorf("Expected point to be of tyep *Point, but is not.")
	}
	if p.x != 1 || p.y != 3 {
		t.Errorf("Expected point{1,3} to be at the top of PriorityQueue, but got point{%d,%d}.", p.x, p.y)
	}

	p = &Point{x: -2, y: 3}
	heap.Push(&pq, p) // pq -> [{-2,3}, {2,-3}]

	popped = heap.Pop(&pq) // pq -> [{2,-3}]
	p, ok = popped.(*Point)
	if !ok {
		t.Errorf("Expected point to be of tyep *Point, but is not.")
	}
	if p.x != -2 || p.y != 3 {
		t.Errorf("Expected point{-2,3} to be at the top of PriorityQueue, but got point{%d,%d}.", p.x, p.y)
	}

	p = &Point{x: 2, y: 3}
	heap.Push(&pq, p) // pq -> [{2,-3}, {2,3}]

	popped = heap.Pop(&pq) // pq -> [{2,3}]
	p, ok = popped.(*Point)
	if !ok {
		t.Errorf("Expected point to be of tyep *Point, but is not.")
	}
	if p.x != 2 || p.y != -3 {
		t.Errorf("Expected point{2,-3} to be at the top of PriorityQueue, but got point{%d,%d}.", p.x, p.y)
	}
}