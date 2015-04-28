// Copyright 2015 Benjamin Campbell, All rights reserved
// The use of this code has been provisioned under an Apache License that is
// provided within this repository.

package graph

import (
	"errors"
	"math"
)

// An implementation of Bellman-Ford shortest path distance algorithm. The
// algorithm uses a variation of dynamic programming to produce a shortest path
// weight for each vertex relative to the given source. In the event of a
// negative cycle it returns nil with an error of a negative cycle.
func (g *Graph) BellmanFordShortestPath(source *Vertex) (map[string]int, error) {

	d := make(map[string]int, 0)

	for _, vert := range g.Vertices {
		d[vert.Label] = math.MaxInt32
	}

	d[source.Label] = 0

	for i := 0; i < len(g.Vertices)-1; i++ {
		for _, edge := range g.Edges {

			u := edge.Tail
			v := edge.Head

			if d[u.Label]+edge.Weight < d[v.Label] {
				d[v.Label] = d[u.Label] + edge.Weight
			}
		}
	}

	// Run the iterative step once more
	for _, edge := range g.Edges {
		u := edge.Tail
		v := edge.Head

		// If a distance decreases without a new vertex to gain on, it is
		// evident of a negative cycle within the graph
		if d[u.Label]+edge.Weight < d[v.Label] {
			return nil, errors.New("Graph contians a negitve weight cycle")
		}
	}

	return d, nil
}
