// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import (
	"container/heap"
	"math"
)

// "A blazingly fast shortest path algorithm" as Tim Roughgarden would put it,
// Dijkstra's Algorithm computes the shortest path from a given source vertex
// and all vertices within the graph. A value of math.MaxInt32 is used to
// represent infinity when a vertex is not reachable from the given source
// vertex.
func (g *Graph) DijkstrasShortestPath(source *Vertex) map[string]int {

	verts := make(map[string]*dijkstraVertex, 0)
	pq := make(dijkstraPriorityQueue, 0)
	dist := make(map[string]int, 0)

	i := 0

	// Fill minheap
	for _, vert := range g.Vertices {
		item := &dijkstraVertex{
			vertex:   vert,
			distance: math.MaxInt32,
			index:    i}

		verts[item.vertex.Label] = item
		pq = append(pq, item)

		i += 1
	}

	// Define shortest path v -> v and heapify
	verts[source.Label].distance = 0
	heap.Init(&pq)

	for pq.Len() > 0 {
		// Grab min edge weight
		u := heap.Pop(&pq).(*dijkstraVertex)

		for _, edge := range u.vertex.outgoing {

			// Compute new edge weights
			v := verts[edge.Head.Label]
			d := u.distance + edge.Weight

			if v.distance > d {
				// Update distances if new min
				v.distance = d
				heap.Fix(&pq, v.index)

				// Correct output vector
				dist[edge.Head.Label] = d
			}
		}
	}

	return dist
}

// A wrapper struct for adding Dijkstra properties to graph's vertices
type dijkstraVertex struct {
	vertex   *Vertex
	distance int
	index    int
}

// A priority queue is an array of items
type dijkstraPriorityQueue []*dijkstraVertex

// Necessary for heap interface
func (pq dijkstraPriorityQueue) Len() int { return len(pq) }

// Necessary for heap interface
func (pq dijkstraPriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

// Necessary for heap interface
func (pq dijkstraPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Adds new edge to heap
func (pq *dijkstraPriorityQueue) Push(x interface{}) {
	vert := x.(*dijkstraVertex)
	vert.index = len(*pq)

	*pq = append(*pq, vert)
}

// Removes item from heap
func (pq *dijkstraPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	vert := old[n-1]
	vert.index = -1
	*pq = old[0 : n-1]

	return vert
}
